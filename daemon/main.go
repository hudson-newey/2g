package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hudson-newey/2g/daemon/actions"
	"github.com/hudson-newey/2g/shared/config"
)

func main() {
	singletonLockLocation := config.DaemonLockLocation()
	configLocation := config.ConfigLocation()

	if _, err := os.Stat(singletonLockLocation); err == nil {
		panic(
			"Failed to start daemon: Daemon is already running.\n" +
				"If this is incorrect, you can delete the lock file at " +
				singletonLockLocation,
		)
	}

	lockFile, err := os.Create(singletonLockLocation)
	if err != nil {
		panic("Failed to create daemon lock file")
	}
	defer lockFile.Close()

	// ensure the daemon cleans itself up if it is terminated
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigs
		cleanup()
		os.Exit(0)
	}()

	log.Println("2g Daemon started")

	for {
		actions.RunConfig(configLocation)
		time.Sleep(10 * time.Second)
	}
}

func cleanup() {
	singletonLockLocation := config.DaemonLockLocation()
	os.Remove(singletonLockLocation)

	log.Println("2g Daemon stopped")
}
