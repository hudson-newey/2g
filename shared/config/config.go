package config

import "os"

func ConfigLocation() string {
	return programStorageLocation() + "/config"
}

func CacheLocation() string {
	return programStorageLocation() + "/cache"
}

func DaemonLockLocation() string {
	return programStorageLocation() + "/daemon.lock"
}

func programStorageLocation() string {
	return "/home/" + currentUser() + "/.local/share/2g"
}

func currentUser() string {
	return os.Getenv("USER")
}
