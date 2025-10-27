package main

import (
	"code"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory;",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "human",
				Aliases: []string{"H"},
				Usage:   "human-readable sizes (auto-select unit)",
			},
			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"a"},
				Usage:   "include hidden files and directories",
			},
		},

		Action: func(ctx context.Context, cmd *cli.Command) error {
			path := cmd.Args().Get(0)
			flags := cmd.LocalFlagNames()
			bytes, err := code.GetSize(path, flags)

			if len(cmd.LocalFlagNames()) > 0 {
				if err != nil {
					return err
				}
				fmt.Println(code.FormatSize(bytes, flags))
			} else {
				fmt.Println(code.FormatSize(bytes, []string{}))
			}
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
