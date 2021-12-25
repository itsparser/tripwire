package binance

import (
	"context"
	"log"
	"net/http"
	"os"
)

type doFunc func(req *http.Request) (*http.Response, error)

// Client define API client - creates instance
//
// If logger or ctx are not provided, NopLogger and Background context are used as default.
// You can use context for one-time request cancel (e.g. when shutting down the app).
type Client struct {
	APIKey, SecretKey, UserAgent string
	Signer                       Signer
	BaseURL                      string
	IsTest                       bool
	HTTPClient                   *http.Client
	Logger                       *log.Logger
	TimeOffset                   int64
	do                           doFunc
	Ctx                          context.Context
}

// NewClient initialize an API client instance with API key and secret key.
// You should always call this function before using this SDK.
// Services will be created by the form client.NewXXXService().
func NewClient(apiKey string, signer Signer, isTest bool) *Client {
	return &Client{
		APIKey:     apiKey,
		Signer:     signer,
		BaseURL:    getAPIEndpoint(isTest),
		UserAgent:  "Binance/golang",
		HTTPClient: http.DefaultClient,
		Logger:     log.New(os.Stderr, "Binance ", log.LstdFlags),
	}
}

// NewWebClient initialize an API client instance with API key and secret key.
// You should always call this function before using this SDK.
// Services will be created by the form client.NewXXXService().
func NewWebClient(isTest bool) *Client {
	return &Client{
		IsTest:     isTest,
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
