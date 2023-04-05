package bybit_exchange

type ApiResponse struct {
	RetCode int    `json:"ret_code"`
	RetMsg  string `json:"ret_msg"`
	ExtCode string `json:"ext_code"`
	ExtInfo string `json:"ext_info"`
	TimeNow string `json:"time_now"`
}

type ResponseForGetSpotMarketPrice struct {
	ApiResponse
	Result SpotMarketPrice `json:"result"`
}

type SpotMarketPrice struct {
	BidPrice string `json:"bidPrice"`
	BidQty   string `json:"bidQty"`
	AskPrice string `json:"askPrice"`
}

type ResponseForGetPerpMarketPrice struct {
	ApiResponse
	Result []PerpMarketPrice `json:"result"`
}

type ResponseForGetPerpSymbolsInfo struct {
	ApiResponse
	Result []PerpSymbolInfo `json:"result"`
}

type PerpSymbolInfo struct {
	Name        string      `json:"name"`
	PriceFilter PriceFilter `json:"price_filter"`
}

type PriceFilter struct {
	MaxPrice string `json:"max_price"`
	TickSize string `json:"tick_size"`
}

type ResponseForGetServerTime struct {
	ApiResponse
	TimeNow string `json:"time_now"`
}

type ResponseForPlacePerpOrder struct {
	ApiResponse
	Result PlacePerpOrderResult `json:"result"`
}

type PlacePerpOrderResult struct {
	UserId        int     `json:"user_id"`
	OrderId       string  `json:"order_id"`
	Symbol        string  `json:"symbol"`
	Side          string  `json:"side"`
	OrderType     string  `json:"order_type"`
	Price         float64 `json:"price"`
	Qty           float64 `json:"qty"`
	TimeInForce   string  `json:"time_in_force"`
	OrderStatus   string  `json:"order_status"`
	LastExecTime  int     `json:"last_exec_time"`
	LastExecPrice float64 `json:"last_exec_price"`
	LeavesQty     float64 `json:"leaves_qty"`
	CumExecQty    float64 `json:"cum_exec_qty"`
	CumExecValue  float64 `json:"cum_exec_value"`
	CumExecFee    float64 `json:"cum_exec_fee"`
	RejectReason  string  `json:"reject_reason"`
	OrderLinkId   string  `json:"order_link_id"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
	TakeProfit    string  `json:"take_profit"`
	StopLoss      string  `json:"stop_loss"`
	TpTriggerBy   string  `json:"tp_trigger_by"`
	SlTriggerBy   string  `json:"sl_trigger_by"`
}

type ResponseForPlaceSpotOrder struct {
	ApiResponse
	Result PlaceSpotOrderResult `json:"result"`
}

type PlaceSpotOrderResult struct {
	OrderId      string `json:"orderId"`
	OrderLinkId  string `json:"orderLinkId"`
	Symbol       string `json:"symbol"`
	SymbolName   string `json:"symbolName"`
	TransactTime string `json:"transactTime"`
	Price        string `json:"price"`
	OrigQty      string `json:"origQty"`
	ExecutedQty  string `json:"executedQty"`
	Type         string `json:"type"`
	Side         string `json:"side"`
	Status       string `json:"status"`
	TimeInForce  string `json:"timeInForce"`
	AccountId    string `json:"accountId"`
}

type ResponseForGetSpotWalletBalance struct {
	ApiResponse
	Result SpotBalances `json:"result"`
}

type SpotBalances struct {
	Balances []SpotBalance `json:"balances"`
}

type SpotBalance struct {
	Coin     string `json:"coin"`
	CoinId   string `json:"coinId"`
	CoinName string `json:"coinName"`
	Total    string `json:"total"`
	Free     string `json:"free"`
	Locked   string `json:"locked"`
}

type ResponseForGetPerpWalletBalance struct {
	ApiResponse
	Result map[string]PerpBalance `json:"result"`
}

type PerpBalance struct {
	Equity           float64 `json:"equity"`
	AvailableBalance float64 `json:"available_balance"`
	UsedMargin       float64 `json:"used_margin"`
	OrderMargin      float64 `json:"order_margin"`
	PositionMargin   float64 `json:"position_margin"`
	OccClosingFee    float64 `json:"occ_closing_fee"`
	OccFundingFee    float64 `json:"occ_funding_fee"`
	WalletBalance    float64 `json:"wallet_balance"`
	RealisedPnl      float64 `json:"realised_pnl"`
	UnrealisedPnl    float64 `json:"unrealised_pnl"`
	CumRealisedPnl   float64 `json:"cum_realised_pnl"`
	GivenCash        float64 `json:"given_cash"`
	ServiceCash      float64 `json:"service_cash"`
}

type PerpMarketPrice struct {
	Symbol     string `json:"symbol"`
	MarkPrice  string `json:"mark_price"`
	IndexPrice string `json:"index_price"`
}

type ResponseForGetSpotOrder struct {
	ApiResponse
	SpotOrder SpotOrder `json:"result"`
}

type SpotOrder struct {
	AccountId           string `json:"accountId"`
	ExchangeId          string `json:"exchangeId"`
	Symbol              string `json:"symbol"`
	SymbolName          string `json:"symbolName"`
	OrderLinkId         string `json:"orderLinkId"`
	OrderId             string `json:"orderId"`
	Price               string `json:"price"`
	OrigQty             string `json:"origQty"`
	ExecutedQty         string `json:"executedQty"`
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
	AvgPrice            string `json:"avgPrice"`
	Status              string `json:"status"`
	TimeInForce         string `json:"timeInForce"`
	Type                string `json:"type"`
	Side                string `json:"side"`
	StopPrice           string `json:"stopPrice"`
	IcebergQty          string `json:"icebergQty"`
	Time                string `json:"time"`
	UpdateTime          string `json:"updateTime"`
	IsWorking           bool   `json:"isWorking"`
	Locked              string `json:"locked"`
}

type ResponseForGetPerpOrder struct {
	ApiResponse
	Result PerpOrders `json:"result"`
}

type PerpOrders struct {
	PerpOrders []PerpOrder `json:"data"`
}

type PerpOrder struct {
	UserId       int    `json:"user_id"`
	PositionIdx  int    `json:"position_idx"`
	OrderStatus  string `json:"order_status"`
	Symbol       string `json:"symbol"`
	Side         string `json:"side"`
	OrderType    string `json:"order_type"`
	Price        string `json:"price"`
	Qty          string `json:"qty"`
	TimeInForce  string `json:"time_in_force"`
	OrderLinkId  string `json:"order_link_id"`
	OrderId      string `json:"order_id"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	LeavesQty    string `json:"leaves_qty"`
	LeavesValue  string `json:"leaves_value"`
	CumExecQty   string `json:"cum_exec_qty"`
	CumExecFee   string `json:"cum_exec_fee"`
	RejectReason string `json:"reject_reason"`
	TakeProfit   string `json:"take_profit"`
	StopLoss     string `json:"stop_loss"`
	TpTriggerBy  string `json:"tp_trigger_by"`
	SlTriggerBy  string `json:"sl_trigger_by"`
}

type ResponseForGetPerpPositions struct {
	ApiResponse
	Result PerpPosition `json:"result"`
}

type PerpPosition struct {
	Id                  int     `json:"id"`
	UserId              int     `json:"user_id"`
	RiskId              int     `json:"risk_id"`
	Symbol              string  `json:"symbol"`
	Side                string  `json:"side"`
	Size                float64 `json:"size"`
	PositionValue       string  `json:"position_value"`
	EntryPrice          string  `json:"entry_price"`
	IsIsolated          bool    `json:"is_isolated"`
	AutoAddMargin       int     `json:"auto_add_margin"`
	Leverage            string  `json:"leverage"`
	EffectiveLeverage   string  `json:"effective_leverage"`
	PositionMargin      string  `json:"position_margin"`
	LiqPrice            string  `json:"liq_price"`
	BustPrice           string  `json:"bust_price"`
	OccClosingFee       string  `json:"occ_closing_fee"`
	OccFundingFee       string  `json:"occ_funding_fee"`
	TakeProfit          string  `json:"take_profit"`
	StopLoss            string  `json:"stop_loss"`
	TrailingStop        string  `json:"trailing_stop"`
	PositionStatus      string  `json:"position_status"`
	DeleverageIndicator int     `json:"deleverage_indicator"`
	OcCalcData          string  `json:"oc_calc_data"`
	OrderMargin         string  `json:"order_margin"`
	WalletBalance       string  `json:"wallet_balance"`
	RealisedPnl         string  `json:"realised_pnl"`
	UnrealisedPnl       float64 `json:"unrealised_pnl"`
	CumRealisedPnl      string  `json:"cum_realised_pnl"`
	CrossSeq            int     `json:"cross_seq"`
	PositionSeq         int     `json:"position_seq"`
	CreatedAt           string  `json:"created_at"`
	UpdatedAt           string  `json:"updated_at"`
	TpSlMode            string  `json:"tp_sl_mode"`
}

type ResponseForGetDepositAddress struct {
	RetCode int                `json:"ret_code"`
	RetMsg  string             `json:"ret_msg"`
	ExtCode string             `json:"ext_code"`
	ExtInfo string             `json:"ext_info"`
	TimeNow int                `json:"time_now"` // can't use APIResponse because TimeNow is an int not a string...
	Result  CoinDepositAddress `json:"result"`
}

type CoinDepositAddress struct {
	Coin   string                      `json:"coin"`
	Chains []CoinDepositAddressOnChain `json:"chains"`
}

type CoinDepositAddressOnChain struct {
	ChainType      string `json:"chain_type"`
	AddressDeposit string `json:"address_deposit"`
	TagDeposit     string `json:"tag_deposit"`
	Chain          string `json:"chain"`
}

type PlaceSpotOrderParams struct {
	Symbol      string  `structs:"symbol"`
	Qty         float64 `structs:"qty"`
	Side        string  `structs:"side"`
	Type        string  `structs:"type"`
	TimeInForce string  `structs:"timeInForce"`
	Price       float64 `structs:"price"`
	OrderLinkId string  `structs:"orderLinkId"`
}

type PlacePerpOrderParams struct {
	Side           string  `structs:"side"`       //required
	Symbol         string  `structs:"symbol"`     //required
	OrderType      string  `structs:"order_type"` //required
	Qty            float64 `structs:"qty"`        //required (in USD)
	Price          float64 `structs:"price"`
	TimeInForce    string  `structs:"time_in_force"` //required
	ReduceOnly     bool    `structs:"reduce_only"`
	CloseOnTrigger bool    `structs:"close_on_trigger"`
	OrderLinkId    string  `structs:"order_link_id"`
	TakeProfit     float64 `structs:"take_profit"`
	StopLoss       float64 `structs:"stop_loss"`
	TpTriggerBy    string  `structs:"tp_trigger_by"`
	SlTriggerBy    string  `structs:"sl_trigger_by"`
}

type ResponseForCancelSpotOrder struct {
	ApiResponse
	Result CancelSpotOrderResult `json:"result"`
}

type CancelSpotOrderResult struct {
	OrderId      string `json:"orderId"`
	OrderLinkId  string `json:"orderLinkId"`
	Symbol       string `json:"symbol"`
	AccountId    string `json:"accountId"`
	TransactTime string `json:"transactTime"`
	Price        string `json:"price"`
	OrigQty      string `json:"origQty"`
	ExecutedQty  string `json:"executedQty"`
	TimeInForce  string `json:"timeInForce"`
	Type         string `json:"type"`
	Status       string `json:"status"`
	Side         string `json:"side"`
}

type ResponseForCancelPerpOrder struct {
	ApiResponse
	Result           CancelPerpOrderResult `json:"result"`
	RateLimitStatus  int                   `json:"rate_limit_status"`
	RateLimitResetMs float64               `json:"rate_limit_reset_ms"`
	RateLimit        int                   `json:"rate_limit"`
}

type CancelPerpOrderResult struct {
	UserId        int     `json:"user_id"`
	OrderId       string  `json:"order_id"`
	Symbol        string  `json:"symbol"`
	Side          string  `json:"side"`
	OrderType     string  `json:"order_type"`
	Price         float64 `json:"price"`
	Qty           float64 `json:"qty"`
	TimeInForce   string  `json:"time_in_force"`
	OrderStatus   string  `json:"order_status"`
	LastExecTime  float64 `json:"last_exec_time"`
	LastExecPrice float64 `json:"last_exec_price"`
	LeavesQty     float64 `json:"leaves_qty"`
	CumExcQty     float64 `json:"cum_exec_qty"`
	CumExecValue  float64 `json:"cum_exec_value"`
	CumExecFee    float64 `json:"cum_exec_fee"`
	RejectReason  string  `json:"reject_reason"`
	OrderLinkId   string  `json:"order_link_id"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
	TakeProfit    string  `json:"take_profit"`
	StopLoss      string  `json:"stop_loss"`
	TpTriggerBy   string  `json:"tp_trigger_by"`
	SlTriggerBy   string  `json:"sl_trigger_by"`
}

type ResponseForWithdrawlFromSpotWallet struct {
	RetCode          int                           `json:"ret_code"`
	RetMsg           string                        `json:"ret_msg"`
	ExtCode          string                        `json:"ext_code"`
	ExtInfo          string                        `json:"ext_info"`
	TimeNow          float64                       `json:"time_now"` //yep...for some reason this endpoint returns a number as TimeNow instead of a string like all the other endpoints...
	Result           WithdrawlFromSpotWalletResult `json:"result"`
	RateLimitStatus  int                           `json:"rate_limit_status"`
	RateLimitResetMs float64                       `json:"rate_limit_reset_ms"`
	RateLimit        int                           `json:"rate_limit"`
}

type WithdrawlFromSpotWalletResult struct {
	WithdrawlId string `json:"id"`
}
