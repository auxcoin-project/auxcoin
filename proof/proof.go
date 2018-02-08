package proof

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"math/big"

	"github.com/auxcoin-project/auxcoin/blockchain"
)

func prepare(b *blockchain.Block) []byte {
	var bx []byte
	buf := bytes.NewBuffer(bx)

	binary.Write(buf, binary.LittleEndian, b.Timestamp)
	binary.Write(buf, binary.LittleEndian, b.Bits)
	buf.Write(b.PrevHash)
	buf.Write(b.Data)
	binary.Write(buf, binary.LittleEndian, b.Nonce)

	return buf.Bytes()
}

func bitsToTarget(bits uint32) *big.Int {
	target := big.NewInt(1)
	return target.Lsh(target, uint(256-bits))
}

func verify(hash []byte, target *big.Int) bool {
	var i big.Int
	i.SetBytes(hash)
	return i.Cmp(target) == -1
}

func Hash(b *blockchain.Block) error {
	b.Nonce = 0
	target := bitsToTarget(b.Bits)

	for {
		data := prepare(b)
		hash := sha256.Sum256(data)
		b.Hash = hash[:]
		if verify(b.Hash, target) {
			break
		}
		b.Nonce++
	}

	return nil
}

func Verify(b *blockchain.Block) bool {
	data := prepare(b)
	hash := sha256.Sum256(data)
	target := bitsToTarget(b.Bits)

	return verify(hash[:], target)
}
