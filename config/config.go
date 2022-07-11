package config

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
	"gopkg.in/yaml.v3"
)

// Config holds all the user options for editor, folder locations, etc...
type Config struct {
	Editor        string
	PublicFolder  string
	PrivateFolder string
	ConfigFolder  string
}

var Cmd = &Z.Cmd{
	Name:     `config`,
	Aliases:  []string{"c", "conf"},
	Commands: []*Z.Cmd{help.Cmd, initialize},
}

var edit = &Z.Cmd{
	Name: `edit`,
	Call: func(_ *Z.Cmd, args ...string) error {
		data, err := readUserConfig()

		return nil
	},
}

var initialize = &Z.Cmd{
	Name: `init`,
	Call: func(_ *Z.Cmd, args ...string) error {
		base, _ := os.UserHomeDir()

		defaultConfig := Config{
			Editor:        "vim",
			PublicFolder:  path.Join(base, "Repos", "remarke"),
			PrivateFolder: path.Join(base, "Private", "remarke"),
			ConfigFolder:  path.Join(base, ".config/remarke"),
		}

		marshallConfig, _ := yaml.Marshal(&defaultConfig)

		fileCreated, err := createConfigFile(defaultConfig.ConfigFolder, marshallConfig)

		if err != nil || !fileCreated {
			log.Fatalf("Could not write the file: %v", err)
		}

		if fileCreated {
			log.Print("File created succesfully")
		}

		return nil
	},
}

func createConfigFile(path string, data []byte) (bool, error) {
	_, err := ioutil.ReadFile(path)

	if err != nil {
		runShell("mkdir", path)
		runShell("touch", fmt.Sprintf("%s/config.yaml", path))

		err = ioutil.WriteFile(fmt.Sprintf("%s/config.yaml", path), data, 0644)

		if err != nil {
			return false, err
		}
	}

	return true, nil
}

func runShell(name string, args ...string) {
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
