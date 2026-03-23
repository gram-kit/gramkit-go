package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/fsnotify/fsnotify"
)

const (
	// debounceDelay prevents rapid rebuilds when multiple files change at once.
	debounceDelay = 300 * time.Millisecond
	buildOutput   = "./tmp/bot"
	defaultWatch  = ":4080"
	defaultDebug  = ":2345"
	// gracefulTimeout is how long to wait after SIGTERM before SIGKILL.
	gracefulTimeout = 2 * time.Second
)

// skipDirs are directories that should not be watched.
var skipDirs = map[string]bool{
	"tmp": true, "vendor": true, "testdata": true, "node_modules": true,
}

// managedCmd wraps exec.Cmd with a done channel to safely track process exit.
// This avoids calling cmd.Wait() from multiple goroutines.
type managedCmd struct {
	cmd  *exec.Cmd
	done chan struct{}
}

func cmdRunDev(debug, watch bool) {
	loadEnv()

	fmt.Println("\033[36m[gramkit:dev]\033[0m Starting in dev mode with hot-reload...")

	if watch {
		os.Setenv("GRAMKIT_WATCH_ADDR", defaultWatch)
		fmt.Printf("\033[35m[gramkit:watch]\033[0m Dashboard will start at http://localhost%s\n", defaultWatch)
	}
	if debug {
		ensureDlv()
		fmt.Printf("\033[33m[gramkit:debug]\033[0m Delve will listen on localhost%s\n", defaultDebug)
		fmt.Println("\033[33m[gramkit:debug]\033[0m Attach your IDE to localhost" + defaultDebug)
	}

	fmt.Println("\033[36m[gramkit:dev]\033[0m Watching for .go file changes (fsnotify)...")
	fmt.Println()

	_ = os.MkdirAll("tmp", 0o755)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fatal(fmt.Sprintf("Failed to create watcher: %v", err))
	}
	defer watcher.Close()

	watchDirs(watcher, ".")

	var (
		proc *managedCmd
		mu   sync.Mutex
	)

	// Graceful shutdown on SIGINT/SIGTERM.
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigCh
		fmt.Println("\n\033[36m[gramkit:dev]\033[0m Shutting down...")
		mu.Lock()
		stopProcess(proc)
		mu.Unlock()
		_ = os.RemoveAll("tmp")
		os.Exit(0)
	}()

	// Initial build and run.
	if err := build(); err != nil {
		fmt.Printf("\033[31m[gramkit:dev]\033[0m Build failed:\n%s\n", err)
	} else {
		mu.Lock()
		proc = startProc(debug)
		mu.Unlock()
	}

	// Debounce timer.
	var timer *time.Timer

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}

			if !isWatchedFile(event.Name) {
				continue
			}

			// Watch new directories.
			if event.Has(fsnotify.Create) {
				if info, err := os.Stat(event.Name); err == nil && info.IsDir() {
					watchDirs(watcher, event.Name)
				}
			}

			// Capture event name for the closure (event is reused each iteration).
			changedFile := event.Name

			if timer != nil {
				timer.Stop()
			}
			timer = time.AfterFunc(debounceDelay, func() {
				fmt.Println()
				fmt.Printf("\033[33m[gramkit:dev]\033[0m File changed: %s\n", changedFile)

				mu.Lock()
				stopProcess(proc)
				proc = nil
				mu.Unlock()

				if err := build(); err != nil {
					fmt.Printf("\033[31m[gramkit:dev]\033[0m Build failed:\n%s\n", err)
					return
				}

				mu.Lock()
				proc = startProc(debug)
				mu.Unlock()
			})

		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			fmt.Printf("\033[31m[gramkit:dev]\033[0m Watcher error: %v\n", err)
		}
	}
}

// build compiles the project with debug flags (no optimizations).
func build() error {
	fmt.Println("\033[36m[gramkit:dev]\033[0m Building...")
	t := time.Now()

	cmd := exec.Command("go", "build", "-gcflags=all=-N -l", "-o", buildOutput, ".")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()

	if err := cmd.Run(); err != nil {
		return err
	}

	fmt.Printf("\033[32m[gramkit:dev]\033[0m Built in %s\n", time.Since(t).Round(time.Millisecond))
	return nil
}

// startProc launches the binary, optionally through delve.
func startProc(debug bool) *managedCmd {
	var cmd *exec.Cmd

	if debug {
		fmt.Println("\033[33m[gramkit:debug]\033[0m Starting with delve...")
		cmd = exec.Command("dlv", "exec", buildOutput,
			"--headless",
			"--listen="+defaultDebug,
			"--api-version=2",
			"--accept-multiclient",
			"--continue",
		)
	} else {
		fmt.Println("\033[36m[gramkit:dev]\033[0m Starting bot...")
		cmd = exec.Command(buildOutput)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Env = os.Environ()

	// Create a new process group so we can kill the entire tree.
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	if err := cmd.Start(); err != nil {
		fmt.Printf("\033[31m[gramkit:dev]\033[0m Failed to start: %v\n", err)
		return nil
	}

	m := &managedCmd{
		cmd:  cmd,
		done: make(chan struct{}),
	}

	// Single goroutine that calls Wait and closes done.
	// No other code should call cmd.Wait().
	go func() {
		_ = cmd.Wait()
		close(m.done)
	}()

	return m
}

// stopProcess gracefully stops a process tree: SIGTERM first, then SIGKILL after timeout.
func stopProcess(m *managedCmd) {
	if m == nil || m.cmd == nil || m.cmd.Process == nil {
		return
	}

	// Check if already exited.
	select {
	case <-m.done:
		return
	default:
	}

	pgid, err := syscall.Getpgid(m.cmd.Process.Pid)

	// Step 1: Send SIGTERM for graceful shutdown.
	if err == nil {
		_ = syscall.Kill(-pgid, syscall.SIGTERM)
	} else {
		_ = m.cmd.Process.Signal(syscall.SIGTERM)
	}

	// Step 2: Wait for process to exit, with a timeout.
	select {
	case <-m.done:
		return
	case <-time.After(gracefulTimeout):
		// Step 3: Force kill if still alive.
		fmt.Println("\033[33m[gramkit:dev]\033[0m Process did not exit, force killing...")
		if err == nil {
			_ = syscall.Kill(-pgid, syscall.SIGKILL)
		} else {
			_ = m.cmd.Process.Kill()
		}
		<-m.done
	}
}

// watchDirs recursively adds directories to the watcher, skipping hidden/ignored dirs.
func watchDirs(w *fsnotify.Watcher, root string) {
	_ = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if !info.IsDir() {
			return nil
		}
		name := info.Name()
		// Skip hidden dirs (but not "." itself), vendor, tmp, etc.
		if name != "." && (skipDirs[name] || name[0] == '.') {
			return filepath.SkipDir
		}
		_ = w.Add(path)
		return nil
	})
}

// isWatchedFile returns true for .go, go.mod, and go.sum files.
func isWatchedFile(path string) bool {
	ext := filepath.Ext(path)
	base := filepath.Base(path)
	if ext == ".go" {
		return !strings.HasPrefix(path, "tmp")
	}
	return base == "go.mod" || base == "go.sum"
}

// ensureDlv checks if delve is installed, and auto-installs it if not.
func ensureDlv() {
	if _, err := exec.LookPath("dlv"); err == nil {
		return
	}

	fmt.Println("\033[33m[gramkit:debug]\033[0m Delve not found, installing automatically...")
	fmt.Println("\033[33m[gramkit:debug]\033[0m Running: go install github.com/go-delve/delve/cmd/dlv@latest")

	cmd := exec.Command("go", "install", "github.com/go-delve/delve/cmd/dlv@latest")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()

	if err := cmd.Run(); err != nil {
		fatal(fmt.Sprintf("Failed to install delve: %v\nInstall manually: go install github.com/go-delve/delve/cmd/dlv@latest", err))
	}

	if _, err := exec.LookPath("dlv"); err != nil {
		fatal("delve was installed but not found in PATH. Make sure $GOPATH/bin (or $GOBIN) is in your PATH.")
	}

	fmt.Println("\033[32m[gramkit:debug]\033[0m Delve installed successfully!")
}
