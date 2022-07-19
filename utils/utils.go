package utils

import (
	"log"
	"os"
	"os/exec"
)

func RunShell(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
