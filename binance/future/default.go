package future

import (
	"github.com/workfoxes/tripwire/binance"
)

type OrderRequestOption struct {
	Symbol           string                   `json:"symbol"`
	Side             binance.SideType         `json:"side"`
	OrderType        binance.OrderType        `json:"type"`
	Quantity         string                   `json:"quantity"`
	NewOrderRespType binance.NewOrderRespType `json:"newOrderRespType"`

	TimeInForce      *binance.TimeInForceType `json:"timeInForce"`
	ReduceOnly       *bool                    `json:"reduceOnly"`
	Price            *string                  `json:"price"`
	NewClientOrderID *string                  `json:"newClientOrderId"`
	StopPrice        *string                  `json:"stopPrice"`
	PriceProtect     *bool                    `json:"priceProtect"`
	ActivationPrice  *string                  `json:"activationPrice"`
	CallbackRate     *string                  `json:"callbackRate"`
	ClosePosition    *bool                    `json:"closePosition"`
	//positionSide     *PositionSideType `json:"positionSide"`
	//workingType      *WorkingType `json:"workingType"`
}

// Order define order info
type Order struct {
	Symbol           string                  `json:"symbol"`
	OrderID          int64                   `json:"orderId"`
	ClientOrderID    string                  `json:"clientOrderId"`
	Price            string                  `json:"price"`
	ReduceOnly       bool                    `json:"reduceOnly"`
	OrigQuantity     string                  `json:"origQty"`
	ExecutedQuantity string                  `json:"executedQty"`
	CumQuantity      string                  `json:"cumQty"`
	CumQuote         string                  `json:"cumQuote"`
	Status           binance.OrderStatusType `json:"status"`
	TimeInForce      binance.TimeInForceType `json:"timeInForce"`
	Type             binance.OrderType       `json:"type"`
	Side             binance.SideType        `json:"side"`
	StopPrice        string                  `json:"stopPrice"`
	Time             int64                   `json:"time"`
	UpdateTime       int64                   `json:"updateTime"`
	ActivatePrice    string                  `json:"activatePrice"`
	PriceRate        string                  `json:"priceRate"`
	AvgPrice         string                  `json:"avgPrice"`
	OrigType         string                  `json:"origType"`
	PriceProtect     bool                    `json:"priceProtect"`
	ClosePosition    bool                    `json:"closePosition"`
	//WorkingType      binance.WorkingType      `json:"workingType"`
	//PositionSide     binance.PositionSideType `json:"positionSide"`
}

type CreateOrderResponse struct {
	Order
	RateLimitOrder10s string `json:"rateLimitOrder10s,omitempty"`
	RateLimitOrder1m  string `json:"rateLimitOrder1m,omitempty"`
}
