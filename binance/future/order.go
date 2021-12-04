package future

import (
	"fmt"
	"github.com/workfoxes/tripwire/binance"
)

// CreateOrder Will create Order in Binance and return the order object
func CreateOrder(client *Client, param OrderRequestOption) (*Order, error) {
	fullURL := fmt.Sprintf("%s%s", client.BaseURL, binance.FUTURE_CREATE_ORDER)
	return nil, nil
}
