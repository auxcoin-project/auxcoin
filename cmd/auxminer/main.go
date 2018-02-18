package main

import (
	"context"
	"log"
	"os"
	"time"

	auxbc "github.com/auxcoin-project/auxcoin/blockchain"
	auxpb "github.com/auxcoin-project/auxcoin/pb"
	"github.com/jawher/mow.cli"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type config struct {
	srvAddr  *string
	coinAddr *string
}

const reward = 25
const bits = 16

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

func newClient(cfg config) auxpb.AuxcoinClient {
	c, err := grpc.Dial(*cfg.srvAddr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return auxpb.NewAuxcoinClient(c)
}

func action(cfg config) func() {
	return func() {
		for {
			var txns []*auxbc.Transaction
			txns = append(txns, coinbaseTxn(reward, *cfg.coinAddr))

			b := auxbc.NewBlock(time.Now().Unix(), bits, txns)

			var p auxbc.Proof
			if err := p.Hash(b); err != nil {
				log.Print(errors.Wrap(err, "failed to hash block"))
				continue
			}

			enc, err := b.Encode()
			if err != nil {
				log.Println(errors.Wrap(err, "failed to encode block"))
				continue
			}

			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

			c := newClient(cfg)
			resp, err := c.AddBlock(ctx, &auxpb.AddBlockRequest{string(enc)})
			if err != nil {
				log.Println(errors.Wrap(err, "failed to add block"))
			}
			if resp.Error != "" {
				log.Println(resp.Error)
			} else {
				log.Print(".")
			}
			cancel()
		}
	}
}

func coinbaseTxn(reward uint32, addr string) *auxbc.Transaction {
	cbTxn := auxbc.NewTransaction()
	cbTxn.TxnOut = append(cbTxn.TxnOut, auxbc.NewTxnOut(reward, addr))

	return cbTxn
}
