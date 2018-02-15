package main

import (
	"fmt"

	"github.com/jawher/mow.cli"
)

func address(cfg config) func(*cli.Cmd) {
	return func(cmd *cli.Cmd) {
		cmd.Command("new", "new address", newAddress(cfg))
		cmd.Command("list ls", "list addresses", listAddress(cfg))
	}
}

func newAddress(cfg config) func(*cli.Cmd) {
	return func(cmd *cli.Cmd) {
		cmd.Action = func() {
			fmt.Println("new address not yet implemented")
		}
	}
}

func listAddress(cfg config) func(*cli.Cmd) {
	return func(cmd *cli.Cmd) {
		priv := cmd.BoolOpt("p priv", false, "show private")

		cmd.Action = func() {
			fmt.Println("list address not yet implemented")
			_ = priv
		}
	}
}
