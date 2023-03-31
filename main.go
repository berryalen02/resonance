package main

import (
	"fmt"
	"os"
	"resonance/pkg/author"
	"resonance/pkg/cmd"

	"github.com/urfave/cli/v2"
)

func main() {
	Scan_app := &cli.App{
		Name:        "resonance",
		Description: "A growing comprehensive scanner.",
		Authors: []*cli.Author{
			&(author.Author),
		},
		Version: "2023.3.31",
		// Usage:   "start a scan server",
		// Action: func(ctx *cli.Context) error {
		// 	return util.Scan(ctx)
		// },
		Commands: []*cli.Command{
			&(cmd.PortScan),
		},
	}
	err := Scan_app.Run(os.Args)
	//panic(err)
	if err != nil {
		fmt.Printf("%v", err)
	}
	_ = err
}
