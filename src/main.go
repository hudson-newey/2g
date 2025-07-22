package main

import (
	"os"
	"strings"

	"github.com/hudson-newey/2g/src/commands"
)

func main() {
	if len(os.Args) < 2 {
		commands.Execute("git")
		return
	}

	if commands.IsCustomCommand(os.Args) {
		commands.ExecuteCustomCommand(os.Args)
		return
	}

	concatenatedArgs := strings.Join(os.Args[1:], " ")
	gitCommand := "git " + concatenatedArgs

	commands.Execute(gitCommand)
}
