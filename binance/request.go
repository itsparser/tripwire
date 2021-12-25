package binance

import (
	"github.com/workfoxes/tripwire/internal/utils"
	"log"
	"net/http"
)

// BinanceRequest is a wrapper for http.Request will add the validation for Binance request
func (client *Client) BinanceRequest(req utils.Request, isSigned bool) (*http.Response, error) {
	var err error
	req.SetDefaultValue()
	if client.APIKey != "" {
		req.Headers.Add("X-MBX-APIKEY", client.APIKey)
	}
	if isSigned {
		query := req.Value.Encode()
		_signature := client.Signer.Sign([]byte(query))
		client.Info("Query String - ", query)
		req.Value.Add("signature", _signature)
		client.Debug("signature", _signature)
		req.RawQuery = req.Value.Encode()
	}
	//req.Headers.Set("X-MBX-APIKEY", c.APIKey)
	////req.Headers.Set("Content-Type", "application/json")
	////req.Headers.Set("Accept", "application/json")
	//if isSigned {
	//	bodyBytes, err := json.Marshal(req.Body)
	//	if err != nil {
	//		return nil, err
	//	}
	//	queryString := req.Value.Encode()
	//	bodyString := string(bodyBytes)
	//	req.Data = bytes.NewBufferString(string(bodyBytes))
	//	raw := fmt.Sprintf("%s%s", queryString, bodyString)
	//	mac := hmac.New(sha256.New, []byte(c.SecretKey))
	//	_, err = mac.Write([]byte(raw))
	//	if err != nil {
	//		return nil, err
	//	}
	//	req.Value.Set(Timestampkey, strconv.FormatInt(butils.CurrentTimestamp()-c.TimeOffset, 10))
	//	req.Value.Set(SignatureKey, fmt.Sprintf("%x", mac.Sum(nil)))
	//}
	//req.URL.RawQuery = q.Encode()
	response, err := req.Do()
	log.Print(response, err)
	return response, err
}
