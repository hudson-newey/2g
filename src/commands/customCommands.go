package commands

import (
	"fmt"
	"os"
	"strings"
)

func IsCustomCommand(command []string) bool {
	switch command[1] {
	case "explore":
		return true
	}

	return false
}

func ExecuteCustomCommand(command string) {
	splitCommand := strings.Split(command, " ")

	switch splitCommand[0] {
	case "explore":
		ExploreRepo(splitCommand[1])
	}
}

func ExploreRepo(resourceUrl string) {
	if resourceUrl == "" {
		fmt.Println("Please provide a git URL to explore")
		os.Exit(1)
	}

	tempDir := "/tmp/2g"
	Execute("mkdir " + tempDir)
	Execute("git clone " + resourceUrl + " " + tempDir)
	Execute("yazi " + tempDir)
	Execute("rm -rf " + tempDir)
}
