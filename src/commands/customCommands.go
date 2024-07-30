package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/hudson-newey/2g/shared/config"
)

func IsCustomCommand(command []string) bool {
	switch command[1] {
	case "explore", "install", "clone-file", "cache-clone":
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
	case "cache-clone":
		CacheCloneRepo(commandParam[1])
	default:
		invalidCommand(commandParam[0])
	}
}

// an optimized version of the clone command that will clone the repo to
// ~/.local/share/2g/cache and then copy the cloned repo to the current
// directory.
// any future attempts to clone the repo will attempt to update the cached
// repository instead of cloning the whole repository again.
//
// TODO: I hope to make this the default clone patch, but I don't have 100%
// confidence in it
func CacheCloneRepo(resourceUrl string) {
	if resourceUrl == "" {
		fmt.Println("Please provide a git URL to clone")
		os.Exit(1)
	}

	repoName := strings.Split(resourceUrl, "/")
	topLevelCacheLocation := config.CacheLocation()
	cacheLocation := expandPath(topLevelCacheLocation + "/" + repoName[len(repoName)-1] + "/")

	currentPath := os.Getenv("PWD")
	localPath := currentPath + "/" + repoName[len(repoName)-1]

	setupCommand := "mkdir -p " + topLevelCacheLocation
	Execute(setupCommand)

	// see if we have a cached version of the repository available
	_, err := os.Stat(cacheLocation)
	if err != nil {
		// we had a cache miss and we should clone the repository
		// to the cache location
		//
		// we do a shallow clone so that the user will get the code as quick
		// as possible (you usually don't need the whole history of a repo to start developing)
		// the git history will be fetched and progressively updated by the daemon in the
		// background while the user is developing, and is hopefully available by the time
		// the user wants to push their changes
		fmt.Println("Cache miss! Cloning", resourceUrl)
		cloneCommand := "git clone --depth 1 --single-branch --branch=main " + resourceUrl + " " + cacheLocation
		Execute(cloneCommand)

		// send a request to the daemon to fetch git history
		// this will be done in the background so that the user can start
		// developing as soon as possible
		configLocation := config.ConfigLocation()
		appendToFile(configLocation, "init-repo:"+cacheLocation+" "+localPath+"\n")
	} else {
		// there was a cache hit! we should attempt to update the cache through
		// a git pull
		fmt.Println("Cache hit! Updating", resourceUrl)
		updateCommand := "git -C " + cacheLocation + " pull --rebase"
		Execute(updateCommand)
	}

	copyCommand := "cp -r " + cacheLocation + " " + localPath
	Execute(copyCommand)

	fmt.Println("Cloned", resourceUrl)
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
		"git clone --depth 1 --single-branch --branch=main " + resourceUrl + " " + tempDir,
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
		"git clone --depth 1 --single-branch --branch=main" + resourceUrl + " " + installLocation,
		"echo 'export PATH=$PATH:" + installLocation + "' >> ~/.bashrc",
		"echo 'export PATH=$PATH:" + installLocation + "' >> ~/.zshrc",
	}

	ExecuteCommands(commandsToRun)
}

func invalidCommand(command string) {
	fmt.Println("Invalid command", command)
	os.Exit(1)
}

func expandPath(path string) string {
	homePath := os.Getenv("HOME")
	result := strings.ReplaceAll(path, "~", homePath)
	return result
}

func appendToFile(path string, contents string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			panic(err)
		}
		file.Close()
	}

	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if _, err := file.WriteString(contents); err != nil {
		panic(err)
	}
}
