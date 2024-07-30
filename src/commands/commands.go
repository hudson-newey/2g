package commands

import (
	"fmt"
	"os"
	"os/exec"
)

func Execute(command string) {
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		fmt.Println("Error executing command: ", err)
		os.Exit(1)
	}
}

func ExecuteCommands(commands []string) {
	for _, command := range commands {
		Execute(command)
	}
}

func ExecuteCommandsInDir(commands []string, dir string) {
	for _, command := range commands {
		virtualCommand := fmt.Sprintf("cd %s && %s", dir, command)
		Execute(virtualCommand)
	}
}
