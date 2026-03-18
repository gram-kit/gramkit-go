package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

func cmdMakeHandler(name string) {
	dir := "handlers"
	if err := os.MkdirAll(dir, 0755); err != nil {
		fatal(fmt.Sprintf("Failed to create directory: %v", err))
	}

	filename := toSnakeCase(name) + ".go"
	path := filepath.Join(dir, filename)

	if _, err := os.Stat(path); err == nil {
		fatal(fmt.Sprintf("Handler already exists: %s", path))
	}

	funcName := toPascalCase(name)

	content := fmt.Sprintf(`package handlers

import (
	"context"

	gramkit "github.com/gram-kit/gramkit-go"
	"github.com/gram-kit/gramkit-go/models"
	"github.com/gram-kit/gramkit-go/params"
)

func %s(ctx context.Context, b *gramkit.Bot, msg *models.Message) {
	b.SendMessage(ctx, params.SendMessage{
		ChatID: msg.Chat.ID,
		Text:   "%s handler",
	})
}
`, funcName, funcName)

	writeFile(path, content)
	fmt.Printf("Handler created: %s\n", path)
	fmt.Printf("Register it in main.go:\n")
	fmt.Printf("  bot.HandleMessage(\"*\", handlers.%s)\n", funcName)
}

func cmdMakeCommand(name string) {
	dir := "handlers"
	if err := os.MkdirAll(dir, 0755); err != nil {
		fatal(fmt.Sprintf("Failed to create directory: %v", err))
	}

	cmdName := strings.TrimPrefix(name, "/")
	filename := toSnakeCase(cmdName) + ".go"
	path := filepath.Join(dir, filename)

	if _, err := os.Stat(path); err == nil {
		fatal(fmt.Sprintf("Command handler already exists: %s", path))
	}

	funcName := toPascalCase(cmdName) + "Command"

	content := fmt.Sprintf(`package handlers

import (
	"context"

	gramkit "github.com/gram-kit/gramkit-go"
	"github.com/gram-kit/gramkit-go/models"
	"github.com/gram-kit/gramkit-go/params"
)

func %s(ctx context.Context, b *gramkit.Bot, msg *models.Message) {
	b.SendMessage(ctx, params.SendMessage{
		ChatID: msg.Chat.ID,
		Text:   "/%s command",
	})
}
`, funcName, cmdName)

	writeFile(path, content)
	fmt.Printf("Command handler created: %s\n", path)
	fmt.Printf("Register it in main.go:\n")
	fmt.Printf("  bot.HandleCommand(\"%s\", handlers.%s)\n", cmdName, funcName)
}

func cmdMakeMiddleware(name string) {
	dir := "middleware"
	if err := os.MkdirAll(dir, 0755); err != nil {
		fatal(fmt.Sprintf("Failed to create directory: %v", err))
	}

	filename := toSnakeCase(name) + ".go"
	path := filepath.Join(dir, filename)

	if _, err := os.Stat(path); err == nil {
		fatal(fmt.Sprintf("Middleware already exists: %s", path))
	}

	funcName := toPascalCase(name)

	content := fmt.Sprintf(`package middleware

import (
	"context"

	gramkit "github.com/gram-kit/gramkit-go"
	"github.com/gram-kit/gramkit-go/models"
)

func %s(ctx context.Context, b *gramkit.Bot, update *models.Update, next gramkit.HandlerFunc) {
	// TODO: implement %s middleware logic

	next(ctx, b, update)
}
`, funcName, funcName)

	writeFile(path, content)
	fmt.Printf("Middleware created: %s\n", path)
	fmt.Printf("Register it in main.go:\n")
	fmt.Printf("  bot.Use(middleware.%s)\n", funcName)
}

// toSnakeCase converts "MyHandler" or "myHandler" to "my_handler"
func toSnakeCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				result.WriteRune('_')
			}
			result.WriteRune(unicode.ToLower(r))
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}

// toPascalCase converts "my_handler" or "my-handler" to "MyHandler"
func toPascalCase(s string) string {
	var result strings.Builder
	upper := true
	for _, r := range s {
		if r == '_' || r == '-' {
			upper = true
			continue
		}
		if upper {
			result.WriteRune(unicode.ToUpper(r))
			upper = false
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}
