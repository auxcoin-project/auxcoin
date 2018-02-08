package main

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProof(t *testing.T) {
	t.Parallel()

	p := Proof{}
	t.Run("hash block", func(t *testing.T) {
		b := &Block{
			Time:     1580601600,
			Bits:     8,
			PrevHash: []byte("prevHash"),
			Data:     []byte("data"),
		}

		err := p.Hash(b)
		require.NoError(t, err)
		require.Equal(t, uint32(75), b.Nonce)
		require.Equal(t,
			"009dca10fa3564afbd05ee487b43a6f69da3502c8c7f31f50e8accdcee935398",
			hex.EncodeToString(b.Hash),
		)
	})

	t.Run("verify block", func(t *testing.T) {
		t.Parallel()

		t.Run("pass", func(t *testing.T) {
			b := &Block{
				Time:     1580601600,
				Bits:     8,
				Nonce:    75,
				Hash:     []byte("hash"),
				PrevHash: []byte("prevHash"),
				Data:     []byte("data"),
			}
			require.True(t, p.Verify(b))
		})

		t.Run("fail", func(t *testing.T) {
			b := &Block{
				Time:     1580601600,
				Bits:     8,
				Nonce:    1,
				Hash:     []byte("hash"),
				PrevHash: []byte("prevHash"),
				Data:     []byte("data"),
			}
			require.False(t, p.Verify(b))
		})
	})
}
