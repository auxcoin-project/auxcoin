package main

import (
	"context"
	"encoding/hex"
	"log"
	"os"
	"time"

	"github.com/jawher/mow.cli"
	"github.com/pkg/errors"
	"google.golang.org/grpc"

	auxbc "github.com/auxcoin-project/auxcoin/blockchain"
	auxpb "github.com/auxcoin-project/auxcoin/pb"
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

func newClient(cfg config) auxpb.AuxcoinClient {
	c, err := grpc.Dial(*cfg.srvAddr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return auxpb.NewAuxcoinClient(c)
}

func action(cfg config) func() {
	return func() {
		c := newClient(cfg)
		for {
			status, err := c.Status(context.Background(), &auxpb.StatusRequest{})
			if err != nil {
				panic(errors.Wrap(err, "failed to fetch status"))
			}

			var txns []*auxbc.Transaction
			txns = append(txns, coinbaseTxn(status.Reward, *cfg.coinAddr))

			b := auxbc.NewBlock(time.Now().Unix(), status.Bits, txns)
			b.PrevHash, err = hex.DecodeString(status.Head)
			if err != nil {
				panic(errors.Wrap(err, "failed to decode head"))
			}

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

			resp, err := c.AddBlock(context.Background(), &auxpb.AddBlockRequest{string(enc)})
			if err != nil {
				log.Println(errors.Wrap(err, "failed to add block"))
			}
			if resp.Error != "" {
				log.Println(resp.Error)
			}
		}
	}
}

func coinbaseTxn(reward uint32, addr string) *auxbc.Transaction {
	cbTxn := auxbc.NewTransaction()
	cbTxn.TxnOut = append(cbTxn.TxnOut, auxbc.NewTxnOut(reward, addr))

	return cbTxn
}
