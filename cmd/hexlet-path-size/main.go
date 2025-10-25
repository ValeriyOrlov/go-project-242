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
				Name:  "human",
				Usage: "human-readable sizes (auto-select unit)",
			},
		},

		Action: func(ctx context.Context, cmd *cli.Command) error {
			path := cmd.Args().Get(0)
			bytes, err := code.GetSize(path)

			if len(cmd.LocalFlagNames()) > 0 {
				flagName := cmd.LocalFlagNames()[0]
				if err != nil {
					return err
				}
				fmt.Println(code.FormatSize(bytes, flagName))
			} else {
				fmt.Println(code.FormatSize(bytes, ""))
			}
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
