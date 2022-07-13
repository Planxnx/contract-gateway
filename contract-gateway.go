package contractgateway

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type ContractGateway struct {
	ethCaller bind.ContractCaller
}

func New(ethCaller bind.ContractCaller) *ContractGateway {
	return &ContractGateway{
		ethCaller: ethCaller,
	}
}
