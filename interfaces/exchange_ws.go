/*
	Exchange Websocket Api interface
*/
package exchange

import "context"

type ExchangeWsClient interface {

	/*
		Creates a go channel to pipe websocket response through

		Returns: websocket response channel
	*/
	CreateChannel() chan WsResponse

	/*
		Connects and subscribes to private websocket channels

		Requires: 	ctx context.Context
					ch go channel to pipe response
					channels []string eg "orders" "fills"

		Returns: error.
	*/
	ConnectToPrivate(ctx context.Context, ch chan WsResponse, channels []string) (err error)

	/*
		Connects and subcribes to public websocket channels.

		Requires: 	ctx context.Context
					ch go channel to pipe response
					channels []string eg "orderbook" "trades"
					symbols []string eg "ETH-PERP" "ETH/USDT"

		Returns: error.
	*/
	ConnectToPublic(ctx context.Context, ch chan WsResponse, channels, symbols []string) (err error)
}
