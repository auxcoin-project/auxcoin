package blockchain

import (
	"log"

	bolt "github.com/coreos/bbolt"
	"github.com/pkg/errors"
)

var (
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
	bc := &Chain{DB: db}

	// ensure buckets exist
	db.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists(blockBucket); err != nil {
			log.Fatal(errors.Wrap(err, "failed to create bucket for blocks"))
		}
		return nil
	})

	db.View(func(tx *bolt.Tx) error {
		bc.Head = tx.Bucket(blockBucket).Get(headKey)
		return nil
	})

	return bc
}

func (bc *Chain) Add(b *Block) error {
	encoded, err := b.Encode()
	if err != nil {
		return err
	}

	err = bc.DB.Update(func(tx *bolt.Tx) error {
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

	bc.Head = b.Hash

	return nil
}

func (bc *Chain) Iterator() *ChainIterator {
	return &ChainIterator{bc.DB, bc.Head}
}

func (bci *ChainIterator) Next() (*Block, error) {
	var encoded []byte
	bci.DB.View(func(tx *bolt.Tx) error {
		encoded = tx.Bucket([]byte(blockBucket)).Get(bci.Current)
		return nil
	})
	if encoded == nil {
		return nil, nil
	}

	b, err := DecodeBlock(encoded)
	if err != nil {
		return nil, err
	}

	bci.Current = b.PrevHash

	return b, nil
}
