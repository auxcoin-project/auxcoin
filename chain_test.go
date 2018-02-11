package main

import (
	"os"
	"testing"
	"time"

	bolt "github.com/coreos/bbolt"
	"github.com/stretchr/testify/require"
)

func openTestDB(t *testing.T, path string) (*bolt.DB, func()) {
	db, err := bolt.Open(path, 0600, &bolt.Options{
		Timeout: 1 * time.Second,
	})
	require.NoError(t, err)

	return db, func() {
		db.Close()
		os.RemoveAll(path)
	}
}

func newTestChain(t *testing.T, db *bolt.DB) *Chain {
	bc := NewChain(db)

	var p Proof
	for i := 0; i < 3; i++ {
		b := newTestBlock()
		b.PrevHash = bc.Head
		b.Txns = []*Transaction{
			newTestCoinbaseTransaction(),
			newTestTransaction(),
		}
		p.Hash(b)

		err := bc.Add(b)
		require.NoError(t, err)
	}

	return bc
}

func TestChain(t *testing.T) {
	t.Parallel()

	t.Run("iterate", func(t *testing.T) {
		db, cleanup := openTestDB(t, "blockchain-test.db")
		defer cleanup()

		bc := newTestChain(t, db)
		iter := bc.Iterator()

		for i := 2; i >= 0; i-- {
			b, err := iter.Next()
			require.NoError(t, err)
			require.NotEmpty(t, b)
		}
		b, err := iter.Next()
		require.NoError(t, err)
		require.Empty(t, b)
	})
}
