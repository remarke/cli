package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func RunShell(name string, args ...string) {
	cmd := exec.Command(name, args...)
	err, _ := cmd.StderrPipe()
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(err)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func GetEnvironmentEditor() string {
	editor := os.Getenv("EDITOR")

	if len(editor) <= 0 {
		return "vim"
	}

	return editor
}
