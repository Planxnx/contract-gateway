package erc20

import (
	"context"

	"github.com/Planxnx/contract-gateway/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

type erc20Builder func(context.Context, *contracts.ERC20Caller) error

type ERC20Gateway struct {
	ethCaller bind.ContractCaller
	address   common.Address
	erc20     *ERC20
	builders  []erc20Builder
	ctx       context.Context
}

func New(ethCaller bind.ContractCaller) *ERC20Gateway {
	return &ERC20Gateway{
		ethCaller: ethCaller,
		erc20:     &ERC20{},
		builders:  make([]erc20Builder, 0, 1),
	}
}

func (erc20Query *ERC20Gateway) WithContext(ctx context.Context) *ERC20Gateway {
	erc20Query.ctx = ctx
	return erc20Query
}

func (erc20Query *ERC20Gateway) WithAddress(address common.Address) *ERC20Gateway {
	erc20Query.address = address
	return erc20Query
}

func (g *ERC20Gateway) Find() (*ERC20, error) {
	erc20Contract, err := contracts.NewERC20Caller(g.address, g.ethCaller)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if g.ctx == nil {
		g.ctx = context.Background()
	}

	for _, b := range g.builders {
		if err := b(g.ctx, erc20Contract); err != nil {
			return nil, errors.WithStack(err)
		}
	}

	return g.erc20, nil
}

func (erc20Query *ERC20Gateway) Name() *ERC20Gateway {
	erc20Query.builders = append(erc20Query.builders, func(ctx context.Context, caller *contracts.ERC20Caller) error {
		val, err := caller.Name(&bind.CallOpts{Context: ctx})
		if err != nil {
			return errors.WithStack(err)
		}
		erc20Query.erc20.Name = val
		return nil
	})
	return erc20Query
}

func (erc20Query *ERC20Gateway) Decimals() *ERC20Gateway {
	erc20Query.builders = append(erc20Query.builders, func(ctx context.Context, caller *contracts.ERC20Caller) error {
		val, err := caller.Decimals(&bind.CallOpts{Context: ctx})
		if err != nil {
			return errors.WithStack(err)
		}
		erc20Query.erc20.Decimals = val
		return nil
	})
	return erc20Query
}

func (erc20Query *ERC20Gateway) Symbol() *ERC20Gateway {
	erc20Query.builders = append(erc20Query.builders, func(ctx context.Context, caller *contracts.ERC20Caller) error {
		val, err := caller.Symbol(&bind.CallOpts{Context: ctx})
		if err != nil {
			return errors.WithStack(err)
		}
		erc20Query.erc20.Symbol = val
		return nil
	})
	return erc20Query
}
