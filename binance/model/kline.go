package binance

type WebSocketRequest struct {
	Method string   `json:"method"`
	Params []string `json:"params"`
	ID     int      `json:"id"`
}

type WebSocketResponse struct {
	EventType string `json:"e"`
	Time      int64  `json:"E"`
	Symbol    string `json:"s"`
	KPI       Kpi    `json:"k"`
}

type Event struct {
	Type   string `json:"e"`
	Time   int64  `json:"E"`
	Symbol string `json:"s"`
}

// Order represents single order information.
type Order struct {
	Price    float64
	Quantity float64
}

// OrderBook represents Bids and Asks.
type OrderBook struct {
	LastUpdateID  int `json:"lastUpdateId"`
	Bids          []*Order
	Asks          []*Order
	BidDepthDelta [][]interface{} `json:"b"`
	AskDepthDelta [][]interface{} `json:"a"`
}

type DepthEvent struct {
	Event
	UpdateID int `json:"u"`
	OrderBook
}

type Kpi struct {
	KStartTime       int64  `json:"t"`
	KSEndTime        int64  `json:"T"`
	Symbol           string `json:"s"`
	Interval         string `json:"i"`
	FirstTradeID     int    `json:"f"`
	LastTradeID      int    `json:"L"`
	OpenPrice        string `json:"o"`
	ClosePrice       string `json:"c"`
	HighestPrice     string `json:"h"`
	LowestPrice      string `json:"l"`
	BassAssetVolume  string `json:"v"`
	NoOfTrade        int    `json:"n"`
	IsKlineClosed    bool   `json:"x"`
	QuoteAssetVolume string `json:"q"`
	Buy              string `json:"V"`
	Quote            string `json:"Q"`
	Ignore           string `json:"B"`
}

// Kline represents single Kline information.
type Kline struct {
	Interval                 string  `json:"i"`
	FirstTradeID             int64   `json:"f"`
	LastTradeID              int64   `json:"L"`
	Final                    bool    `json:"x"`
	OpenTime                 float64 `json:"t"`
	CloseTime                float64 `json:"T"`
	Open                     string  `json:"o"`
	High                     string  `json:"h"`
	Low                      string  `json:"l"`
	Close                    string  `json:"c"`
	Volume                   string  `json:"v"`
	NumberOfTrades           int     `json:"n"`
	QuoteAssetVolume         string  `json:"q"`
	TakerBuyBaseAssetVolume  string  `json:"V"`
	TakerBuyQuoteAssetVolume string  `json:"Q"`
}

type KlineEvent struct {
	Event
	Interval     string
	FirstTradeID int64
	LastTradeID  int64
	Final        bool
	Kline        `json:"k"`
}
