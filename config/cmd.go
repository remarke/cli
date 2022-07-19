package config

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/remarke/cli/git"
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/fs/file"
	"github.com/rwxrob/help"
	"gopkg.in/yaml.v3"
)

var Cmd = &Z.Cmd{
	Name:     `config`,
	Aliases:  []string{"c", "conf"},
	Commands: []*Z.Cmd{help.Cmd, initialize, show, edit},
}

var edit = &Z.Cmd{
	Name:        `edit`,
	Description: `edit command open user configuration using the default $EDITOR`,
	Call: func(_ *Z.Cmd, args ...string) error {
		var config Config

		return file.Edit(config.getConfigFilePath())
	},
}

var show = &Z.Cmd{
	Name: `show`,
	Call: func(_ *Z.Cmd, args ...string) error {
		var config Config
		data, err := config.GetConfig()

		if err != nil {
			log.Fatalf("Error while reading the configuration file %v", err)
		}

		fmt.Printf("Editor: %s\nPublic Folder: %s\nPrivate Folder: %s\n", data.Editor, data.PublicFolder, data.PrivateFolder)

		return nil
	},
}

var initialize = &Z.Cmd{
	Name: `init`,
	Call: func(_ *Z.Cmd, args ...string) error {
		base, _ := os.UserHomeDir()

		config := Config{
			Editor:        "vim",
			PublicFolder:  path.Join(base, "Repos", "remarke"),
			PrivateFolder: path.Join(base, "Private", "remarke"),
		}

		marshallConfig, _ := yaml.Marshal(&config)

		fileCreated, err := config.setConfigFile(marshallConfig)
		os.Mkdir(config.PublicFolder, 0755)
		os.Mkdir(path.Join(base, "Private"), 0755)
		os.Mkdir(config.PrivateFolder, 0755)

		if err := git.Initialize(config.PublicFolder); err != nil {
			log.Fatalf("Could not initialize git repository on public folder: %v", err)
			return err
		}

		if err := git.Initialize(config.PrivateFolder); err != nil {
			log.Fatalf("Could not initialize git repository on public folder: %v", err)
			return err
		}

		if err != nil || !fileCreated {
			log.Fatalf("Could not write the file: %v", err)
			return err
		}

		if fileCreated {
			log.Print("File created succesfully")
		}

		return nil
	},
}
