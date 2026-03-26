package main

import (
	"fmt"
	"os"
	"strings"
)

const version = "1.2.0"

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

	case "make:command":
		if len(args) < 1 {
			fatal("Usage: gramkit make:command <name>")
		}
		cmdMakeCommand(args[0])

	case "make:middleware":
		if len(args) < 1 {
			fatal("Usage: gramkit make:middleware <name>")
		}
		cmdMakeMiddleware(args[0])

	case "run":
		flags := parseFlags(args)
		cmdRun(flags["debug"], flags["watch"])

	case "run:dev":
		flags := parseFlags(args)
		cmdRunDev(flags["debug"], flags["watch"])

	case "doctor":
		cmdDoctor()

	case "update":
		cmdUpdate()

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

Scaffolding:
  new <name>              Create a new bot project
  make:handler <name>     Generate a message handler
  make:command <name>     Generate a command handler (/start, /help, etc.)
  make:middleware <name>   Generate a middleware

Running:
  run                     Run the bot (go run .)
  run:dev                 Run with hot-reload (dev mode)

Bot API:
  webhook:set <url>       Set webhook URL
  webhook:delete          Delete webhook
  webhook:info            Get webhook info

Tools:
  doctor                  Check environment and configuration
  update                  Update gramkit CLI to latest version
  version                 Print version
  help                    Show this help

Flags (for run / run:dev):
  --debug                 Start delve debugger (attach IDE to :2345)
  --watch                 Start web dashboard for updates (:4080)`)
}

func fatal(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}

// parseFlags extracts --flag style flags from args.
func parseFlags(args []string) map[string]bool {
	flags := make(map[string]bool)
	for _, arg := range args {
		if strings.HasPrefix(arg, "--") {
			flags[arg[2:]] = true
		}
	}
	return flags
}
