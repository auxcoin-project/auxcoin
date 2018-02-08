package blockchain

import (
	"log"

	bolt "github.com/coreos/bbolt"
	"github.com/pkg/errors"
)

var (
	dbPath      = "blockchain.db"
	blockBucket = []byte("blocks")
	headKey     = []byte("l")
)

type (
	Chain struct {
		DB   *bolt.DB
		Head []byte
	}

	ChainIterator struct {
		DB      *bolt.DB
		Current []byte
	}
)

func NewChain(db *bolt.DB) *Chain {
	c := &Chain{DB: db}

	// ensure buckets exist
	db.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists(blockBucket); err != nil {
			log.Fatal(errors.Wrap(err, "failed to create bucket for blocks"))
		}
		return nil
	})

	db.View(func(tx *bolt.Tx) error {
		c.Head = tx.Bucket(blockBucket).Get(headKey)
		return nil
	})

	return c
}

func (c *Chain) Add(b *Block) error {
	encoded, err := b.Encode()
	if err != nil {
		return err
	}

	err = c.DB.Update(func(tx *bolt.Tx) error {
		if err := tx.Bucket(blockBucket).Put(b.Hash, encoded); err != nil {
			return errors.Wrap(err, "failed to persist block")
		}
		if err := tx.Bucket(blockBucket).Put(headKey, b.Hash); err != nil {
			return errors.Wrap(err, "failed to update blockchain head")
		}
		return nil
	})
	if err != nil {
		return err
	}

	c.Head = b.Hash

	return nil
}

func (c *Chain) Iterator() *ChainIterator {
	return &ChainIterator{c.DB, c.Head}
}

func (ci *ChainIterator) Next() (*Block, error) {
	var encoded []byte
	ci.DB.View(func(tx *bolt.Tx) error {
		encoded = tx.Bucket([]byte(blockBucket)).Get(ci.Current)
		return nil
	})
	if encoded == nil {
		return nil, nil
	}

	b, err := DecodeBlock(encoded)
	if err != nil {
		return nil, err
	}

	ci.Current = b.PrevHash

	return b, nil
}
