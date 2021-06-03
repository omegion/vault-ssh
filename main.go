package main

import (
	"os"

	"github.com/omegion/vault-ssh/cmd"
)

func main() {
	commander := cmd.NewCommander()
	commander.Setup()

	if err := commander.Root.Execute(); err != nil {
		os.Exit(1)
	}
}
