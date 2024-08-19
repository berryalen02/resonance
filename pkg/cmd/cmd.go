package cmd

import (
	"resonance/pkg/task"
	"resonance/pkg/util"

	"github.com/urfave/cli/v2"
)

var Cli_PortScan = cli.Command{
	Name:        "portscan",
	Usage:       "start a port scan",
	Description: "start a port scan",
	Action: func(ctx *cli.Context) error {
		util.TargetsInit(ctx)
		task.PortScan()
		return nil
	},
	Flags: []cli.Flag{
		stringFlag("iplist", "i", "", "ip list"),
		stringFlag("port", "p", "", "port range (default: 'CommonPort')"),
		boolFlag("full", "f", "full port scan"),
		&cli.BoolFlag{
			Name:  "tcp",
			Value: true,
		},
		&cli.BoolFlag{
			Name: "syn",
		},
		intFlag("level", "l", 2, "Scan intensity level 0-4"),
		intFlag("timeout", "t", 3, "timeout"),
		intFlag("concurrency", "c", 1000, "concurrency"),
	},
}

func stringFlag(name, aliases, value, usage string) *cli.StringFlag {
	return &cli.StringFlag{
		Name:    name,
		Aliases: []string{aliases},
		Value:   value,
		Usage:   usage,
	}
}

func intFlag(name string, aliases string, value int, usage string) *cli.IntFlag {
	return &cli.IntFlag{
		Name:    name,
		Aliases: []string{aliases},
		Value:   value,
		Usage:   usage,
	}
}

func boolFlag(name string, aliases string, usage string) *cli.BoolFlag {
	return &cli.BoolFlag{
		Name:    name,
		Aliases: []string{aliases},
		// Value:   value,
		Usage: usage,
	}
}
