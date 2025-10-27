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
		Usage: "print size of a file or directory; supports -r (recursive), -H (human-readable), -a (include hidden)",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "recursive",
				Aliases: []string{"r"},
				Usage:   "recursive size of directories (default: false)",
			},
			&cli.BoolFlag{
				Name:    "human",
				Aliases: []string{"H"},
				Usage:   "human-readable sizes (auto-select unit) (default: false)",
			},
			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"a"},
				Usage:   "include hidden files and directories (default: false)",
			},
		},

		Action: func(ctx context.Context, cmd *cli.Command) error {
			path := cmd.Args().Get(0)
			flags := cmd.LocalFlagNames()
			bytes, err := code.GetSize(path, flags)
			if err != nil {
				return err
			}
			result := fmt.Sprintf("%s	%s", code.FormatSize(bytes, flags), path)
			fmt.Println(result)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
