package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"

	"github.com/jawher/mow.cli"
	"github.com/jbenet/go-base58"
	"golang.org/x/crypto/ripemd160"
)

func address(cfg config) func(*cli.Cmd) {
	return func(cmd *cli.Cmd) {
		cmd.Command("new", "new address", newAddress(cfg))
		cmd.Command("list ls", "list addresses", listAddress(cfg))
	}
}

func newAddress(cfg config) func(*cli.Cmd) {
	return func(cmd *cli.Cmd) {
		cmd.Action = func() {
			key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
			if err != nil {
				panic("failed to generate key")
			}
			pub := append(key.PublicKey.X.Bytes(), key.PublicKey.Y.Bytes()...)
			sum := sha256.Sum256(pub)

			ripe := ripemd160.New()
			if _, err := ripe.Write(sum[:]); err != nil {
				panic("failed to write sum")
			}
			addr := append([]byte{0x01}, ripe.Sum(nil)...)

			// checksum
			s := sha256.Sum256(addr)
			s = sha256.Sum256(s[:])
			addr = append(addr, s[:4]...)

			fmt.Println("Public: Ax" + base58.Encode(addr))
			fmt.Println("Private: " + base58.Encode(key.D.Bytes()))
		}
	}
}

func listAddress(cfg config) func(*cli.Cmd) {
	return func(cmd *cli.Cmd) {
		priv := cmd.BoolOpt("p priv", false, "show private")

		cmd.Action = func() {
			fmt.Println("list address not yet implemented")
			_ = priv
		}
	}
}
