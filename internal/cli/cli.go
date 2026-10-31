// Package cli parses argv for the fx-repro commands.
// scan, plan, record, status, and export are stubs.
package cli

import (
	"fmt"
	"io"
)

// NotImplemented is the stub result for every command.
const NotImplemented = "not implemented yet"

const helpText = `fx-repro reduces a FiveM (FXServer) crash to a minimal shareable repro.

You confirm each trial. fx-repro does not detect client crashes.

Usage:
  fx-repro <command>

Commands:
  scan     Read resources and their manifests
  plan     Propose the next reduced configuration
  record   Record a trial as crash, works, or unknown
  status   Show the resumable trial history
  export   Write a minimal shareable repro

Status: early scaffold, not usable yet.
`

// commands are the planned workflow verbs. Each one is a stub.
var commands = []string{"scan", "plan", "record", "status", "export"}

// Run executes one invocation. It returns the process exit code.
func Run(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 || args[0] == "-h" || args[0] == "--help" || args[0] == "help" {
		fmt.Fprint(stdout, helpText)
		return 0
	}

	name := args[0]
	for _, command := range commands {
		if command == name {
			fmt.Fprintln(stdout, NotImplemented)
			return 1
		}
	}

	fmt.Fprintf(stderr, "unknown command %q\n\n%s", name, helpText)
	return 2
}
