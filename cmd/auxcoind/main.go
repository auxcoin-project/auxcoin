package main

import (
	"log"
	"net"
	"os"

	"github.com/auxcoin-project/auxcoin/pb"
	"github.com/jawher/mow.cli"
	"google.golang.org/grpc"
)

type config struct {
	port *string
}

func main() {
	app := cli.App("auxcoind", "auxcoin blockchain service")

	app.Action = action(config{
		port: app.String(cli.StringOpt{
			Name:  "p port",
			Value: ":9999",
			Desc:  "listen port",
		}),
	})

	app.Run(os.Args)
}

func action(cfg config) func() {
	return func() {
		auxd := New()
		gsrv := grpc.NewServer()
		pb.RegisterAuxcoinServer(gsrv, auxd)

		l, err := net.Listen("tcp", *cfg.port)
		if err != nil {
			log.Fatal("failed to listen", err)
		}

		gsrv.Serve(l)
	}
}
