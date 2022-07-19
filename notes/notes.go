package notes

import (
	"fmt"
	"log"
	"os"

	"github.com/remarke/cli/config"
	"github.com/remarke/cli/git"
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/fs/file"
	"github.com/rwxrob/help"
	"github.com/rwxrob/uniq"
)

var Cmd = &Z.Cmd{
	Name:     `notes`,
	Aliases:  []string{"n"},
	Commands: []*Z.Cmd{help.Cmd, create},
}

var create = &Z.Cmd{
	Name:    `create`,
	Aliases: []string{"c"},
	Call: func(_ *Z.Cmd, args ...string) error {
		var config config.Config
		data, err := config.GetConfig()

		if err != nil {
			log.Fatalf("Error while reading the configuration file %v", err)
		}

		isosec := uniq.Isosec()
		os.Mkdir(fmt.Sprintf("%s/%s", data.PublicFolder, isosec), 0755)

		file.Edit(fmt.Sprintf("%s/%s/README.md", data.PublicFolder, isosec))

		if err := git.AddAll(data.PublicFolder); err != nil {
			return fmt.Errorf("Error while add your new note %s: %v", isosec, err)
		}

		return git.Commit(data.PublicFolder, fmt.Sprintf("Added %s note", isosec))
	},
}
