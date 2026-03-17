package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
)

func cmdRun(debug, watch bool) {
	loadEnv()

	if watch {
		os.Setenv("GRAMKIT_WATCH_ADDR", defaultWatch)
		fmt.Printf("\033[35m[gramkit:watch]\033[0m Dashboard will start at http://localhost%s\n", defaultWatch)
	}

	fmt.Println("Starting bot...")

	var cmd *exec.Cmd

	if debug {
		ensureDlv()
		fmt.Printf("\033[33m[gramkit:debug]\033[0m Delve listening on localhost%s\n", defaultDebug)
		fmt.Println("\033[33m[gramkit:debug]\033[0m Attach your IDE to localhost" + defaultDebug)

		// dlv debug compiles with debug flags and starts the debugger.
		cmd = exec.Command("dlv", "debug", ".",
			"--headless",
			"--listen="+defaultDebug,
			"--api-version=2",
			"--accept-multiclient",
			"--continue",
		)
	} else {
		cmd = exec.Command("go", "run", ".")
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Env = os.Environ()

	// Create a new process group so Ctrl+C kills the entire tree.
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	// Handle Ctrl+C: kill the process group and exit cleanly.
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	if err := cmd.Start(); err != nil {
		fatal(fmt.Sprintf("Failed to run: %v", err))
	}

	go func() {
		<-sigCh
		killProcessGroup(cmd)
	}()

	_ = cmd.Wait()
}

// loadEnv loads .env file into environment variables.
func loadEnv() {
	f, err := os.Open(".env")
	if err != nil {
		return // .env is optional
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		key, val, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}
		key = strings.TrimSpace(key)
		val = strings.TrimSpace(val)
		// Don't override existing env vars.
		if os.Getenv(key) == "" {
			os.Setenv(key, val)
		}
	}
}
