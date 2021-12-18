package future

import (
	"encoding/json"
	"fmt"

	"github.com/workfoxes/tripwire/binance"
	"github.com/workfoxes/tripwire/internal/utils"
)

// CreateOrder Will create Order in Binance and return the order object
func CreateOrder(client *binance.Client, param OrderRequestOption) (*Order, error) {
	fullURL := fmt.Sprintf("%s%s", client.BaseURL, binance.FutureCreateOrder)
	_response, err := client.BinanceRequest(utils.Request{URL: fullURL, Method: "POST", Body: param}, true)
	if err != nil {
		return nil, err
	}
	response := new(Order)
	getResponse, err := utils.GetResponse(_response)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(getResponse, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
