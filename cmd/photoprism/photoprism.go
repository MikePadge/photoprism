package main

import (
	"os"

	"github.com/mikepadge/photoprism/internal/commands"
	"github.com/mikepadge/photoprism/internal/config"
	"github.com/mikepadge/photoprism/internal/event"
	"github.com/urfave/cli"
)

var version = "development"
var log = event.Log

func main() {
	app := cli.NewApp()
	app.Name = "PhotoPrism"
	app.Usage = "Browse your life in pictures"
	app.Version = version
	app.Copyright = "(c) 2018-2020 The PhotoPrism contributors <hello@photoprism.org>"
	app.EnableBashCompletion = true
	app.Flags = config.GlobalFlags

	app.Commands = []cli.Command{
		commands.StartCommand,
		commands.StopCommand,
		commands.IndexCommand,
		commands.ImportCommand,
		commands.CopyCommand,
		commands.ConvertCommand,
		commands.ThumbsCommand,
		commands.MigrateCommand,
		commands.ConfigCommand,
		commands.VersionCommand,
	}

	if err := app.Run(os.Args); err != nil {
		log.Error(err)
	}
}
