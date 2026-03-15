package main

import (
	"fmt"
	"os"
)

const version = "1.0.0"

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	switch cmd {
	case "new":
		if len(args) < 1 {
			fatal("Usage: gramkit new <project-name>")
		}
		cmdNew(args[0])

	case "make:handler":
		if len(args) < 1 {
			fatal("Usage: gramkit make:handler <name>")
		}
		cmdMakeHandler(args[0])

	case "make:middleware":
		if len(args) < 1 {
			fatal("Usage: gramkit make:middleware <name>")
		}
		cmdMakeMiddleware(args[0])

	case "run":
		cmdRun()

	case "webhook:set":
		if len(args) < 1 {
			fatal("Usage: gramkit webhook:set <url>")
		}
		cmdWebhookSet(args[0])

	case "webhook:delete":
		cmdWebhookDelete()

	case "webhook:info":
		cmdWebhookInfo()

	case "version", "--version", "-v":
		fmt.Printf("gramkit v%s\n", version)

	case "help", "--help", "-h":
		printHelp()

	default:
		fmt.Printf("Unknown command: %s\n\n", cmd)
		printHelp()
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Println(`gramkit - Telegram Bot Framework for Go

Usage:
  gramkit <command> [arguments]

Available Commands:
  new <name>              Create a new bot project
  make:handler <name>     Generate a new handler
  make:middleware <name>   Generate a new middleware
  run                     Run the bot (go run .)
  webhook:set <url>       Set webhook URL
  webhook:delete          Delete webhook
  webhook:info            Get webhook info
  version                 Print version
  help                    Show this help`)
}

func fatal(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}
