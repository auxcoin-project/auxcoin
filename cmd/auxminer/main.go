package main

import (
	"fmt"
	"os"

	"github.com/auxcoin-project/auxcoin/pb"
	"github.com/jawher/mow.cli"
	"google.golang.org/grpc"
)

type config struct {
	srvAddr  *string
	coinAddr *string
}

func main() {
	app := cli.App("auxminer", "auxcoin miner")

	app.Spec = "[-a] ADDRESS"
	app.Action = action(config{
		srvAddr: app.String(cli.StringOpt{
			Name:  "a addr",
			Value: ":9999",
			Desc:  "address of auxcoind instance",
		}),
		coinAddr: app.StringArg("ADDRESS", "", "coin address"),
	})

	app.Run(os.Args)
}

func newClient(cfg config) pb.AuxcoinClient {
	c, err := grpc.Dial(*cfg.srvAddr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return pb.NewAuxcoinClient(c)
}

func action(cfg config) func() {
	return func() {
		fmt.Println("mining not yet implemented")
	}
}
