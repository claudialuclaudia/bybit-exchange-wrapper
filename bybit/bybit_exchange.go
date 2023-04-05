package bybit_exchange

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/fatih/structs"
	"github.com/pingcap/log"
)

type BybitExchange struct {
	Client    *http.Client
	SecretKey string
	ApiKey    string
}

// runtime bybit exchange client instance
var bybitExchange_ *BybitExchange

// -------- Runtime State
var isInitRest_ bool

var isTestnet = true

// Generates new bybit exchange client.
// Requires: Secret and Api keys
func NewBybitExchangeClient(secretKey, apiKey string) {
	bybitExchange_ = &BybitExchange{
		Client:    &http.Client{},
		SecretKey: secretKey,
		ApiKey:    apiKey,
	}
	isInitRest_ = true
}

// Returns initiated bybit exchange client instance with access to methods
func GetBybitExchangeRestService() (*BybitExchange, error) {
	// check if initialised
	if isInitRest_ {
		return bybitExchange_, nil
	}
	err := errors.New("uninitialised exchange rest client")
	return bybitExchange_, err
}

// ---------------------------- GETTERS ----------------------------

/*
	Gets a spot order

	Requires: 
		orderId string

	Returns: 
		spotOrder SpotOrder
		err error

	Ref: https://bybit-exchange.github.io/docs/spot/v1/#t-getactive
*/
func (bybit *BybitExchange) GetSpotOrder(orderId string) (spotOrder SpotOrder, err error) {
	functionName := "GetSpotOrder"
	params := map[string]interface{}{}
	params["orderId"] = orderId

	// create request
	req := bybit.signRequest(http.MethodGet, SPOT_ORDER, params)

	body, err := bybit.getResponseBody(functionName, req)

	var response = new(ResponseForGetSpotOrder)
	if err = json.Unmarshal(body, &response); err != nil {
		err_msg := fmt.Sprintf("%v failed to parse response body: %v", functionName, err)
		err = errors.New(err_msg)
		log.Error(err.Error())
		return spotOrder, err
	}

	if response.ApiResponse.RetMsg != "OK" {
		err_msg := fmt.Sprintf("%v failed: %v", functionName, response.ApiResponse)
		err = errors.New(err_msg)
		log.Error(err.Error())
		return spotOrder, err
	}

	return response.SpotOrder, err
}

/*
	Gets a list of perp orders.

	Requires: 
		symbol string

	Returns:
		perpOrders []PerpOrder
		err error

	Ref: https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-getactive
*/
func (bybit *BybitExchange) GetPerpOrder(symbol string) (perpOrders []PerpOrder, err error) {
	functionName := "GetPerpOrder"
	params := map[string]interface{}{}
	params["symbol"] = symbol

	// create request
	req := bybit.signRequest(http.MethodGet, GET_PERP_ORDER, params)

	body, err := bybit.getResponseBody(functionName, req)

	var response = new(ResponseForGetPerpOrder)
	if err = json.Unmarshal(body, &response); err != nil {
		err_msg := fmt.Sprintf("%v failed to parse response body: %v", functionName, err)
		err = errors.New(err_msg)
		log.Error(err.Error())
		return perpOrders, err
	}

	if response.ApiResponse.RetMsg != "OK" {
		err_msg := fmt.Sprintf("%v failed: %v", functionName, response.ApiResponse)
		err = errors.New(err_msg)
		log.Error(err.Error())
		return perpOrders, err
	}

	return response.Result.PerpOrders, err
}

/*
	Gets spot wallet balance for a particular symbol.

	Requires: 
		symbol string

	Returns: 
		balance float64
		err error

	Refs: https://bybit-exchange.github.io/docs/spot/v1/#t-balance
*/
func (bybit *BybitExchange) GetSpotWalletBalanceForSymbol(symbol string) (balance float64, err error) {
	balances, err := bybit.GetSpotWalletBalance()

	for _, coin := range balances {
		if coin.Coin == symbol {
			balance, err := strconv.ParseFloat(coin.Free, 64)
			return balance, err
		}
	}

	return balance, err
}

/*
	Gets the entire spot wallet balance.

	Requires:
		-

	Returns:
		balances []SpotBalance
		err error
		
	Ref: https://bybit-exchange.github.io/docs/spot/v1/#t-balance
*/
func (bybit *BybitExchange) GetSpotWalletBalance() (balances []SpotBalance, err error) {
	functionName := "GetSpotWalletBalance"
	// create request
	req := bybit.signRequest(http.MethodGet, GET_SPOT_BALANCE, nil)

	body, err := bybit.getResponseBody(functionName, req)

	var response = new(ResponseForGetSpotWalletBalance)
	if err = json.Unmarshal(body, &response); err != nil {
		err_msg := fmt.Sprintf("%v failed to parse response body: %v", functionName, err)
		err = errors.New(err_msg)
		log.Error(err.Error())
		return balances, err
	}

	if response.ApiResponse.RetMsg != "OK" {
		err_msg := fmt.Sprintf("%v failed: %v", functionName, response.ApiResponse)
		err = errors.New(err_msg)
		log.Error(err.Error())
		return balances, err
	}

	return response.Result.Balances, err
}

/*
	Gets the entire perp wallet balance.

	Requires:
		-

	Returns:
		a map of PerpBalances, each one keyed to its token name (e.g. "BTC", "DOT")
		err error

	Ref: https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-balance
*/
func (bybit *BybitExchange) GetPerpWalletBalance() (balances map[string]PerpBalance, err error) {
	functionName := "GetPerpWalletBalance"

	// create request
	req := bybit.signRequest(http.MethodGet, GET_PERP_BALANCE, nil)

	body, err := bybit.getResponseBody(functionName, req)

	var response = new(ResponseForGetPerpWalletBalance)
	if err = json.Unmarshal(body, &response); err != nil {
		err_msg := fmt.Sprintf("%v failed to parse response body: %v", functionName, err)
		err = errors.New(err_msg)
		log.Error(err.Error())
		return balances, err
	}

	if response.ApiResponse.RetMsg != "OK" {
		err_msg := fmt.Sprintf("%v failed: %v", functionName, response.ApiResponse)
		err = errors.New(err_msg)
		log.Error(err.Error())
		return balances, err
	}

	return response.Result, err
}

/*
	Gets the entire spot and perp wallet balance in USD.
	
	Requires:
		-

	Returns:
		balance float64
		err error

	Refs: https://bybit-exchange.github.io/docs/spot/v1/#t-wallet
*/
func (bybit *BybitExchange) GetTotalAcctUsdValue() (balance float64, err error) {
	spotBalances, err := bybit.GetSpotWalletBalance()

	var price float64
	for _, coin := range spotBalances {
		if amount, err := strconv.ParseFloat(coin.Total, 32); err == nil {
			if coin.Coin == "USDT" {
				price = 1
			} else {
				price, _ = bybit.GetMarketPrice(coin.Coin, "USDT", "spot")
			}
			balance += amount * price
		}
	}

	perpBalances, err := bybit.GetPerpWalletBalance()
	for coinName, coin := range perpBalances {
		if coin.WalletBalance > 0 {
			price, _ = bybit.GetMarketPrice(coinName, "USDT", "perp")
			balance += coin.WalletBalance * price
		}
	}

	return balance, err
}

/*
	Gets perp position of symbol.

	Requires:
		symbol string

	Returns:
		position PerpPosition
		err error

	Refs: https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-myposition
*/
func (bybit *BybitExchange) GetPerpPosition(symbol string) (position PerpPosition, err error) {
	functionName := "GetPerpPosition"
	params := map[string]interface{}{}
	params["symbol"] = symbol
	// create request
	req := bybit.signRequest(http.MethodGet, GET_PERP_POSITION, params)

	body, err := bybit.getResponseBody(functionName, req)

	var response = new(ResponseForGetPerpPositions)
	if err = json.Unmarshal(body, &response); err != nil {
		err_msg := fmt.Sprintf("%v failed to parse response body: %v", functionName, err)
		err = errors.New(err_msg)
		log.Error(err.Error())
		return position, err
	}

	if response.ApiResponse.RetMsg != "OK" {
		err_msg := fmt.Sprintf("%v failed: %v", functionName, response.ApiResponse)
		err = errors.New(err_msg)
		log.Error(err.Error())
		return position, err
	}

	return response.Result, err
}

/*
	Gets deposit address for symbol.

	Requires:
		symbol string - currency symbol, e.g. "BTC"
		network string - network symbol, e.g. "LTC"
	
	Returns:
		address string
		err error

	Refs: https://bybit-exchange.github.io/docs/account_asset/v1/#t-deposit_addr_info
*/
func (bybit *BybitExchange) GetDepositAddress(symbol, network string) (address string, err error) {
	functionName := "GetDepositAddress"
	params := map[string]interface{}{}
	params["coin"] = symbol

	// create request
	req := bybit.signRequest(http.MethodGet, GET_DEPOSIT_ADDRESS, params)

	body, err := bybit.getResponseBody(functionName, req)

	var response = new(ResponseForGetDepositAddress)
	if err = json.Unmarshal(body, &response); err != nil {
		err_msg := fmt.Sprintf("%v failed to parse response body: %v", functionName, err)
		err = errors.New(err_msg)
		log.Error(err.Error())
		return address, err
	}

	if response.RetMsg != "OK" {
		err_msg := fmt.Sprintf("%v failed: %v", functionName, response.RetMsg)
		err = errors.New(err_msg)
		log.Error(err.Error())
		return address, err
	}

	// for _, coin := range balances {
	// 	if coin.Coin == symbol {
	// 		balance, err := strconv.ParseFloat(coin.Free, 64)
	// 		return balance, err
	// 	}
	// }

	return response.Result.Chains[0].AddressDeposit, err
}

/*
	Get market price for perp or spot.

	Requires: 
		baseCurr string
		quoteCurr string
		marketType string

	Returns:
		marketPrice float64
		err error
*/
func (bybit *BybitExchange) GetMarketPrice(baseCurr, quoteCurr, marketType string) (marketPrice float64, err error) {
	symbol := baseCurr + quoteCurr
	switch marketType {
	case "perp":
		return bybit.GetPerpMarketPrice(symbol)
	case "spot":
		return bybit.GetSpotMarketPrice(symbol)
	}
	return marketPrice, err
}

/*
	Gets spot market price using best bid price.

	Requires: 
		symbol (baseCurr + quoteCurr) string

	Returns: 
		marketPrice float64
		err error

	Refs: https://bybit-exchange.github.io/docs/spot/v1/#t-bestbidask
*/
func (bybit *BybitExchange) GetSpotMarketPrice(symbol string) (marketPrice float64, err error) {
	functionName := "GetSpotMarketPrice"
	params := map[string]interface{}{}
	params["symbol"] = symbol

	// create request
	req := bybit.createRequest(http.MethodGet, SPOT_MARKET_PRICE, params)

	body, err := bybit.getResponseBody(functionName, req)

	// parse response
	var response = new(ResponseForGetSpotMarketPrice)

	if err := json.Unmarshal(body, &response); err != nil {
		err_msg := fmt.Sprintf("%v failed to parse response body: %v", functionName, err)
		err = errors.New(err_msg)
		log.Error(err.Error())
		return marketPrice, err
	}

	if response.Result.BidPrice == "" {
		err_msg := fmt.Sprintf("%v failed: %v", functionName, response.ApiResponse)
		err = errors.New(err_msg)
		log.Error(err.Error())
		return marketPrice, err
	}

	price, err := strconv.ParseFloat(response.Result.BidPrice, 64)
	return price, err
}

/*
	Gets Perp Market Price using mark price

	Requires:
		symbol (baseCurr + quoteCurr) string

	Returns:
		marketPrice float64
		err error

	Refs: https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-latestsymbolinfo
*/
func (bybit *BybitExchange) GetPerpMarketPrice(symbol string) (marketPrice float64, err error) {
	// runtime variables
	functionName := "GetPerpMarketPrice"
	params := map[string]interface{}{}
	params["symbol"] = symbol

	// create request
	req := bybit.createRequest(http.MethodGet, TICKER, params)

	body, err := bybit.getResponseBody(functionName, req)

	// parse response
	var response = new(ResponseForGetPerpMarketPrice)

	if err := json.Unmarshal(body, &response); err != nil {
		err_msg := fmt.Sprintf("%v failed to parse response body: %v", functionName, err)
		err = errors.New(err_msg)
		log.Error(err.Error())
		return marketPrice, err
	}

	if response.ApiResponse.RetMsg != "OK" {
		err_msg := fmt.Sprintf("%v failed: %v", functionName, response.ApiResponse)
		err = errors.New(err_msg)
		log.Error(err.Error())
		return marketPrice, err
	}

	price, err := strconv.ParseFloat(response.Result[0].MarkPrice, 64)
	return price, err
}

/*
	Gets Perp Tick Size. Useful for calculating things like the allowed bid price range for order (https://bybit-exchange.github.io/docs/testnet/futuresV2/inverse/#price-price)

	Requires:
		symbol (baseCurr + quoteCurr) string - e.g. "BTCUSD"

	Returns: 
		maxPrice float64
		tickSize float64
		err error

	Refs: https://bybit-exchange.github.io/docs/testnet/futuresV2/inverse/#t-querysymbol
*/
func (bybit *BybitExchange) GetPerpTickSize(symbol string) (maxPrice, tickSize float64, err error) {
	functionName := "GetPerpTickSize"

	// create request
	req := bybit.createRequest(http.MethodGet, GET_PERP_SYMBOLS_INFO, nil)

	body, err := bybit.getResponseBody(functionName, req)

	// parse response
	var response = new(ResponseForGetPerpSymbolsInfo)

	if err := json.Unmarshal(body, &response); err != nil {
		err_msg := fmt.Sprintf("%v failed to parse response body: %v", functionName, err)
		err = errors.New(err_msg)
		log.Error(err.Error())
		return maxPrice, tickSize, err
	}

	if response.ApiResponse.RetMsg != "OK" {
		err_msg := fmt.Sprintf("%v failed: %v", functionName, response.ApiResponse)
		err = errors.New(err_msg)
		log.Error(err.Error())
		return maxPrice, tickSize, err
	}

	for _, perpSymbolInfo := range response.Result {
		if perpSymbolInfo.Name == symbol {
			maxPrice, err = strconv.ParseFloat(perpSymbolInfo.PriceFilter.MaxPrice, 64)
			tickSize, err = strconv.ParseFloat(perpSymbolInfo.PriceFilter.TickSize, 64)
			return maxPrice, tickSize, err
		}
	}

	return maxPrice, tickSize, err
}

/*
	Get Bybit Server Time. All requests are required to be sent with a timestamp such that
	server_time - recv_window (unit in millisecond and default value is 5,000) <= timestamp < server_time + 1000

	Returns: 
		serverTime float64
		err error

	Refs: https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-servertime
*/
func (bybit *BybitExchange) GetServerTime() (serverTime float64, err error) {
	functionName := "GetServerTime"
	url := SERVER_TIME

	req := bybit.createRequest(http.MethodGet, url, nil)
	body, err := bybit.getResponseBody(functionName, req)

	// parse response
	var response = new(ResponseForGetServerTime)
	if err := json.Unmarshal(body, &response); err != nil {
		err_msg := fmt.Sprintf("get server time failed to parse response body: %v", err)
		err = errors.New(err_msg)
		log.Error(err.Error())
		return serverTime, err
	}

	if response.ApiResponse.RetMsg != "OK" {
		err_msg := fmt.Sprintf("get server time failed: %v", response.ApiResponse)
		err = errors.New(err_msg)
		log.Error(err.Error())
		return serverTime, err
	}

	time, err := strconv.ParseFloat(response.TimeNow, 64)
	return time, err
}

// ---------------------------- POST CALLS ----------------------------

/*
	Creates a spot order.

	Requires:
		params PlaceSpotOrderParams

	Returns: 
		orderId string
		err error

	Ref : https://bybit-exchange.github.io/docs/spot/v1/#t-placeactive
*/
func (bybit *BybitExchange) PlaceSpotOrder(params PlaceSpotOrderParams) (orderId string, err error) {
	functionName := "PlaceSpotOrder"
	// prepare request body by turning exchange.PlaceSpotOrderParams into map[string]interface{}:
	params_map := structs.Map(&params)

	// create request
	req := bybit.signRequest(http.MethodPost, SPOT_ORDER, params_map)

	body, err := bybit.getResponseBody(functionName, req)

	var response = new(ResponseForPlaceSpotOrder)
	if err = json.Unmarshal(body, &response); err != nil {
		err_msg := fmt.Sprintf("%v failed to parse response body: %v", functionName, err)
		err = errors.New(err_msg)
		log.Error(err.Error())
		return orderId, err
	}

	if response.ApiResponse.RetMsg != "OK" {
		err_msg := fmt.Sprintf("%v failed: %v", functionName, response.ApiResponse)
		err = errors.New(err_msg)
		log.Error(err.Error())
		return orderId, err
	}

	return response.Result.OrderId, err
}

/*
	Creates a perp order

	Requires:
		params PlacePerpOrderParams
	
	Returns:
		orderId string
		err error

	Ref: https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-activeorders
*/
func (bybit *BybitExchange) PlacePerpOrder(params PlacePerpOrderParams) (orderId string, err error) {
	functionName := "PlacePerpOrder"
	// prepare request body by turning exchange.PlacePerpOrderParams into map[string]interface{}:
	params_map := structs.Map(&params)

	// create request
	req := bybit.signRequest(http.MethodPost, PLACE_PERP_ORDER, params_map)

	body, err := bybit.getResponseBody(functionName, req)

	var response = new(ResponseForPlacePerpOrder)
	if err = json.Unmarshal(body, &response); err != nil {
		err_msg := fmt.Sprintf("%v failed to parse response body: %v", functionName, err)
		err = errors.New(err_msg)
		log.Error(err.Error())
		return orderId, err
	}

	if response.ApiResponse.RetMsg != "OK" {
		err_msg := fmt.Sprintf("%v failed: %v", functionName, response.ApiResponse)
		err = errors.New(err_msg)
		log.Error(err.Error())
		return orderId, err
	}

	return response.Result.OrderId, err
}

/*
	Withdrawls from exchange. Withdrawals are always made from the SPOT wallet.

	Requires:
		coin string - e.g. "BTC"
		amount float64 - amont to withdrawl
		destinationAddress string - wallet adress
		network string - network name. e.g. "LTC"

	Returns:
		withdrawOrderId string
		err error

	Refs: https://bybit-exchange.github.io/docs/account_asset/v1/#t-withdraw_info
*/
func (bybit *BybitExchange) WithdrawFromExchange(coin string, amount float64, destinationAddress, network string) (withdrawOrderId string, err error) {
	functionName := "WithdrawFromExchange"

	params := map[string]interface{}{}
	params["coin"] = coin
	params["amount"] = fmt.Sprintf("%f", amount)
	params["chain"] = network
	params["address"] = destinationAddress

	// create request
	req := bybit.signRequestWithSignAsABodyParam(http.MethodPost, WITHDRAW_FROM_SPOT_WALLET, params)

	body, err := bybit.getResponseBody(functionName, req)

	var response = new(ResponseForWithdrawlFromSpotWallet)
	if err = json.Unmarshal(body, &response); err != nil {
		err_msg := fmt.Sprintf("%v failed to parse response body: %v", functionName, err)
		err = errors.New(err_msg)
		log.Error(err.Error())
		return withdrawOrderId, err
	}

	if response.RetMsg != "OK" {
		err_msg := fmt.Sprintf("%v failed: %v", functionName, response.RetMsg)
		err = errors.New(err_msg)
		log.Error(err.Error())
		return withdrawOrderId, err
	}

	return response.Result.WithdrawlId, err
}

// ---------------------------- DELETE CALLS ----------------------------

/*
	Cancels a spot order

	Requires:
		orderId string

	Returns: 
		status bool
		err error

	Ref: https://bybit-exchange.github.io/docs/spot/v1/#t-cancelactive
*/
func (bybit *BybitExchange) CancelSpotOrder(orderId string) (status bool, err error) {
	functionName := "CancelSpotOrder"
	params := map[string]interface{}{}
	params["orderId"] = orderId

	// create request
	req := bybit.signRequest(http.MethodDelete, SPOT_ORDER, params)

	body, err := bybit.getResponseBody(functionName, req)

	var response = new(ResponseForCancelSpotOrder)
	if err = json.Unmarshal(body, &response); err != nil {
		err_msg := fmt.Sprintf("%v failed to parse response body: %v", functionName, err)
		err = errors.New(err_msg)
		log.Error(err.Error())
		return false, err
	}

	if response.ApiResponse.RetMsg != "OK" {
		err_msg := fmt.Sprintf("%v failed: %v", functionName, response.ApiResponse)
		err = errors.New(err_msg)
		log.Error(err.Error())
		return false, err
	}

	return true, err
}

/*
	Cancels a perp order

	Requires:
		symbol string
		orderId string

	Returns: 
		status bool
		err error

	Ref: https://bybit-exchange.github.io/docs/futuresV2/inverse/#t-cancelactive
*/
func (bybit *BybitExchange) CancelPerpOrder(symbol, orderId string) (status bool, err error) {
	functionName := "CancelPerpOrder"
	params := map[string]interface{}{}
	params["symbol"] = symbol
	params["order_id"] = orderId

	// create request
	req := bybit.signRequest(http.MethodPost, CANCEL_PERP_ORDER, params)

	body, err := bybit.getResponseBody(functionName, req)

	var response = new(ResponseForCancelPerpOrder)
	if err = json.Unmarshal(body, &response); err != nil {
		err_msg := fmt.Sprintf("%v failed to parse response body: %v", functionName, err)
		err = errors.New(err_msg)
		log.Error(err.Error())
		return false, err
	}

	if response.ApiResponse.RetMsg != "OK" {
		err_msg := fmt.Sprintf("%v failed: %v", functionName, response.ApiResponse)
		err = errors.New(err_msg)
		log.Error(err.Error())
		return false, err
	}

	return true, err
}

// ---------------------------- HELPERS ----------------------------

func (bybit *BybitExchange) getResponseBody(functionName string, req *http.Request) (body []byte, err error) {
	// create request
	if err != nil {
		err_msg := fmt.Sprintf("%v failed to create request object: %v", functionName, err)
		err = errors.New(err_msg)
		log.Error(err.Error())
		return body, err
	}

	// make request
	res, err := bybit.Client.Do(req)
	if err != nil {
		err_msg := fmt.Sprintf("%v failed to make request: %v", functionName, err)
		err = errors.New(err_msg)
		log.Error(err.Error())
		return body, err
	}
	defer res.Body.Close()

	// read response
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		err_msg := fmt.Sprintf("%v failed to read response body: %v", functionName, err)
		err = errors.New(err_msg)
		log.Error(err.Error())
		return body, err
	}

	return body, err
}

func (bybit *BybitExchange) signRequest(method string, path string, params map[string]interface{}) *http.Request {
	if params == nil {
		params = map[string]interface{}{}
	}
	server_time, _ := bybit.GetServerTime()
	// timestamp must be (in milliseconds): serverTime - recvWindow (default 5,000) <= timestamp < serverTime + 1000
	timestamp := int(server_time*1000 - 100)

	params["api_key"] = bybit.ApiKey
	params["timestamp"] = fmt.Sprintf("%d", timestamp)

	param := bybit.makeParamString(params)

	signature := bybit.sign(param)
	param += "&sign=" + signature

	var fullURL string
	if isTestnet {
		fullURL = TESTNET_URL + path + "?" + param
	} else {
		fullURL = MAINNET_URL + path + "?" + param
	}

	req, _ := http.NewRequest(method, fullURL, bytes.NewReader(make([]byte, 0)))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return req
}

func (bybit *BybitExchange) signRequestWithSignAsABodyParam(method string, path string, params map[string]interface{}) *http.Request {
	if params == nil {
		params = map[string]interface{}{}
	}
	server_time, _ := bybit.GetServerTime()
	// timestamp must be (in milliseconds): serverTime - recvWindow (default 5,000) <= timestamp < serverTime + 1000
	timestamp := int(server_time*1000 - 100)

	params["api_key"] = bybit.ApiKey
	params["timestamp"] = fmt.Sprintf("%d", timestamp)

	param := bybit.makeParamString(params)

	signature := bybit.sign(param)
	params["sign"] = signature

	postBody, _ := json.Marshal(params)

	var fullURL string
	if isTestnet {
		fullURL = TESTNET_URL + path
	} else {
		fullURL = MAINNET_URL + path
	}

	req, _ := http.NewRequest(method, fullURL, bytes.NewBuffer(postBody))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return req
}

func (bybit *BybitExchange) makeParamString(params map[string]interface{}) (param string) {
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	// keys must be sorted in alphabetical order when signed
	sort.Strings(keys)

	var p []string
	for _, k := range keys {
		p = append(p, fmt.Sprintf("%v=%v", k, params[k]))
	}

	param = strings.Join(p, "&")
	return param
}

func (bybit *BybitExchange) createRequest(method, path string, params map[string]interface{}) *http.Request {
	if params == nil {
		params = map[string]interface{}{}
	}

	param := bybit.makeParamString(params)

	var fullURL string
	if isTestnet {
		fullURL = TESTNET_URL + path + "?" + param
	} else {
		fullURL = MAINNET_URL + path + "?" + param
	}

	req, _ := http.NewRequest(method, fullURL, bytes.NewReader(make([]byte, 0)))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return req
}

func (bybit *BybitExchange) sign(signaturePayload string) string {
	mac := hmac.New(sha256.New, []byte(bybit.SecretKey))
	mac.Write([]byte(signaturePayload))
	return hex.EncodeToString(mac.Sum(nil))
}
