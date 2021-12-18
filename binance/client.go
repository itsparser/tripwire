package binance

import (
	"log"
	"net/http"
	"os"
)

type doFunc func(req *http.Request) (*http.Response, error)

// Client define API client
type Client struct {
	APIKey     string
	SecretKey  string
	BaseURL    string
	UserAgent  string
	HTTPClient *http.Client
	Logger     *log.Logger
	TimeOffset int64
	do         doFunc
}

// NewClient initialize an API client instance with API key and secret key.
// You should always call this function before using this SDK.
// Services will be created by the form client.NewXXXService().
func NewClient(apiKey, secretKey string, isTest bool) *Client {
	return &Client{
		APIKey:     apiKey,
		SecretKey:  secretKey,
		BaseURL:    getAPIEndpoint(isTest),
		UserAgent:  "Binance/golang",
		HTTPClient: http.DefaultClient,
		Logger:     log.New(os.Stderr, "Binance ", log.LstdFlags),
	}
}

// getAPIEndpoint return the base endpoint of the Rest API according the UseTestnet flag
func getAPIEndpoint(isTest bool) string {
	if isTest {
		return BaseAPITestnetURL
	}
	return BaseAPIMainURL
}
