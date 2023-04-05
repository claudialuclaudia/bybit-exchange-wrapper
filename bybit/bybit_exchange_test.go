package bybit_exchange

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type BybitTestSuite struct {
	// Dependencies
	suite.Suite
	Exchange *BybitExchange
}

// this function executes before the test suite begins execution
// and loads all test dependencies
func (suite *BybitTestSuite) SetupSuite() {
	viper.AddConfigPath("../")
	viper.SetConfigName("config") // Register config file name (no extension)
	viper.SetConfigType("env")    // Look for specific type
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	apiKey := viper.Get("BYBIT_API_KEY").(string)
	secretKey := viper.Get("BYBIT_API_SECRET").(string)
	NewBybitExchangeClient(secretKey, apiKey)
	suite.Exchange, err = GetBybitExchangeRestService()
}

func (suite *BybitTestSuite) TestGetMarketPriceSpot() {
	fmt.Println(">>> From TestGetMarketPriceSpot")

	// Test variables
	baseCurr := "BTC"
	quoteCurr := "USDT"
	mktType := "spot"

	// Run test
	marketPrice, err := suite.Exchange.GetMarketPrice(baseCurr, quoteCurr, mktType)

	fmt.Printf("Spot Market Price for: %v/%v is: %v\n", baseCurr, quoteCurr, marketPrice)
	fmt.Printf("---------------------------------\n")

	// Assert test
	suite.NoError(err, "Couldn't get spot market price.")
	suite.NotZero(marketPrice, "Returned spot market price is zero")
}

func (suite *BybitTestSuite) TestGetMarketPricePerp() {
	fmt.Println(">>> From TestGetMarketPricePerp")

	// Test variables
	baseCurr := "BTC"
	quoteCurr := "USDT"
	mktType := "perp"

	// Run test
	marketPrice, err := suite.Exchange.GetMarketPrice(baseCurr, quoteCurr, mktType)

	fmt.Printf("Perp Market Price for: %v/%v is: %v\n", baseCurr, quoteCurr, marketPrice)
	fmt.Printf("---------------------------------\n")

	// Assert test
	suite.NoError(err, "Couldn't get perp market price.")
	suite.NotZero(marketPrice, "Returned perp market price is zero")
}

func (suite *BybitTestSuite) TestPlaceSpotOrder() {
	fmt.Println(">>> From TestPlaceSpotOrder")

	// Setup test
	orderParams := PlaceSpotOrderParams{
		Symbol: "BTCUSDT",
		Side:   ORDER_SIDE_BUY,
		Type:   ORDER_TYPE_MARKET,
		Qty:    1.000,
	}

	// Run test
	orderId, err := suite.Exchange.PlaceSpotOrder(orderParams)

	fmt.Printf("Order Id: %v\n", orderId)
	fmt.Printf("---------------------------------\n")

	// Assert test
	suite.NoError(err, "Couldn't place spot order.")
	suite.NotZero(orderId, "Returned order Id is zero")
}

func (suite *BybitTestSuite) TestGetSpotOrder() {
	fmt.Println(">>> From TestGetSpotOrder")

	// Place an order first
	orderParams := PlaceSpotOrderParams{
		Symbol: "BTCUSDT",
		Side:   ORDER_SIDE_BUY,
		Type:   ORDER_TYPE_MARKET,
		Qty:    1.000,
	}

	// Run test
	orderId, err := suite.Exchange.PlaceSpotOrder(orderParams)

	// Run test
	order, err := suite.Exchange.GetSpotOrder(orderId)

	orderJson, _ := json.MarshalIndent(order, "", "\t")

	fmt.Println(string(orderJson))
	fmt.Printf("---------------------------------\n")

	// Assert test
	suite.NoError(err, "Couldn't get spot order.")
	suite.NotZero(order.OrderId, "Returned spot order Id is 0")
}

func (suite *BybitTestSuite) TestPlacePerpOrder() {
	fmt.Println(">>> From TestPlacePerpOrder")

	// Run test
	orderParams := PlacePerpOrderParams{
		Side:        PLACE_PERP_BUY,
		Symbol:      "BTCUSD",
		OrderType:   PLACE_PERP_MARKET,
		Qty:         1,
		TimeInForce: PLACE_PERP_GTC,
	}
	orderId, err := suite.Exchange.PlacePerpOrder(orderParams)

	fmt.Printf("Order Id: %v\n", orderId)
	fmt.Printf("---------------------------------\n")

	// Assert test
	suite.NoError(err, "Couldn't place perp order.")
	suite.NotZero(orderId, "Returned order Id is zero")
}

func (suite *BybitTestSuite) TestGetPerpOrder() {
	fmt.Println(">>> From TestGetPerpOrder")

	// Set up test
	symbol := "BTCUSD"

	// Place an order first
	orderParams := PlacePerpOrderParams{
		Side:        PLACE_PERP_BUY,
		Symbol:      symbol,
		OrderType:   PLACE_PERP_MARKET,
		Qty:         1,
		TimeInForce: PLACE_PERP_GTC,
	}
	_, err := suite.Exchange.PlacePerpOrder(orderParams)

	// Run test
	perpOrders, err := suite.Exchange.GetPerpOrder(symbol)

	perpOrdersJson, _ := json.MarshalIndent(perpOrders, "", "\t")

	fmt.Println(string(perpOrdersJson))
	fmt.Printf("---------------------------------\n")

	// Assert test
	suite.NoError(err, "Couldn't get perp orders.")
	suite.NotNil(perpOrders, "Returned nil perp orders. Should at least be an empty list, even if the place call at first failed.")
}

func (suite *BybitTestSuite) TestGetSpotBalance() {
	fmt.Println(">>> From TestGetSpotBalance")

	// Run test
	spotBalance, err := suite.Exchange.GetSpotWalletBalanceForSymbol("BTC")

	fmt.Printf("Spot balance: %v\n", spotBalance)
	fmt.Printf("---------------------------------\n")

	// Assert test
	suite.NoError(err, "Couldn't get spot balance.")
	suite.NotZero(spotBalance, "Balance shouldn't be zero unless we have no fund in this current spot account")
}

func (suite *BybitTestSuite) TestGetTotalAcctUsdValue() {
	fmt.Println(">>> From TestGetTotalAcctUsdValue")

	// Run test
	totalAcctUsdValue, err := suite.Exchange.GetTotalAcctUsdValue()

	fmt.Printf("Total Account USD Value (Spot + Perp): %v\n", totalAcctUsdValue)
	fmt.Printf("---------------------------------\n")

	// Assert test
	suite.NoError(err, "Couldn't get totalAcctUsdValue.")
	suite.NotZero(totalAcctUsdValue, "TotalAcctUsdValue shouldn't be zero unless we have no fund in this current account")
}

func (suite *BybitTestSuite) TestGetPerpWalletBalance() {
	fmt.Println(">>> From TestGetPerpWalletBalance")

	// Run test
	balances, _ := suite.Exchange.GetPerpWalletBalance()

	balancesJson, err := json.MarshalIndent(balances, "", "\t")

	fmt.Println(string(balancesJson))
	fmt.Printf("---------------------------------\n")

	// Assert test
	suite.NoError(err, "Couldn't get perp wallet balance.")
	suite.Equal(len(balances), 11) // Perp balances should have a length of 11 (ADA, BIT, BTC, DOT, EOS, ETH, LTC, MANA, SOL, USDT, XRP) even if balance of that currency is 0.
}

func (suite *BybitTestSuite) TestGetPerpPosition() {
	fmt.Println(">>> From TestGetPerpPosition")

	// Set up test
	symbol := "BTCUSD"

	// place a perp order first
	orderParams := PlacePerpOrderParams{
		Side:        PLACE_PERP_BUY,
		Symbol:      symbol,
		OrderType:   PLACE_PERP_MARKET,
		Qty:         1,
		TimeInForce: PLACE_PERP_GTC,
	}
	_, err := suite.Exchange.PlacePerpOrder(orderParams)

	// Run test
	position, err := suite.Exchange.GetPerpPosition(symbol)

	positionJson, _ := json.MarshalIndent(position, "", "\t")

	fmt.Println(string(positionJson))
	fmt.Printf("---------------------------------\n")

	// Assert test
	suite.NoError(err, "Couldn't get perp position.")
	// suite.Equal(position.UserId, 623225)
	suite.NotZero(position.UserId, "UserId is Zero.")
}

func (suite *BybitTestSuite) TestGetDepositAddress() {
	fmt.Println(">>> From TestGetDepositAddress")

	// Run test
	address, err := suite.Exchange.GetDepositAddress("BTC", "LTC")

	fmt.Printf("Deposit Address: %v\n", address)
	fmt.Printf("---------------------------------\n")

	// Assert Test
	suite.NoError(err, "Couldn't get address.")
	suite.NotNil(address, "Deposit address is nil.")
	suite.NotEmpty(address, "Deposit address is empty string")
}

func (suite *BybitTestSuite) TestCancelSpotOrder() {
	fmt.Println(">>> From TestCancelSpotOrder")

	// Place an order first
	orderParams := PlaceSpotOrderParams{
		Symbol: "BTCUSDT",
		Side:   ORDER_SIDE_BUY,
		Type:   ORDER_TYPE_LIMIT,
		Price:  1, // a low price so it won't get fulfilled and we can then make the cancel request
		Qty:    1.000,
	}
	orderId, err := suite.Exchange.PlaceSpotOrder(orderParams)

	// Run test
	status, err := suite.Exchange.CancelSpotOrder(orderId)

	fmt.Printf("Status of order cancellation: %v\n", status)
	fmt.Printf("---------------------------------\n")

	// Assert test
	suite.NoError(err, "Couldn't cancel order.")
	suite.True(status, "Cancel order request failed.")
}

func (suite *BybitTestSuite) TestCancelPerpOrder() {
	fmt.Println(">>> From TestCancelPerpOrder")

	// test params
	symbol := "BTCUSD"
	// might have to get perp max price and tick size because Price has some limitations https://bybit-exchange.github.io/docs/futuresV2/inverse/#price-price
	// perpMaxPrice, perpTickSize, err := suite.Exchange.GetPerpTickSize(symbol)
	price := 5000.0 // making price way below market so it doesn't get fulfilled so we can immediately cancel it

	orderParams := PlacePerpOrderParams{
		Side:        PLACE_PERP_BUY,
		Symbol:      symbol,
		OrderType:   PLACE_PERP_LIMIT,
		Qty:         1,
		TimeInForce: PLACE_PERP_GTC,
		Price:       price,
	}
	orderId, err := suite.Exchange.PlacePerpOrder(orderParams)

	// Run test
	status, err := suite.Exchange.CancelPerpOrder(symbol, orderId)

	fmt.Printf("Status of perp order cancellation: %v\n", status)
	fmt.Printf("---------------------------------\n")

	// Assert test
	suite.NoError(err, "Couldn't cancel placed perp order.")
	suite.True(status, "Cancel perp order request failed.")
}

func (suite *BybitTestSuite) TestWithdrawFromExchange() {
	fmt.Println(">>> From TestWithdrawFromExchange")

	whitelistedAddress := "0xabcde" // this address needs to be specifically whitelisted in your account settings

	// Run test
	WithdrawlId, err := suite.Exchange.WithdrawFromExchange("BTC", 0.001, whitelistedAddress, "BTC")

	WithdrawlIdJson, _ := json.MarshalIndent(WithdrawlId, "", "\t")

	fmt.Println(string(WithdrawlIdJson))
	fmt.Printf("---------------------------------\n")

	// Assert test
	suite.NoError(err, "Couldn't withdraw from exchange.")
	suite.NotZero(WithdrawlId, "Returned withdrawl id is zero.")
}

func TestBybitExchangeTestSuite(t *testing.T) {
	suite.Run(t, new(BybitTestSuite))
}
