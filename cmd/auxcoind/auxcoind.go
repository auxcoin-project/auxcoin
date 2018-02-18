package main

import (
	"context"

	auxbc "github.com/auxcoin-project/auxcoin/blockchain"
	auxpb "github.com/auxcoin-project/auxcoin/pb"
	"github.com/pkg/errors"
)

type auxcoind struct {
	bc *auxbc.Chain
}

func New(bc *auxbc.Chain) *auxcoind {
	return &auxcoind{bc}
}

func (a *auxcoind) AddBlock(ctx context.Context, req *auxpb.AddBlockRequest) (*auxpb.AddBlockResponse, error) {
	b, err := auxbc.DecodeBlock([]byte(req.GetBlock()))
	if err != nil {
		err = errors.Wrap(err, "failed to decode block")
		return &auxpb.AddBlockResponse{err.Error()}, err
	}

	if err := a.bc.Add(b); err != nil {
		err = errors.Wrap(err, "failed to add block")
		return &auxpb.AddBlockResponse{err.Error()}, err
	}

	return &auxpb.AddBlockResponse{""}, nil
}
