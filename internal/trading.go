package internal

//Pair represent an asset pairing that can be trading on an echange.
type Pair struct{
	Base Asset
	Quote Asset
}

var(
	//BITCUSD pair represents the BTC/USD pair typically found as a basis 
	BITCUSD = Pair{
		Base: BTC,
		Quote: USD,

	}

	//ETHUSD pair represent the ETH/USD pair found on exchanges.
	ETHUSD = Pair{
		Base: ETH,
		Quote: USD,
	}
)