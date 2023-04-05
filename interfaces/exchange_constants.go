package exchange

// BLOCKCHAIN NETWORKS
const (
	BSC_NETWORK = "BSC"
)

// ORDERS
const (
	ORDER_STATUS_FILLED = "FILLED"
	ORDER_TYPE_MARKET   = "market"
	ORDER_TYPE_LIMIT    = "limit"
	ORDER_SIDE_BUY      = "buy"
	ORDER_SIDE_SELL     = "sell"
)

// MARKETS
const (
	MARKET_TYPE_SPOT = "spot"
	MARKET_TYPE_PERP = "perp"
)

// WS
const (
	UNDEFINED = iota
	ERROR
	TICKER
	TRADES
	ORDERBOOK
	ORDERS_WS
	FILLS
)
