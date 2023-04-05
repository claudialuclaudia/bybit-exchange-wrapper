package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/0xSaiki/pawo-exchange-wrappers/ftx_exchange"
	"github.com/0xSaiki/pawo-exchange-wrappers/util"
)

/*
	Initiates ftx websocket runtime and listens for subscribed channels
*/
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Load config
	config, config_err := util.LoadConfig("config", ".")
	if config_err != nil {
		log.Fatal("cannot load config:", config_err)
	}

	// Init Ftx websocket wrapper
	ftx_exchange.NewFtxExchangeWsClient(config.FtxSecretKey, config.FtxApiKey, "")

	// Setup cleanup and bind Ctrl+C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	defer func() {
		signal.Stop(c)
		cancel()
	}()

	// Listen for Ctrl+C
	go func() {
		select {
		case <-c:
			signal.Stop(c)
			cancel()
		case <-ctx.Done():
		}
	}()

	// Get websocket instance
	exchange_ws, _ := ftx_exchange.GetFtxExchangeWsService()

	// create response channel
	ft_ch := exchange_ws.CreateChannel()

	// initiate and subscribe to channels
	go exchange_ws.ConnectToPublic(ctx, ft_ch, []string{"orderbook", "trades"}, []string{"ETH/USDT", "BTC/USDT", "SOL/USDT"})

	for {
		select {
		case v := <-ft_ch:
			switch v.Type {
			case ftx_exchange.TICKER:
				fmt.Printf("%s	%+v\n", v.Symbol, v.Ticker)

			case ftx_exchange.TRADES:
				fmt.Printf("%s	%+v\n", v.Symbol, v.Trades)
				for i := range v.Trades {
					if v.Trades[i].Liquidation {
						fmt.Printf("-----------------------------%+v\n", v.Trades[i])
					}
				}

			case ftx_exchange.ORDERBOOK:
				fmt.Printf("%s	%+v\n", v.Symbol, v.Orderbook)

			case ftx_exchange.ORDERS_WS:
				fmt.Printf("%d	%+v\n", v.Type, v.Orders)

			case ftx_exchange.FILLS:
				fmt.Printf("%d	%+v\n", v.Type, v.Fills)

			case ftx_exchange.UNDEFINED:
				fmt.Printf("UNDEFINED %s	%s\n", v.Symbol, v.Error.Error())
			}
		}
	}
}
