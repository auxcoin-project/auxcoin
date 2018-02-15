package main

import (
	"os"

	"github.com/auxcoin-project/auxcoin/pb"
	"github.com/jawher/mow.cli"
	"google.golang.org/grpc"
)

type config struct {
	srvAddr *string
}

func main() {
	app := cli.App("auxwallet", "auxcoin wallet")

	cfg := config{
		srvAddr: app.String(cli.StringOpt{
			Name:  "a addr",
			Value: ":9999",
			Desc:  "address of auxcoind instance",
		}),
	}

	app.Command("address", "manage addresses", address(cfg))
	app.Command("send", "send coins", send(cfg))
	app.Command("list", "list transactions", list(cfg))

	app.Run(os.Args)
}

func newClient(cfg config) pb.AuxcoinClient {
	c, err := grpc.Dial(*cfg.srvAddr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return pb.NewAuxcoinClient(c)
}
