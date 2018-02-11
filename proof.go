package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"math/big"
)

type Proof struct{}

func (p Proof) prepare(b *Block) []byte {
	var bx []byte
	buf := bytes.NewBuffer(bx)

	binary.Write(buf, binary.LittleEndian, b.Time)
	binary.Write(buf, binary.LittleEndian, b.Bits)
	buf.Write(b.PrevHash)
	for _, txn := range b.Txns {
		buf.Write(txn.Bytes())
	}
	binary.Write(buf, binary.LittleEndian, b.Nonce)

	return buf.Bytes()
}

func (p Proof) bitsToTarget(bits uint32) *big.Int {
	target := big.NewInt(1)
	return target.Lsh(target, uint(256-bits))
}

func (p Proof) verify(hash []byte, target *big.Int) bool {
	var i big.Int
	i.SetBytes(hash)
	return i.Cmp(target) == -1
}

func (p Proof) Hash(b *Block) error {
	b.Nonce = 0
	target := p.bitsToTarget(b.Bits)

	for {
		data := p.prepare(b)
		hash := sha256.Sum256(data)
		b.Hash = hash[:]
		if p.verify(b.Hash, target) {
			break
		}
		b.Nonce++
	}

	return nil
}

func (p Proof) Verify(b *Block) bool {
	data := p.prepare(b)
	hash := sha256.Sum256(data)
	target := p.bitsToTarget(b.Bits)

	return p.verify(hash[:], target)
}
