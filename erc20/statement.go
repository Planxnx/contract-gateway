package erc20

import (
	"context"

	"github.com/Planxnx/contract-gateway/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type erc20Builder func(*bind.CallOpts, *contracts.ERC20Caller, *ERC20) error

type statement struct {
	ctx      context.Context
	address  common.Address
	builders []erc20Builder
}

func (s *statement) clone() *statement {

	newS := &statement{
		ctx:     s.ctx,
		address: s.address,
	}

	if len(s.builders) > 0 {
		newS.builders = make([]erc20Builder, len(s.builders))
		copy(newS.builders, s.builders)
	}

	return newS
}
