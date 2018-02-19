package main

import (
	"log"
	"net"
	"os"
	"time"

	bolt "github.com/coreos/bbolt"
	"github.com/jawher/mow.cli"
	"github.com/pkg/errors"
	"google.golang.org/grpc"

	auxbc "github.com/auxcoin-project/auxcoin/blockchain"
	auxpb "github.com/auxcoin-project/auxcoin/pb"
)

type config struct {
	port   *string
	dbPath *string
}

func main() {
	app := cli.App("auxcoind", "auxcoin blockchain service")

	app.Action = action(config{
		port: app.String(cli.StringOpt{
			Name:  "p port",
			Value: ":9999",
			Desc:  "listen port",
		}),
		dbPath: app.String(cli.StringOpt{
			Name:  "d db",
			Value: "blockchain.db",
			Desc:  "path to database",
		}),
	})

	app.Run(os.Args)
}

func action(cfg config) func() {
	return func() {
		db, err := bolt.Open(*cfg.dbPath, 0600, &bolt.Options{
			Timeout: 1 * time.Second,
		})
		if err != nil {
			log.Fatal(errors.Wrap(err, "failed to open database"))
		}

		gsrv := grpc.NewServer()
		auxd := newAuxcoind(auxbc.NewChain(db), 8, 25)

		auxpb.RegisterAuxcoinServer(gsrv, auxd)

		l, err := net.Listen("tcp", *cfg.port)
		if err != nil {
			log.Fatal("failed to listen", err)
		}

		gsrv.Serve(l)
	}
}
