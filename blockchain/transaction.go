package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
)

const reward uint32 = 25

type (
	Transaction struct {
		ID     []byte
		TxnIn  []*TxnIn
		TxnOut []*TxnOut
	}

	TxnIn struct {
		TxnID  []byte
		Index  int
		Script string
	}

	TxnOut struct {
		Value  uint32
		Script string
	}
)

func NewTransaction() *Transaction {
	return &Transaction{}
}

func (txn Transaction) Bytes() []byte {
	var b []byte
	buf := bytes.NewBuffer(b)

	for _, in := range txn.TxnIn {
		buf.Write(in.TxnID)
		binary.Write(buf, binary.LittleEndian, in.Index)
		buf.Write([]byte(in.Script))
	}

	for _, out := range txn.TxnOut {
		binary.Write(buf, binary.LittleEndian, out.Value)
		buf.Write([]byte(out.Script))
	}

	return buf.Bytes()
}

func (txn *Transaction) Hash() {
	hash := sha256.Sum256(txn.Bytes())
	txn.ID = hash[:]
}

func NewTxnIn(txnID []byte, out int, script string) *TxnIn {
	return &TxnIn{txnID, out, script}
}

func NewTxnOut(value uint32, script string) *TxnOut {
	return &TxnOut{value, script}
}
