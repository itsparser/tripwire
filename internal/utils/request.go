package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// Request BinanceRequestParam will be the param that will allow to process the request
type Request struct {
	URL, Method string
	Headers     http.Header
	Body        interface{}
	Value       url.Values
	Cookies     []http.Cookie
	Data        io.Reader
}

func (r *Request) SetDefaultValue() {
	r.Headers = make(http.Header)
	r.Headers.Set("Content-Type", "application/json")
	r.Headers.Set("Accept", "application/json")
	r.Value = make(url.Values)
	r.Cookies = make([]http.Cookie, 0)
}

func (r *Request) ReformatBody(url string) {
	r.URL = url
}

// Do will process the request
func (r *Request) Do() (*http.Response, error) {
	var err error
	if r.Data == nil {
		r.Data, err = BodyConstructor(r.Body)
		if err != nil {
			return nil, err
		}
	}
	request, err := http.NewRequest(r.Method, r.URL, r.Data)
	if err != nil {
		return nil, err
	}
	if r.Headers != nil {
		request.Header = r.Headers
	}
	request.Header.Set("Content-Type", "application/json")
	_query := request.URL.Query()
	if r.Value != nil {
		for k, v := range r.Value {
			_query.Add(k, v[0])
		}
	}
	if r.Cookies != nil && len(r.Cookies) > 0 {
		for _, c := range r.Cookies {
			request.AddCookie(&c)
		}
	}
	request.AddCookie(&http.Cookie{Name: "CST", Value: "1"})
	request.URL.RawQuery = _query.Encode()
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Print("Error While triggering the request with payload - ", request, ", error -", err)
		return nil, err
	}
	return resp, nil
}

// WebRequest is a wrapper for http.Request that can Handle the web request
func WebRequest(r Request) (*http.Response, error) {
	return r.Do()
}

// BodyConstructor is a wrapper for json.Marshal that can handle the body
func BodyConstructor(body interface{}) (*bytes.Buffer, error) {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return bytes.NewBufferString(string(bodyBytes)), nil
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

// HeadersConstructor is a wrapper for http.Header that can handle the headers
func HeadersConstructor(headers map[string]string) http.Header {
	header := http.Header{}
	for key, value := range headers {
		header.Set(key, value)
	}
	return header
}

// QueryParamsConstructor is a wrapper for url.Values that can handle the query params
func QueryParamsConstructor(values map[string]string) url.Values {
	values_ := url.Values{}
	for key, value := range values {
		values_.Set(key, value)
	}
	return values_
}
