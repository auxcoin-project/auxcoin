package blockchain

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
		require.Equal(t, uint32(773), b.Nonce)
		require.Equal(t,
			"00ff8fea864bb4b68e65f9d197d1476161345f29a6c60df29f3c1b5d9a1a7d31",
			hex.EncodeToString(b.Hash),
		)
	})

	t.Run("verify block", func(t *testing.T) {
		t.Parallel()

		t.Run("pass", func(t *testing.T) {
			b := newTestBlock()
			b.PrevHash, _ = hex.DecodeString("00c12b7b65fc30c29648a5eda0acb51a5d9cc754d860a8b4579a5b9f8f7c8896")
			b.Hash, _ = hex.DecodeString("00ff8fea864bb4b68e65f9d197d1476161345f29a6c60df29f3c1b5d9a1a7d31")
			b.Nonce = 773

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
