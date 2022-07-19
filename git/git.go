package git

import (
	"fmt"
	"os"
	"os/exec"
)

func Initialize(path string) error {
	cmd := exec.Command("git", "init")
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func AddAll(path string) error {
	cmd := exec.Command("git", "add", ".")
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func Commit(path string, message string) error {
	cmd := exec.Command("git", "commit", "-m", fmt.Sprintf("\"%s\"", message))
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func Push(path string) error {
	cmd := exec.Command("git", "push")
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
