// Command fx-repro is the CLI entry point. Subcommands are stubs.
package main

import (
	"os"

	"github.com/Anatomydotexe/fx-repro/internal/cli"
)

func main() {
	os.Exit(cli.Run(os.Args[1:], os.Stdout, os.Stderr))
}
