package bybit_exchange

const (
	TESTNET_URL   = "https://api-testnet.bybit.com"
	MAINNET_URL   = "https://api.bybit.com"
	MAINNET_URL_2 = "https://api.bytick.com"
)

// API ENDPOINTS
const (
	TICKER                    = "/v2/public/tickers"
	SERVER_TIME               = "/v2/public/time"
	GET_PERP_SYMBOLS_INFO     = "/v2/public/symbols"
	PLACE_PERP_ORDER          = "/v2/private/order/create"
	GET_PERP_ORDER            = "/v2/private/order/list"
	GET_PERP_POSITION         = "/v2/private/position/list"
	SPOT_MARKET_PRICE         = "/spot/quote/v1/ticker/book_ticker"
	SPOT_ORDER                = "/spot/v1/order"
	GET_SPOT_BALANCE          = "/spot/v1/account"
	GET_PERP_BALANCE          = "/v2/private/wallet/balance"
	GET_DEPOSIT_ADDRESS       = "/asset/v1/private/deposit/address"
	CANCEL_PERP_ORDER         = "/v2/private/order/cancel"
	WITHDRAW_FROM_SPOT_WALLET = "/asset/v1/private/withdraw"
)

// ORDERS
const (
	ORDER_STATUS_FILLED = "FILLED"
	ORDER_SIDE_BUY      = "Buy"
	ORDER_SIDE_SELL     = "Sell"
)

// Order type (order_type)
const (
	ORDER_TYPE_MARKET = "market"
	ORDER_TYPE_LIMIT  = "Limit"
)

// PlacePerpOrderParams
const (
	PLACE_PERP_BUY                 = "Buy"
	PLACE_PERP_SELL                = "Sell"
	PLACE_PERP_LIMIT               = "Limit"
	PLACE_PERP_MARKET              = "Market"
	PLACE_PERP_GTC                 = "GoodTillCancel"
	PLACE_PERP_IMMEDIATE_OR_CANCEL = "ImmediateOrCancel"
	PLACE_PERP_FILL_OR_KILL        = "FillOrKill"
	PLACE_PERP_POST_ONLY           = "PostOnly"
)
