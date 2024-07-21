package commands

func IsCustomCommand(command string) bool {
	switch command {
	case "ls", "pwd":
		return true
	}

	return false
}

func ExecuteCustomCommand(command string) {
	switch command {
	case "ls":
		Execute("ls")
	case "pwd":
		Execute("pwd")
	}
}
