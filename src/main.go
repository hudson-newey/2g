package main

import (
	"os"
	"strings"

	"github.com/hudson-newey/2g/src/commands"
)

func main() {
	// remove the first argument which is the program name and join the rest
	args := strings.Join(os.Args[1:], "")
	gitCommand := "git " + args

	if commands.IsCustomCommand(args) {
		commands.ExecuteCustomCommand(args)
		return
	}

	commands.Execute(gitCommand)
}
