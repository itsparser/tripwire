package main

import (
	"github.com/workfoxes/tripwire/binance"
	"time"
)

func main() {
	// Create a new Binance client
	c := binance.NewWebClient(false)
	//_ch := make(chan []byte, 5)
	//make(chan *common.Item, 20)
	//url := fmt.Sprintf("wss://stream.binance.com:9443/ws/bnbbtc@kline_1m")
	_ch, _ := c.KlineWebsocket("bnbbtc", "1m")

	go func() {
		select {
		default:
			for cValue := range _ch {
				c.Info(cValue)
				if cValue.Final {
					c.Info("Final")
				}
				//var a b_model.WebSocketResponse
				//err := json.Unmarshal(cValue, &a)
				//if err != nil {
				//	fmt.Println("error:", err)
				//}
				//c.Info(a)
			}
		}

	}()
	// Get the current ticker for BTC/USDT
	//err := c.BinanceWebsocket(url, &_ch)
	//if err != nil {
	//	panic(err)
	//}
	for {
		time.Sleep(time.Second * 1000)
	}
}
