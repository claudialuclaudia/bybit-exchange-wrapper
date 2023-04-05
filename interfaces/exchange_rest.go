/*
	Exchange Rest Api interface
*/
package exchange

type ExchangeRestClient interface {

	/*
		Gets spot balance for currency.

		Requires: symbol string - currency symbol

		Returns: float, error.
	*/
	GetSpotBalance(symbol string) (balance float64, err error)

	/*
		Gets account tvl.

		Returns: float, error.
	*/
	GetTotalAcctUsdValue() (float64, error)

	/*
		Gets order.

		Requires:
				params type GetOrderParams struct {
					OrderId       int 	- order ID
					BaseCurrency  string
					QuoteCurrency string
					MarketType    string - SPOT or PERP order
				}

		Returns: order, error.
	*/
	GetOrder(orderRequest GetOrderParams) (Order, error)

	/*
		Gets perp position for currency.

		Requires:	symbol string - currency symbol

		Returns: perpPosition, error.
	*/
	GetPerpPosition(symbol string) (Position, error)

	/*
		Gets perp collateral balance for currency.

		Requires:	symbol string - currency symbol

		Returns: float, error.
	*/
	GetPerpMktCollateralUSD(symbol string) (float64, error)

	/*
		Gets deposit address for currency.

		Requires:	symbol string - currency symbol
					network string - currency network

		Returns: string, error.
	*/
	GetDepositAddress(symbol, network string) (address string, err error)

	/*
		Gets market price for currency.

		Requires:
				baseCurr string - base currency symbol
				quoteCurr string - quote currency symbol
				marketType string - PERP or SPOT

		Returns: float, error.
	*/
	GetMarketPrice(baseCurr, quoteCurr, marketType string) (float64, error)

	/*
		Creates Spot order.

		Requires: placeOrderParams PlaceSpotOrderParams struct -> {
				BaseCurrency string - base currency symbol
				QuoteCurrency string - quote currency symbol
				Side string - SELL or BUY
				Price string - decimal price of base currency - required for all LIMIT orders
				Type string - Market or Limit order
				Size float64 - decimal quantity of base currency to buy or sell
				ReduceOnly bool - reduce only order
				Ioc        bool -  immediate or cancel
				PostOnly   bool - postonly order
				ClientId   *string - client order id
			}

		Returns: orderId, error.
	*/
	PlaceSpotOrder(params PlaceOrderParams) (orderId string, err error)

	/*
		Creates Perp order.

		Requires: placeOrderParams PlaceSpotOrderParams struct -> {
				BaseCurrency string - base currency symbol
				QuoteCurrency string - quote currency symbol
				Side string - SELL or BUY
				Price string - decimal price of base currency - required for all LIMIT orders
				Type string - Market or Limit order
				Size float64 - decimal quantity of base currency to buy or sell
				ReduceOnly bool - reduce only order
				Ioc        bool -  immediate or cancel
				PostOnly   bool - postonly order
				ClientId   *string - client order id
			}

		Returns: orderId, error.
	*/
	PlacePerpOrder(params PlaceOrderParams) (orderId string, err error)

	/*
		Cancels order.

		Requires: type CancelOrderParams struct {
			OrderId       int64
			BaseCurrency  string
			QuoteCurrency string
			MarketType    string - SPOT or PERP
		}

		Returns: status, error.
	*/
	CancelOrder(params CancelOrderParams) (status bool, err error)

	/*
		Withdraws token from exhcnage.

		Requires: symbol string - currency symbol
				: amount string - decimal amount to withdraw
				: destination string - destination address
				: network string - destination address network

		Returns: string, error.
	*/
	WithdrawFromExchange(symbol string, amount float64, destinationAddress, network string) (withdrawOrderId string, err error)
}
