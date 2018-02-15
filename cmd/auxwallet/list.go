package main

import (
	"fmt"

	"github.com/jawher/mow.cli"
)

func list(cfg config) func(*cli.Cmd) {
	return func(cmd *cli.Cmd) {
		cmd.Action = func() {
			fmt.Println("list transactions not yet implemented")
		}
	}
}
