package contracts

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestABIEvent(t *testing.T) {
	t.Run("ERC20", func(t *testing.T) {
		assert := assert.New(t)
		{
			actual := ABIERC20.Events["Transfer"].ID
			expected := common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
			assert.Equal(expected, actual)
		}
		{
			actual := ABIERC20.Events["Approval"].ID
			expected := common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925")
			assert.Equal(expected, actual)
		}
	})
	t.Run("UniswapV2_Factory", func(t *testing.T) {
		assert := assert.New(t)
		actual := ABIUniswapV2Factory.Events["PairCreated"].ID
		expected := common.HexToHash("0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9")
		assert.Equal(expected, actual)
	})
	t.Run("UniswapV2_Pair", func(t *testing.T) {
		assert := assert.New(t)
		{
			actual := ABIUniswapV2Pair.Events["Swap"].ID
			expected := common.HexToHash("0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822")
			assert.Equal(expected, actual)
		}
		{
			actual := ABIUniswapV2Pair.Events["Sync"].ID
			expected := common.HexToHash("0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1")
			assert.Equal(expected, actual)
		}
		{
			actual := ABIUniswapV2Pair.Events["Mint"].ID
			expected := common.HexToHash("0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f")
			assert.Equal(expected, actual)
		}
		{
			actual := ABIUniswapV2Pair.Events["Burn"].ID
			expected := common.HexToHash("0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496")
			assert.Equal(expected, actual)
		}
		{
			actual := ABIUniswapV2Pair.Events["Transfer"].ID
			expected := common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
			assert.Equal(expected, actual)
		}
		{
			actual := ABIUniswapV2Pair.Events["Approval"].ID
			expected := common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925")
			assert.Equal(expected, actual)
		}
	})
}

func TestABIUnpack(t *testing.T) {
	t.Run("ERC20", func(t *testing.T) {
		assert := assert.New(t)
		{
			actualEvent := &ERC20Transfer{}
			err := ABIERC20.UnpackIntoInterface(actualEvent, "Transfer", []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1})
			assert.NoError(err)
			assert.Equal(big.NewInt(1), actualEvent.Value)
		}
		{
			actualEvent := &ERC20Approval{}
			err := ABIERC20.UnpackIntoInterface(actualEvent, "Approval", []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255})
			assert.NoError(err)
			assert.Equal(big.NewInt(255), actualEvent.Value)
		}
	})
	t.Run("UniswapV2_Factory", func(t *testing.T) {
		assert := assert.New(t)
		actualEvent := &UniswapV2FactoryPairCreated{}
		data, _ := hex.DecodeString("000000000000000000000000392fd6554747989fcfc2c5ff45f13ac5392b209a00000000000000000000000000000000000000000000000000000000000129ec")
		err := ABIUniswapV2Factory.UnpackIntoInterface(actualEvent, "PairCreated", data)

		assert.NoError(err)
		assert.Equal(common.HexToAddress("0x392fd6554747989fcfc2c5ff45f13ac5392b209a"), actualEvent.Pair)
	})
	t.Run("UniswapV2_Pair", func(t *testing.T) {
		assert := assert.New(t)
		{
			actualEvent := &UniswapV2PairSwap{}
			data, _ := hex.DecodeString("000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000795c552f000000000000000000000000000000000000000000000000000c44683e256b4b0000000000000000000000000000000000000000000000000000000000000000")
			err := ABIUniswapV2Pair.UnpackIntoInterface(actualEvent, "Swap", data)

			assert.NoError(err)
			assert.Equal(big.NewInt(3452914230455115), actualEvent.Amount0Out)
			assert.Equal(big.NewInt(2036094255), actualEvent.Amount1In)
		}
		{
			actualEvent := &UniswapV2PairSync{}
			expectedReserve0, _ := new(big.Int).SetString("700369211779219078493", 10)
			expectedReserve1 := big.NewInt(353556955659508025)
			data, _ := hex.DecodeString("000000000000000000000000000000000000000000000025f7934765e8ae3d5d00000000000000000000000000000000000000000000000004e8162d8df41939")
			err := ABIUniswapV2Pair.UnpackIntoInterface(actualEvent, "Sync", data)

			assert.NoError(err)
			assert.Equal(expectedReserve0, actualEvent.Reserve0)
			assert.Equal(expectedReserve1, actualEvent.Reserve1)
		}
	})
}
