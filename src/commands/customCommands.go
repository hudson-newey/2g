package commands

import (
	"fmt"
	"os"
	"strings"
)

func IsCustomCommand(command []string) bool {
	switch command[1] {
	case "explore", "install":
		return true
	}

	return false
}

func ExecuteCustomCommand(command string) {
	splitCommand := strings.Split(command, " ")

	switch splitCommand[0] {
	case "explore":
		{
			ExploreRepo(splitCommand[1])
			break
		}
	case "install":
		{
			InstallRepo(splitCommand[1])
			break
		}
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

func InstallRepo(resourceUrl string) {
	repoName := strings.Split(resourceUrl, "/")
	installLocation := "~/.local/bin/" + repoName[len(repoName)-1]

	commandsToRun := []string{
		"git clone " + resourceUrl + " " + installLocation,
		"echo 'export PATH=$PATH:" + installLocation + "' >> ~/.bashrc",
		"echo 'export PATH=$PATH:" + installLocation + "' >> ~/.zshrc",
	}

	for _, command := range commandsToRun {
		Execute(command)
	}

	fmt.Println("Installed " + repoName[len(repoName)-1])
}
