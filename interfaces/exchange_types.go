package exchange

import (
	"math"
	"time"

	jsoniter "github.com/json-iterator/go"
)

// -------------------------- FUNCTION PARAMS --------------------------
type GetOrderParams struct {
	OrderId       int64
	BaseCurrency  string
	QuoteCurrency string
	MarketType    string
}

type CancelOrderParams struct {
	OrderId       int64
	BaseCurrency  string
	QuoteCurrency string
	MarketType    string
}

type PlaceOrderParams struct {
	BaseCurrency  string  `json:"baseCurrency"`
	QuoteCurrency string  `json:"quoteCurrency"`
	Market        string  `json:"market"`
	OrderSide     string  `json:"orderSide"`
	PositionSide  string  `json:"positionSide"`
	Price         float64 `json:"price"`
	OrderType     string  `json:"type"`
	Size          float64 `json:"size"`
	ReduceOnly    bool    `json:"reduceOnly"`
	Ioc           bool    `json:"ioc"`
	PostOnly      bool    `json:"postOnly"`
	ClientId      *string `json:"clientId"`
}

// -------------------------- REST API --------------------------
type ApiResponse struct {
	Error   string `json:"error,omitempty"`
	Success bool   `json:"success"`
}

type ApiResponseResult struct {
	Error   string      `json:"error,omitempty"`
	Success bool        `json:"success"`
	Result  interface{} `json:"result,omitempty"`
}

// ACCOUNT
type Account struct {
	Username string `json:"username"`

	Collateral               float64 `json:"collateral"`
	FreeCollateral           float64 `json:"freeCollateral"`
	TotalAccountValue        float64 `json:"totalAccountValue"`
	TotalPositionSize        float64 `json:"totalPositionSize"`
	InitialMarginRequirement float64 `json:"initialMarginRequirement"`
	Leverage                 float64 `json:"leverage"`

	MakerFee                     float64 `json:"makerFee"`
	TakerFee                     float64 `json:"takerFee"`
	MaintenanceMarginRequirement float64 `json:"maintenanceMarginRequirement"`

	MarginFraction     float64 `json:"marginFraction"`
	OpenMarginFraction float64 `json:"openMarginFraction"`

	Positions []Position `json:"positions"`

	BackstopProvider bool `json:"backstopProvider"`
	Liquidating      bool `json:"liquidating"`
}

type ResponseForAccountInformation struct {
	ApiResponse
	Result Account `json:"result,omitempty"`
}

// BALANCE
type Balances []Balance

type Balance struct {
	Coin                   string  `json:"coin"`
	Free                   float64 `json:"free"`
	Total                  float64 `json:"total"`
	USDValue               float64 `json:"usdvalue"`
	StopBorrow             float64 `json:"spotBorrow"`
	AvailableWithoutBorrow float64 `json:"availableWithoutBorrow"`
}

type ResponseForBalancesAll struct {
	Result Balances `json:"result,omitempty"`
	ApiResponse
}

// FILLS
type Fill struct {
	Future    string `json:"future"`
	Market    string `json:"market"`
	Type      string `json:"type"`
	Liquidity string `json:"liquidity"`

	// only rest follow 2factor
	BaseCurrency  string `json:"baseCurrency"`
	QuoteCurrency string `json:"quoteCurrency"`
	FeeCurrency   string `json:"feeCurrency"`

	Side string `json:"side"`

	Price   float64 `json:"price"`
	Size    float64 `json:"size"`
	Fee     float64 `json:"fee"`
	FeeRate float64 `json:"feeRate"`

	Time time.Time `json:"time"`

	ID      int `json:"id"`
	OrderID int `json:"orderId"`
	TradeID int `json:"tradeId"`
}

// POSITIONS
type Position struct {
	Future string `json:"future"`
	Side   string `json:"side"`

	InitialMarginRequirement     float64 `json:"initialMarginRequirement"`
	MaintenanceMarginRequirement float64 `json:"maintenanceMarginRequirement"`

	EntryPrice                float64 `json:"entryPrice"`
	EstimatedLiquidationPrice float64 `json:"estimatedLiquidationPrice,omitempty"`

	Size           float64 `json:"size"`
	NetSize        float64 `json:"netSize"`
	OpenSize       float64 `json:"openSize"`
	LongOrderSize  float64 `json:"longOrderSize"`
	ShortOrderSize float64 `json:"shortOrderSize"`

	Cost          float64 `json:"cost"`
	UnrealizedPnl float64 `json:"unrealizedPnl"`
	RealizedPnl   float64 `json:"realizedPnl"`

	CollateralUsed         float64 `json:"collateralUsed,omitempty"`
	RecentAverageOpenPrice float64 `json:"recentAverageOpenPrice,omitempty"`
	RecentPnl              float64 `json:"recentPnl,omitempty"`
	RecentBreakEvenPrice   float64 `json:"recentBreakEvenPrice,omitempty"`
	CumulativeBuySize      float64 `json:"cumulativeBuySize,omitempty"`
	CumulativeSellSize     float64 `json:"cumulativeSellSize,omitempty"`
}
type ResponseForPositionsAll struct {
	Result []Position `json:"result,omitempty"`
	ApiResponse
}

// MARKETS
type Orderbook struct {
	Asks [][]float64 `json:"asks"`
	Bids [][]float64 `json:"bids"`
}
type ResponseForMarketDepth struct {
	Result Orderbook `json:"result,omitempty"`
	ApiResponse
}

// ORDERS
type Order struct {
	ClientID string `json:"clientId"`
	Status   string `json:"status"`
	Type     string `json:"type"`
	Future   string `json:"future"`
	Market   string `json:"market"`
	Side     string `json:"side"`

	Price         float64 `json:"price"`
	AvgFillPrice  float64 `json:"avgFillPrice"`
	Size          float64 `json:"size"`
	RemainingSize float64 `json:"remainingSize"`
	FilledSize    float64 `json:"filledSize"`

	ID          int  `json:"id"`
	Ioc         bool `json:"ioc"`
	PostOnly    bool `json:"postOnly"`
	ReduceOnly  bool `json:"reduceOnly"`
	Liquidation bool `json:"liquidation"`

	CreatedAt time.Time `json:"createdAt"`
}

type ResponseForOrder struct {
	ApiResponse
	Result Order `json:"result,omitempty"`
}

// TRADES
type Trade struct {
	ID          int64     `json:"id"`
	Liquidation bool      `json:"liquidation"`
	Price       float64   `json:"price"`
	Side        string    `json:"side"`
	Size        float64   `json:"size"`
	Time        time.Time `json:"time"`
}

type Ticker struct {
	Bid     float64 `json:"bid"`
	Ask     float64 `json:"ask"`
	BidSize float64 `json:"bidSize"`
	AskSize float64 `json:"askSize"`
	Last    float64 `json:"last"`
	Time    Time_   `json:"time"`
}

// WALLET
type ResponseForDepositAddress struct {
	ApiResponse
	Result DepositAddress `json:"result,omitempty"`
}

type DepositAddress struct {
	Address string `json:"address"`
	Tag     string `json:"tag"`
}
type WithdrawRequestParams struct {
	Coin    string  `json:"coin"`
	Size    float64 `json:"size"`
	Address string  `json:"address"`
	// Optionals
	Tag      string `json:"tag,omitempty"`
	Method   string `json:"method,omitempty"`
	Password string `json:"password,omitempty"`
	Code     string `json:"code,omitempty"`
}

type Withdraw struct {
	Coin    string  `json:"coin"`
	Address string  `json:"address"`
	Tag     string  `json:"tag,omitempty"`
	Fee     float64 `json:"fee,omitempty"`
	Id      int     `json:"id,omitempty"`
	Size    float64 `json:"size,omitempty"`
	Status  string  `json:"status,omitempty"`
	Time    string  `json:"code,omitempty"`
	Txid    string  `json:"txid,omitempty"`
}

type ResponseForWithdraw struct {
	ApiResponse
	Result Withdraw `json:"result,omitempty"`
}

// -------------------------- WEBSOCKET API --------------------------

type WsRequest struct {
	Op      string `json:"op"`
	Channel string `json:"channel"`
	Market  string `json:"market"`
}

// {"op": "login", "args": {"key": "<api_key>", "sign": "<signature>", "time": 1111}}
type WsRequestForPrivate struct {
	Op   string                 `json:"op"`
	Args map[string]interface{} `json:"args"`
}

type WsResponse struct {
	Type   int
	Symbol string

	Ticker    Ticker
	Trades    []Trade
	Orderbook Orderbook

	Orders Order
	Fills  Fill

	Error error
}

type WsOrderbook struct {
	Bids [][]float64 `json:"bids"`
	Asks [][]float64 `json:"asks"`
	// Action return update/partial
	Action   string `json:"action"`
	Time     Time_  `json:"time"`
	Checksum int64  `json:"checksum"`
}

// -------------------------- HELPERS --------------------------

var json_ = jsoniter.ConfigCompatibleWithStandardLibrary

const (
	BUY    = "buy"
	SELL   = "sell"
	MARKET = "market"
	LIMIT  = "limit"
	STOP   = "stop"
)

type Time_ struct {
	Time time.Time
}

func (p *Time_) UnmarshalJSON(data []byte) error {
	var f float64
	if err := json_.Unmarshal(data, &f); err != nil {
		return err
	}

	sec, nsec := math.Modf(f)
	p.Time = time.Unix(int64(sec), int64(nsec))
	return nil
}
