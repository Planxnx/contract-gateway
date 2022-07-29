# Contract Gateway

EVM-Based Smart-contract caller tools (my research project about design patterns by creating this tools)

#### TODO

- [ ] wrap ethclient with circuit breaker and timeout
- [ ] support uniswapv2 pair, factory contract

### Example

```go
client, err := ethclient.DialContext(ctx, "https://bsc-dataseed1.binance.org/")
if err != nil {
	panic(err)
}

// create contract gateway instance
contract := contractgateway.New(client)

bnbAddress := common.HexToAddress("0xbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c")
busdAddress := common.HexToAddress("0xe9e7cea3dedca5984780bafc599bd69add087d56")


// new session for each address
bnbContract := contract.WithAddress(bnbAddress)
busdContract := contract.WithAddress(busdAddress)


var bnb erc20.ERC20
if err := bnbContract.Symbol().Decimals().Find(&bnb).Error; err != nil {
	panic(err)
}
fmt.Printf("Data: %+v\n", bnb)

// new session with context
tx := bnbContract.WithContext(ctx)
if err := tx.BalanceOf(bnbAddress).Find(&bnb).Error; err != nil {
	panic(err)
}
fmt.Printf("BalanceOf: %+v\n", bnb.BalanceOf)


var busd erc20.ERC20
if err := busdContract.WithBlockNumber(12345678).BalanceOf(busdContract).Find(&busd).Error; err != nil {
	panic(err)
}
fmt.Printf("Data: %+v\n", busd)

```
