package cmd

import (
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

type CmdArg func() error

func Execute(version string) error {
	app := cli.NewApp()
	app.Name = "sonarcheck"
	app.Usage = "sonarcheck"
	app.Version = version
	app.EnableBashCompletion = true
	app.Compiled = time.Now()
	app.Copyright = "(c) 2022 heloo"
	app.Authors = []*cli.Author{{Name: "heloo", Email: "developer@heloo.com.br"}}

	app.Commands = []*cli.Command{
		Check,
	}

	return app.Run(os.Args)
}
