package main

import (
	"resonance/pkg/author"
	"resonance/pkg/cmd"

	"github.com/urfave/cli/v2"
)

func main() {
	// app := cli.App()
	// app.Name = "port_scanner"
	// app.Author = "netxfly"
	// app.Email = "x@xsec.io"
	// app.Version = "2020/3/8"
	// app.Usage = "tcp syn/connect port scanner"
	// app.Commands = []cli.Command{cmd.Scan}
	// app.Flags = append(app.Flags, cmd.Scan.Flags...)
	// err := app.Run(os.Args)
	// _ = err
	app := &cli.App{
		Name: "port_sacnner",
		Authors: []*cli.Author{
			&(author.Author),
		},
		Version: "2023.3.28",
		Usage:   "start to port scan",
		Commands: []*cli.Command{
			&(cmd.Scan),
		},
		Flags: []cli.Flag{
			cmd.Scan.Flags,
		},
	}
}
