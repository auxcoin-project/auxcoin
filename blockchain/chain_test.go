package blockchain

import (
	"os"
	"testing"
	"time"

	bolt "github.com/coreos/bbolt"
	"github.com/stretchr/testify/require"
)

func openDB(t *testing.T, path string) (*bolt.DB, func()) {
	db, err := bolt.Open(path, 0600, &bolt.Options{
		Timeout: 1 * time.Second,
	})
	require.NoError(t, err)

	return db, func() {
		db.Close()
		os.RemoveAll(path)
	}
}

func TestBlockChain(t *testing.T) {
	db, cleanup := openDB(t, "test-blockchain.db")
	defer cleanup()

	c := NewChain(db)

	// add blocks
	b1 := NewBlock(time.Now().Unix(), 42, []byte("data1"))
	b1.Hash = []byte("hash1")
	err := c.Add(b1)
	require.NoError(t, err)

	b2 := NewBlock(time.Now().Unix(), 42, []byte("data2"))
	b2.Hash = []byte("hash2")
	b2.PrevHash = b1.Hash
	err = c.Add(b2)
	require.NoError(t, err)

	b3 := NewBlock(time.Now().Unix(), 42, []byte("data3"))
	b3.Hash = []byte("hash3")
	b3.PrevHash = b2.Hash
	err = c.Add(b3)
	require.NoError(t, err)

	// iterate
	iter := c.Iterator()

	b, err := iter.Next()
	require.NoError(t, err)
	require.Equal(t, b3, b)

	b, err = iter.Next()
	require.NoError(t, err)
	require.Equal(t, b2, b)

	b, err = iter.Next()
	require.NoError(t, err)
	require.Equal(t, b1, b)

	b, err = iter.Next()
	require.NoError(t, err)
	require.Nil(t, b)
}
