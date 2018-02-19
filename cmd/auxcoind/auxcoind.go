package main

import (
	"context"
	"encoding/hex"

	auxbc "github.com/auxcoin-project/auxcoin/blockchain"
	auxpb "github.com/auxcoin-project/auxcoin/pb"
	"github.com/pkg/errors"
)

type auxcoind struct {
	chain  *auxbc.Chain
	bits   uint32
	reward uint32
}

func New(chain *auxbc.Chain, bits uint32, reward uint32) *auxcoind {
	return &auxcoind{chain, bits, reward}
}

func (a *auxcoind) Status(ctx context.Context, req *auxpb.StatusRequest) (*auxpb.StatusResponse, error) {
	return &auxpb.StatusResponse{
		Head:   hex.EncodeToString(a.chain.Head),
		Bits:   a.bits,
		Reward: a.reward,
	}, nil
}

func (a *auxcoind) AddBlock(ctx context.Context, req *auxpb.AddBlockRequest) (*auxpb.AddBlockResponse, error) {
	b, err := auxbc.DecodeBlock([]byte(req.GetBlock()))
	if err != nil {
		err = errors.Wrap(err, "failed to decode block")
		return &auxpb.AddBlockResponse{err.Error()}, err
	}

	if err := a.chain.Add(b); err != nil {
		err = errors.Wrap(err, "failed to add block")
		return &auxpb.AddBlockResponse{err.Error()}, err
	}

	return &auxpb.AddBlockResponse{""}, nil
}
