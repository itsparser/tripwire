package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// WebRequest is a wrapper for http.Request that can Handle the web request
func WebRequest(url string, method string, headers map[string]string, body interface{}) (*http.Response, error) {
	_body, err := BodyConstructor(body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(method, url, _body)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	if headers != nil {
		for key, value := range headers {
			request.Header.Set(key, value)
		}
	}
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Print("Error While triggering the request with payload - ", request, ", error -", err)
		return nil, err
	}
	return resp, nil
}

// BodyConstructor is a wrapper for json.Marshal that can handle the body
func BodyConstructor(body interface{}) (*bytes.Buffer, error) {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(bodyBytes), nil
}

func GetResponse(response *http.Response) ([]byte, error) {
	//defer func(Body io.ReadCloser) {
	//	err := Body.Close()
	//	if err != nil {
	//
	//	}
	//}(response.Body)
	return ioutil.ReadAll(response.Body)
}
