package binance

import (
	"github.com/workfoxes/tripwire/internal/utils"
	"io"
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
	Debug      bool
	Logger     *log.Logger
	TimeOffset int64
	do         doFunc
}

// NewClient initialize an API client instance with API key and secret key.
// You should always call this function before using this SDK.
// Services will be created by the form client.NewXXXService().
func NewClient(apiKey, secretKey string) *Client {
	return &Client{
		APIKey:     apiKey,
		SecretKey:  secretKey,
		BaseURL:    getAPIEndpoint(),
		UserAgent:  "Binance/golang",
		HTTPClient: http.DefaultClient,
		Logger:     log.New(os.Stderr, "Binance ", log.LstdFlags),
	}
}

// getAPIEndpoint return the base endpoint of the Rest API according the UseTestnet flag
func getAPIEndpoint() string {
	if false {
		return BaseAPITestnetURL
	}
	return BaseAPIMainURL
}

func (client Client) BinanceRequest(method, url string, body interface{}) (*http.Response, error) {
	response, err := utils.WebRequest(method, url, nil, body)
	if err != nil {
		return nil, err
	}
	_response, err := utils.GetResponse(response)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	return http.DefaultClient.Do(req)
}
