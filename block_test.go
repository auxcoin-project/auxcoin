package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func newTestBlock() *Block {
	return NewBlock(1580601600, 8, []*Transaction{
		newTestCoinbaseTransaction(),
		newTestTransaction(),
	})
}

func TestBlock(t *testing.T) {
	t.Parallel()

	t.Run("encoding", func(t *testing.T) {
		block := newTestBlock()

		encoded, err := block.Encode()
		require.NoError(t, err)
		require.NotEmpty(t, encoded)

		decoded, err := DecodeBlock(encoded)
		require.NoError(t, err)
		require.Equal(t, block, decoded)
	})
}
