package binance

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	butils "github.com/workfoxes/tripwire/binance/utils"
	"github.com/workfoxes/tripwire/internal/utils"
)

// BinanceRequest is a wrapper for http.Request will add the validation for Binance request
func (c *Client) BinanceRequest(req utils.Request, isSigned bool) (*http.Response, error) {
	var err error
	req.SetDefaultValue()
	req.Headers.Set("X-MBX-APIKEY", c.APIKey)
	//req.Headers.Set("Content-Type", "application/json")
	//req.Headers.Set("Accept", "application/json")
	if isSigned {
		bodyBytes, err := json.Marshal(req.Body)
		if err != nil {
			return nil, err
		}
		queryString := req.Value.Encode()
		bodyString := string(bodyBytes)
		req.Data = bytes.NewBufferString(string(bodyBytes))
		raw := fmt.Sprintf("%s%s", queryString, bodyString)
		mac := hmac.New(sha256.New, []byte(c.SecretKey))
		_, err = mac.Write([]byte(raw))
		if err != nil {
			return nil, err
		}
		req.Value.Set(Timestampkey, strconv.FormatInt(butils.CurrentTimestamp()-c.TimeOffset, 10))
		req.Value.Set(SignatureKey, fmt.Sprintf("%x", mac.Sum(nil)))
	}
	response, err := req.Do()
	log.Print(response, err)
	return response, err
}
