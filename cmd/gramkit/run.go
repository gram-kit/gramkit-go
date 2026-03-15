package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func cmdRun() {
	// Load .env if exists
	loadEnv()

	fmt.Println("Starting bot...")
	cmd := exec.Command("go", "run", ".")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Env = os.Environ()

	if err := cmd.Run(); err != nil {
		fatal(fmt.Sprintf("Failed to run: %v", err))
	}
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
		// Don't override existing env vars
		if os.Getenv(key) == "" {
			os.Setenv(key, val)
		}
	}
}
