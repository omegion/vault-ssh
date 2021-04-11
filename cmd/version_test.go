package cmd

import (
	"testing"
)

func Test_GetCommand(t *testing.T) {
	_, err := executeCommand(Sign(), "--public-key=/Users/hakan/.ssh/id_rsa.pub")
	if err != nil {
		t.Errorf("Command Error: %v", err)
	}
}
