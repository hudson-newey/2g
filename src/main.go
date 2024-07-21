package main

import (
	"os"
	"strings"

	"github.com/hudson-newey/2g/src/commands"
)

func main() {
	args := strings.Join(os.Args[1:], " ")

	if commands.IsCustomCommand(os.Args) {
		commands.ExecuteCustomCommand(args)
		return
	}

	gitCommand := "git " + args
	commands.Execute(gitCommand)
}
