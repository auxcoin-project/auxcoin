package main

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProof(t *testing.T) {
	t.Parallel()

	var p Proof
	t.Run("hash block", func(t *testing.T) {
		b := newTestBlock()
		b.PrevHash, _ = hex.DecodeString("00c12b7b65fc30c29648a5eda0acb51a5d9cc754d860a8b4579a5b9f8f7c8896")

		err := p.Hash(b)
		require.NoError(t, err)
		require.Equal(t, uint32(7), b.Nonce)
		require.Equal(t,
			"00c2d161c5b31b8f1bacd6fff5fa31add5f16eec66b33450820897585ac80594",
			hex.EncodeToString(b.Hash),
		)
	})

	t.Run("verify block", func(t *testing.T) {
		t.Parallel()

		t.Run("pass", func(t *testing.T) {
			b := newTestBlock()
			b.PrevHash, _ = hex.DecodeString("00c12b7b65fc30c29648a5eda0acb51a5d9cc754d860a8b4579a5b9f8f7c8896")
			b.Hash, _ = hex.DecodeString("00c2d161c5b31b8f1bacd6fff5fa31add5f16eec66b33450820897585ac80594")
			b.Nonce = 7

			require.True(t, p.Verify(b))
		})

		t.Run("fail", func(t *testing.T) {
			b := newTestBlock()
			b.PrevHash, _ = hex.DecodeString("00c12b7b65fc30c29648a5eda0acb51a5d9cc754d860a8b4579a5b9f8f7c8896")
			b.Hash, _ = hex.DecodeString("0b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde")

			require.False(t, p.Verify(b))
		})
	})
}
