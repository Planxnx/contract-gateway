package erc20

import (
	"context"

	"github.com/Planxnx/contract-gateway/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

type ERC20Gateway struct {
	ethCaller bind.ContractCaller
	statement *statement
	Error     error
}

func New(ethCaller bind.ContractCaller) *ERC20Gateway {
	return &ERC20Gateway{
		ethCaller: ethCaller,
		statement: &statement{},
	}
}

func (g *ERC20Gateway) getInstance() *ERC20Gateway {
	if g.statement == nil {
		// it's should create a new session or not?
		g.statement = &statement{}
	}

	if g.statement.ctx == nil {
		g.statement.ctx = context.Background()
	}

	return g
}

func (g *ERC20Gateway) Session() *ERC20Gateway {
	session := &ERC20Gateway{
		ethCaller: g.ethCaller,
		statement: g.statement.clone(),
	}
	return session
}

func (g *ERC20Gateway) WithAddress(address common.Address) *ERC20Gateway {
	newGateway := g.Session()
	newGateway.statement.address = address
	return newGateway
}

func (g *ERC20Gateway) WithContext(ctx context.Context) *ERC20Gateway {
	newGateway := g.Session()
	newGateway.statement.ctx = ctx
	return newGateway
}

func (g *ERC20Gateway) AddError(err error) error {
	if g.Error == nil {
		g.Error = err
	} else if g.Error != nil {
		g.Error = errors.Wrap(g.Error, err.Error())
	}
	return g.Error
}

func (g *ERC20Gateway) Find(result *ERC20) (tx *ERC20Gateway) {
	tx = g.getInstance()

	erc20Contract, err := contracts.NewERC20Caller(g.statement.address, g.ethCaller)
	if err != nil {
		tx.AddError(errors.WithStack(err))
		return
	}

	if g.statement.ctx == nil {
		g.statement.ctx = context.Background()
	}

	for _, builder := range g.statement.builders {
		if err := builder(g.statement.ctx, erc20Contract, result); err != nil {
			tx.AddError(errors.WithStack(err))
		}
	}

	return
}

func (g *ERC20Gateway) Name() (tx *ERC20Gateway) {
	tx = g.getInstance()
	tx.statement.builders = append(tx.statement.builders, func(ctx context.Context, caller *contracts.ERC20Caller, erc20 *ERC20) error {
		val, err := caller.Name(&bind.CallOpts{Context: ctx})
		if err != nil {
			return errors.WithStack(err)
		}
		erc20.Name = val
		return nil
	})
	return tx
}

func (g *ERC20Gateway) Decimals() *ERC20Gateway {
	g.statement.builders = append(g.statement.builders, func(ctx context.Context, caller *contracts.ERC20Caller, erc20 *ERC20) error {
		val, err := caller.Decimals(&bind.CallOpts{Context: ctx})
		if err != nil {
			return errors.WithStack(err)
		}
		erc20.Decimals = val
		return nil
	})
	return g
}

func (g *ERC20Gateway) Symbol() *ERC20Gateway {
	g.statement.builders = append(g.statement.builders, func(ctx context.Context, caller *contracts.ERC20Caller, erc20 *ERC20) error {
		val, err := caller.Symbol(&bind.CallOpts{Context: ctx})
		if err != nil {
			return errors.WithStack(err)
		}
		erc20.Symbol = val
		return nil
	})
	return g
}
