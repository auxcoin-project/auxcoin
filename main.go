package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"time"

	bolt "github.com/coreos/bbolt"
	"github.com/jawher/mow.cli"

	"github.com/auxcoin-project/auxcoin/blockchain"
	"github.com/auxcoin-project/auxcoin/proof"
)

func main() {
	app := cli.App("auxcoin", "auxcoin blockchain")

	app.Command("add", "add an entry to the block chain", add)

	app.Command("view", "view the blockchain", view)

	app.Run(os.Args)
}

func openDB() *bolt.DB {
	db, err := bolt.Open("blockchain.db", 0600, &bolt.Options{
		Timeout: 1 * time.Second,
	})
	if err != nil {
		log.Fatal("unable to open database")
	}
	return db
}

func add(cmd *cli.Cmd) {
	data := cmd.StringArg("DATA", "", "data to add to chain")
	cmd.Action = func() {
		db := openDB()
		bc := blockchain.NewChain(db)

		b := blockchain.NewBlock(time.Now().Unix(), 8, []byte(*data))

		b.PrevHash = bc.Head
		proof.Hash(b)
		bc.Add(b)

		db.Close()
	}
}

func view(cmd *cli.Cmd) {
	cmd.Action = func() {
		db := openDB()
		bc := blockchain.NewChain(db)

		iter := bc.Iterator()

		for {
			b, _ := iter.Next()
			if b == nil {
				break
			}
			fmt.Printf("Timestamp: %v\n", b.Timestamp)
			fmt.Printf("Bits: %v\n", b.Bits)
			fmt.Printf("Nonce: %v\n", b.Nonce)
			fmt.Printf("Hash: %v\n", hex.EncodeToString(b.Hash))
			fmt.Printf("PrevHash: %v\n", hex.EncodeToString(b.PrevHash))
			fmt.Printf("Data: %v\n\n", string(b.Data))
		}
		db.Close()
	}
}
