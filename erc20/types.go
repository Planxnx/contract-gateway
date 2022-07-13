package erc20

import "github.com/holiman/uint256"

type ERC20 struct {
	// static fields
	Name     string
	Decimals uint8
	Symbol   string

	// dynamic fields
	TotalSupply *uint256.Int
	BalanceOf   *uint256.Int
	Allowance   *uint256.Int
}
