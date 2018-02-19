package main

import (
	"context"
	"encoding/hex"

	"github.com/pkg/errors"

	auxbc "github.com/auxcoin-project/auxcoin/blockchain"
	auxpb "github.com/auxcoin-project/auxcoin/pb"
)

type auxcoind struct {
	chain  *auxbc.Chain
	bits   uint32
	reward uint32
}

func newAuxcoind(chain *auxbc.Chain, bits uint32, reward uint32) *auxcoind {
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

	var p auxbc.Proof
	if ok := p.Verify(b); !ok {
		err = errors.New("block is not valid")
		return &auxpb.AddBlockResponse{err.Error()}, err
	}

	if err := a.chain.Add(b); err != nil {
		err = errors.Wrap(err, "failed to add block")
		return &auxpb.AddBlockResponse{err.Error()}, err
	}

	return &auxpb.AddBlockResponse{""}, nil
}
