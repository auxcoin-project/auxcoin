package blockchain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBlock(t *testing.T) {
	t.Parallel()

	t.Run("new block", func(t *testing.T) {
		expected := &Block{
			Timestamp: time.Now().Unix(),
			Bits:      42,
			Data:      []byte("data"),
		}
		actual := NewBlock(expected.Timestamp, expected.Bits, expected.Data)
		assert.Equal(t, expected, actual)
	})

	t.Run("encoding", func(t *testing.T) {
		block := NewBlock(time.Now().Unix(), 42, []byte("data"))

		encoded, err := block.Encode()
		assert.NoError(t, err)
		assert.NotEmpty(t, encoded)

		decoded, err := DecodeBlock(encoded)
		assert.NoError(t, err)
		assert.Equal(t, block, decoded)
	})
}
