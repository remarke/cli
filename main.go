package main

import (
	"log"

	"github.com/remarke/cli/config"
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

func main() {
	log.SetFlags(0)

	Cmd.Run()
}

var Cmd = &Z.Cmd{
	Name:     `rem`,
	Commands: []*Z.Cmd{help.Cmd, config.Cmd},
}
