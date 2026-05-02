package main

import (
	"fmt"
	"os"

	"github.com/acme/actio/internal/adapters/cli"
)

func main() {
	if err := cli.Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
