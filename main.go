package main

import (
	"fmt"
	"os"
	"resonance/pkg/author"
	"resonance/pkg/cmd"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {
	// 定义字符画
	asciiArt := []string{
		"  _____                                                   ",
		" |  __ \\                                                  ",
		" | |__) | ___  ___   ___   _ __    __ _  _ __    ___  ___ ",
		" |  _  / / _ \\/ __| / _ \\ | '_ \\  / _` || '_ \\  / __|/ _ \\",
		" | | \\ \\|  __/\\__ \\| (_) || | | || (_| || | | || (__|  __/",
		" |_|  \\_\\\\___||___/ \\___/ |_| |_| \\__,_||_| |_| \\___|\\___|",
	}

	// 将字符画转换为字符串
	var sb strings.Builder
	for _, line := range asciiArt {
		sb.WriteString(line)
		sb.WriteString("\n")
	}
	asciiArtString := sb.String()

	fmt.Println(asciiArtString)
	Scan_app := &cli.App{
		Name:        "resonance",
		Description: "https://github.com/berryalen02/resonance",
		Authors: []*cli.Author{
			&(author.Author),
		},
		Version: "1.0",
		Usage:   "Your loyal attack assistant",
		// Action: func(ctx *cli.Context) error {
		// 	return util.Scan(ctx)
		// },
		Commands: []*cli.Command{
			&(cmd.Cli_PortScan),
		},
	}
	err := Scan_app.Run(os.Args)
	//panic(err)
	if err != nil {
		fmt.Printf("%v", err)
	}
	_ = err
}
