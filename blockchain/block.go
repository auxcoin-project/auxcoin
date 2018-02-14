package blockchain

import (
	"bytes"
	"encoding/gob"

	"github.com/pkg/errors"
)

type (
	Block struct {
		Time     int64
		Bits     uint32
		Nonce    uint32
		Hash     []byte
		PrevHash []byte
		Txns     []*Transaction
	}
)

func NewBlock(tstamp int64, bits uint32, txns []*Transaction) *Block {
	return &Block{
		Time: tstamp,
		Bits: bits,
		Txns: txns,
	}
}

func (b Block) Encode() ([]byte, error) {
	var result bytes.Buffer
	if err := gob.NewEncoder(&result).Encode(b); err != nil {
		return nil, errors.Wrap(err, "failed to encode block")
	}
	return result.Bytes(), nil
}

func DecodeBlock(encoded []byte) (*Block, error) {
	var b Block
	buf := bytes.NewBuffer(encoded)
	if err := gob.NewDecoder(buf).Decode(&b); err != nil {
		return nil, errors.Wrap(err, "failed to decode block")
	}
	return &b, nil
}
