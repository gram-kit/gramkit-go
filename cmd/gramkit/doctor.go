package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func cmdDoctor() {
	fmt.Println("gramkit doctor")
	fmt.Println(strings.Repeat("─", 40))
	fmt.Println()

	allOk := true

	// 1. Go version
	allOk = checkGo() && allOk

	// 2. Delve
	allOk = checkDlv() && allOk

	// 3. .env file
	allOk = checkEnvFile() && allOk

	// 4. BOT_TOKEN validity
	allOk = checkBotToken() && allOk

	// 5. go.mod exists (are we in a project?)
	allOk = checkGoMod() && allOk

	// 6. System info
	printInfo("OS", runtime.GOOS+"/"+runtime.GOARCH)
	printInfo("gramkit", "v"+version)

	fmt.Println()
	fmt.Println(strings.Repeat("─", 40))
	if allOk {
		fmt.Println("\033[32mAll checks passed!\033[0m")
	} else {
		fmt.Println("\033[33mSome checks failed. See above for details.\033[0m")
	}
}

func checkGo() bool {
	out, err := exec.Command("go", "version").Output()
	if err != nil {
		printFail("Go", "not found in PATH")
		return false
	}
	ver := strings.TrimSpace(string(out))
	// Extract just "go1.24.1" from "go version go1.24.1 darwin/arm64"
	parts := strings.Fields(ver)
	if len(parts) >= 3 {
		ver = parts[2]
	}
	printOk("Go", ver)
	return true
}

func checkDlv() bool {
	out, err := exec.Command("dlv", "version").Output()
	if err != nil {
		printWarn("Delve", "not installed (needed for --debug, will auto-install on first use)")
		return true // not a hard failure
	}
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	ver := "installed"
	for _, line := range lines {
		if strings.HasPrefix(line, "Version:") {
			ver = strings.TrimSpace(strings.TrimPrefix(line, "Version:"))
			break
		}
	}
	printOk("Delve", ver)
	return true
}

func checkEnvFile() bool {
	if _, err := os.Stat(".env"); err != nil {
		printWarn(".env", "not found (create one with BOT_TOKEN=your_token)")
		return true // not a hard failure
	}
	printOk(".env", "found")
	return true
}

func checkBotToken() bool {
	loadEnv()
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		printFail("BOT_TOKEN", "not set")
		return false
	}
	if token == "your_bot_token_here" {
		printFail("BOT_TOKEN", "still set to placeholder value")
		return false
	}

	// Try getMe to validate token
	resp := apiCall(token, "getMe", nil)
	if !resp.Ok {
		printFail("BOT_TOKEN", fmt.Sprintf("invalid (%s)", resp.Description))
		return false
	}

	var bot struct {
		Username string `json:"username"`
		Name     string `json:"first_name"`
	}
	if err := json.Unmarshal(resp.Result, &bot); err == nil {
		printOk("BOT_TOKEN", fmt.Sprintf("valid (@%s — %s)", bot.Username, bot.Name))
	} else {
		printOk("BOT_TOKEN", "valid")
	}
	return true
}

func checkGoMod() bool {
	if _, err := os.Stat("go.mod"); err != nil {
		printWarn("Project", "no go.mod found (are you in a project directory?)")
		return true
	}
	printOk("Project", "go.mod found")
	return true
}

func printOk(name, detail string) {
	fmt.Printf("  \033[32m✓\033[0m %-14s %s\n", name, detail)
}

func printFail(name, detail string) {
	fmt.Printf("  \033[31m✗\033[0m %-14s %s\n", name, detail)
}

func printWarn(name, detail string) {
	fmt.Printf("  \033[33m!\033[0m %-14s %s\n", name, detail)
}

func printInfo(name, detail string) {
	fmt.Printf("  \033[36m·\033[0m %-14s %s\n", name, detail)
}
