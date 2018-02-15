package main

import (
	"fmt"

	"github.com/jawher/mow.cli"
)

func send(cfg config) func(*cli.Cmd) {
	return func(cmd *cli.Cmd) {
		addr := cmd.StringArg("ADDRESS", "", "recipient address")
		amt := cmd.StringArg("AMOUNT", "", "send amount")

		cmd.Action = func() {
			fmt.Println("send not yet implemented")
			_, _ = addr, amt
		}
	}
}
