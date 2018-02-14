package blockchain

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newTestCoinbaseTransaction() *Transaction {
	txn := NewTransaction()
	txn.TxnOut = append(txn.TxnOut, NewTxnOut(reward, "ADDRESS_1"))
	return txn
}

func newTestTransaction() *Transaction {
	txn := NewTransaction()
	hash, _ := hex.DecodeString("0cce18c0fb799350e7f7e1cfda8994255cbd75513dd8cdbf4b33cfb8dafc702d")
	txn.TxnIn = append(txn.TxnIn, NewTxnIn(hash, 0, "ADDRESS_1"))
	txn.TxnOut = append(txn.TxnOut, NewTxnOut(reward, "ADDRESS_2"))
	return txn
}

func TestTransaction(t *testing.T) {
	t.Parallel()

	t.Run("transaction hash", func(t *testing.T) {
		txn := newTestTransaction()
		txn.Hash()
		assert.Equal(t,
			"46f69fc76abf1264f5a809b82e7887554bb2166f97c97a41353307e90ed42da7",
			hex.EncodeToString(txn.ID),
		)
	})
}
