package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProof(t *testing.T) {
	t.Parallel()

	t.Run("hash block", func(t *testing.T) {
		b := &Block{
			Timestamp: 1580601600,
			Bits:      8,
			PrevHash:  []byte("prevHash"),
			Data:      []byte("data"),
		}

		err := (Proof{}).HashBlock(b)
		require.NoError(t, err)
		require.Equal(t, uint32(75), b.Nonce)
		require.Equal(t, []byte{0x0, 0x9d, 0xca, 0x10, 0xfa, 0x35, 0x64,
			0xaf, 0xbd, 0x5, 0xee, 0x48, 0x7b, 0x43, 0xa6, 0xf6, 0x9d, 0xa3,
			0x50, 0x2c, 0x8c, 0x7f, 0x31, 0xf5, 0xe, 0x8a, 0xcc, 0xdc, 0xee,
			0x93, 0x53, 0x98}, b.Hash)
	})

	t.Run("verify block", func(t *testing.T) {
		t.Parallel()

		t.Run("pass", func(t *testing.T) {
			b := &Block{
				Timestamp: 1580601600,
				Bits:      8,
				Nonce:     75,
				Hash:      []byte("hash"),
				PrevHash:  []byte("prevHash"),
				Data:      []byte("data"),
			}
			require.True(t, (Proof{}).VerifyBlock(b))
		})

		t.Run("fail", func(t *testing.T) {
			b := &Block{
				Timestamp: 1580601600,
				Bits:      8,
				Nonce:     1,
				Hash:      []byte("hash"),
				PrevHash:  []byte("prevHash"),
				Data:      []byte("data"),
			}
			require.False(t, (Proof{}).VerifyBlock(b))
		})
	})
}
