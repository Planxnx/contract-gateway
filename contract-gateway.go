package contractgateway

import (
	"github.com/Planxnx/contract-gateway/erc20"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type ContractGateway struct {
	ethCaller bind.ContractCaller
	*erc20.ERC20Gateway
}

func New(ethCaller bind.ContractCaller) *ContractGateway {
	return &ContractGateway{
		ethCaller:    ethCaller,
		ERC20Gateway: erc20.New(ethCaller),
	}
}
