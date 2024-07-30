package actions

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/hudson-newey/2g/src/commands"
)

func RunConfig(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return
	}

	fileLines := readFile(path)
	removeFile(path)

	commands := map[string]func(string){
		"init-repo:": initRepo,
	}

	for _, line := range fileLines {
		if line == "" {
			continue
		}

		for command, action := range commands {
			if strings.HasPrefix(line, command) {
				action(strings.TrimPrefix(line, command))
			}
		}
	}
}

func initRepo(command string) {
	log.Println("init repo: " + command)

	splitCommand := strings.Split(command, " ")
	if len(splitCommand) < 2 {
		panic("Invalid init repo command. Format: init-repo: <cache-location> <local-path>")
	}

	cachePath := splitCommand[0]
	localPath := splitCommand[1]

	commandsToRun := []string{
		"git fetch --unshallow",

		// replace the local .git with the fully fetched .git directory
		// TODO: this should probably use a rebase instead
		"rm -rf " + localPath + "/.git",
		"cp -r " + cachePath + "/.git/ " + localPath + "/.git",
	}

	commands.ExecuteCommandsInDir(commandsToRun, cachePath)
}

func readFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func removeFile(path string) {
	err := os.Remove(path)
	if err != nil {
		panic(err)
	}
}
