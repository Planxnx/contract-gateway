package contracts

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/pkg/errors"
)

var (
	ABIERC20            abi.ABI = MustParseABI(ERC20ABI)
	ABIUniswapV2Factory abi.ABI = MustParseABI(UniswapV2FactoryABI)
	ABIUniswapV2Pair    abi.ABI = MustParseABI(UniswapV2PairABI)
)

func MustParseABI(strABI string) abi.ABI {
	abi, err := abi.JSON(strings.NewReader(strABI))
	if err != nil {
		panic(errors.Wrapf(err, "Can't parse contract abi"))
	}
	return abi
}
