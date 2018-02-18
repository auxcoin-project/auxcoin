package main

import (
	"log"
	"net"
	"os"
	"time"

	auxbc "github.com/auxcoin-project/auxcoin/blockchain"
	auxpb "github.com/auxcoin-project/auxcoin/pb"
	bolt "github.com/coreos/bbolt"
	"github.com/jawher/mow.cli"
	"github.com/pkg/errors"
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
		db, err := bolt.Open("blockchain.db", 0600, &bolt.Options{
			Timeout: 1 * time.Second,
		})
		if err != nil {
			log.Fatal(errors.Wrap(err, "failed to open database"))
		}

		gsrv := grpc.NewServer()
		auxd := New(auxbc.NewChain(db))

		auxpb.RegisterAuxcoinServer(gsrv, auxd)

		l, err := net.Listen("tcp", *cfg.port)
		if err != nil {
			log.Fatal("failed to listen", err)
		}

		gsrv.Serve(l)
	}
}
