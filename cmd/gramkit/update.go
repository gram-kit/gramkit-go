package main

import (
	"fmt"
	"os"
	"os/exec"
)

func cmdUpdate() {
	fmt.Println("Updating gramkit CLI...")

	cmd := exec.Command("go", "install", "github.com/gram-kit/gramkit-go/cmd/gramkit@latest")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()

	if err := cmd.Run(); err != nil {
		fatal(fmt.Sprintf("Failed to update: %v", err))
	}

	fmt.Println("\033[32mgramkit updated successfully!\033[0m")
	fmt.Println("Run 'gramkit version' to check the new version.")
}
