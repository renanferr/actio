package cli

import (
	"fmt"
	"os"

	"github.com/acme/actio/internal/adapters/cli/commands"
)

func Run(args []string) error {
	if len(args) == 0 {
		return printHelp()
	}

	switch args[0] {
	case "init":
		return commands.RunInit(args[1:])
	case "validate":
		return commands.RunValidate(args[1:])
	case "apply":
		return commands.RunApply(args[1:])
	case "list":
		return commands.RunList(args[1:])
	case "inspect":
		return commands.RunInspect(args[1:])
	case "run":
		return commands.RunRun(args[1:])
	case "graph":
		return commands.RunGraph(args[1:])
	case "help":
		return printHelp()
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n", args[0])
		return printHelp()
	}
}

func printHelp() error {
	fmt.Println("Actio CLI (skeleton)")
	fmt.Println("Usage: actio <command> [options]")
	fmt.Println("Commands: init, validate, apply, list, inspect, run, graph")
	return nil
}
