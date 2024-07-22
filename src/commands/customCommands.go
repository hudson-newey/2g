package commands

import (
	"fmt"
	"os"
	"strings"
)

func IsCustomCommand(command []string) bool {
	switch command[1] {
	case "explore", "install", "clone-file":
		return true
	}

	// I have implemented a clone patch so that you can clone a single file
	// from a repository without downloading the whole repository
	// therefore, we have to conditionally use the custom clone command
	// if we are cloning a file instead of a repository
	if command[1] == "clone" {
		repositoryPath := strings.Split(command[2], ".git")[1]
		pathSplit := strings.Split(repositoryPath, "/")
		return len(pathSplit) > 1
	}

	return false
}

func ExecuteCustomCommand(command string) {
	commandParam := strings.Split(command, " ")

	switch commandParam[0] {
	case "explore":
		ExploreRepo(commandParam[1])
	case "install":
		InstallRepo(commandParam[1])
	case "clone":
		CloneSingle(commandParam[1])
	case "clone-file":
		CloneSingle(commandParam[1])
	}
}

// this is a patch for the clone command that allows you to clone a single file
// from a repository without downloading the whole repository
func CloneSingle(resourceUrl string) {
	splitUrl := strings.Split(resourceUrl, ".git")
	repositoryUrl := splitUrl[0]
	// we join by .git because files like .gitignore have .git in their names
	// but they are really files.
	pathUrl := strings.Join(splitUrl[1:], ".git")

	fmt.Println("Cloning file", pathUrl, "from", repositoryUrl)

	tempCloneDir := "/tmp/2g"

	commandsToRun := []string{
		"mkdir " + tempCloneDir,
		"git clone --depth 1 " + repositoryUrl + " " + tempCloneDir,
		"cp -r " + tempCloneDir + "/" + pathUrl + " .",
		"rm -rf " + tempCloneDir,
	}

	ExecuteCommands(commandsToRun)
}

func ExploreRepo(resourceUrl string) {
	if resourceUrl == "" {
		fmt.Println("Please provide a git URL to explore")
		os.Exit(1)
	}

	tempDir := "/tmp/2g"

	commandsToRun := []string{
		"mkdir " + tempDir,
		"git clone --depth 1 " + resourceUrl + " " + tempDir,
		"yazi " + tempDir,
		"rm -rf " + tempDir,
	}

	ExecuteCommands(commandsToRun)
}

func InstallRepo(resourceUrl string) {
	if resourceUrl == "" {
		fmt.Println("Please provide a git URL to install")
		os.Exit(1)
	}

	repoName := strings.Split(resourceUrl, "/")
	installLocation := "~/.local/bin/" + repoName[len(repoName)-1]

	commandsToRun := []string{
		"git clone " + resourceUrl + " " + installLocation,
		"echo 'export PATH=$PATH:" + installLocation + "' >> ~/.bashrc",
		"echo 'export PATH=$PATH:" + installLocation + "' >> ~/.zshrc",
	}

	ExecuteCommands(commandsToRun)
}
