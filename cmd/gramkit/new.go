package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func cmdNew(name string) {
	dir, err := filepath.Abs(name)
	if err != nil {
		fatal(fmt.Sprintf("Invalid path: %v", err))
	}

	if _, err := os.Stat(dir); err == nil {
		fatal(fmt.Sprintf("Directory %s already exists", name))
	}

	fmt.Printf("Creating new gramkit project: %s\n", name)

	// Create directory structure
	dirs := []string{
		"",
		"handlers",
		"middleware",
	}
	for _, d := range dirs {
		if err := os.MkdirAll(filepath.Join(dir, d), 0755); err != nil {
			fatal(fmt.Sprintf("Failed to create directory: %v", err))
		}
	}

	moduleName := "github.com/user/" + name

	// Write main.go
	writeFile(filepath.Join(dir, "main.go"), mainTemplate(moduleName))

	// Write handlers/start.go
	writeFile(filepath.Join(dir, "handlers", "start.go"), startHandlerTemplate(moduleName))

	// Write handlers/help.go
	writeFile(filepath.Join(dir, "handlers", "help.go"), helpHandlerTemplate(moduleName))

	// Write middleware/logger.go
	writeFile(filepath.Join(dir, "middleware", "logger.go"), loggerMiddlewareTemplate(moduleName))

	// Write .env
	writeFile(filepath.Join(dir, ".env"), "BOT_TOKEN=your_bot_token_here\n")

	// Write .gitignore
	writeFile(filepath.Join(dir, ".gitignore"), dotGitignore())

	// Initialize go module
	fmt.Println("Initializing Go module...")
	runCmd(dir, "go", "mod", "init", moduleName)
	runCmd(dir, "go", "get", "github.com/gram-kit/gramkit-go@latest")
	runCmd(dir, "go", "mod", "tidy")

	fmt.Println()
	fmt.Println(strings.Repeat("─", 40))
	fmt.Printf("Project %s created successfully!\n\n", name)
	fmt.Println("Next steps:")
	fmt.Printf("  cd %s\n", name)
	fmt.Println("  # Set your bot token in .env")
	fmt.Println("  gramkit run")
	fmt.Println(strings.Repeat("─", 40))
}

func writeFile(path, content string) {
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		fatal(fmt.Sprintf("Failed to write %s: %v", path, err))
	}
}

func runCmd(dir, name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fatal(fmt.Sprintf("Command failed: %s %s: %v", name, strings.Join(args, " "), err))
	}
}

func mainTemplate(module string) string {
	return fmt.Sprintf(`package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	gramkit "github.com/gram-kit/gramkit-go"

	"%s/handlers"
	"%s/middleware"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("BOT_TOKEN is not set")
	}

	bot, err := gramkit.New(token,
		gramkit.WithDefaultHandler(handlers.Default),
		gramkit.WithMiddleware(middleware.Logger),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Register command handlers
	bot.HandleCommand("start", handlers.Start)
	bot.HandleCommand("help", handlers.Help)

	fmt.Println("Bot is running...")
	bot.StartPolling(ctx)
}
`, module, module)
}

func startHandlerTemplate(module string) string {
	return `package handlers

import (
	"context"

	gramkit "github.com/gram-kit/gramkit-go"
	"github.com/gram-kit/gramkit-go/models"
	"github.com/gram-kit/gramkit-go/params"
)

func Start(ctx context.Context, b *gramkit.Bot, msg *models.Message) {
	b.SendMessage(ctx, params.SendMessage{
		ChatID: msg.Chat.ID,
		Text:   "Welcome! I'm a bot built with gramkit.",
	})
}
`
}

func helpHandlerTemplate(module string) string {
	return `package handlers

import (
	"context"

	gramkit "github.com/gram-kit/gramkit-go"
	"github.com/gram-kit/gramkit-go/models"
	"github.com/gram-kit/gramkit-go/params"
)

func Help(ctx context.Context, b *gramkit.Bot, msg *models.Message) {
	b.SendMessage(ctx, params.SendMessage{
		ChatID: msg.Chat.ID,
		Text:   "Commands:\n/start - Start the bot\n/help - Show help",
	})
}
`
}

func loggerMiddlewareTemplate(module string) string {
	return `package middleware

import (
	"context"
	"log"

	gramkit "github.com/gram-kit/gramkit-go"
	"github.com/gram-kit/gramkit-go/models"
)

func Logger(ctx context.Context, b *gramkit.Bot, update *models.Update, next gramkit.HandlerFunc) {
	if update.Message != nil {
		log.Printf("[%d] %s: %s", update.UpdateID, update.Message.From.FirstName, update.Message.Text)
	}
	next(ctx, b, update)
}
`
}

func dotGitignore() string {
	return `.env
tmp/
*.exe
*.exe~
*.dll
*.so
*.dylib
*.test
*.out
vendor/
`
}
