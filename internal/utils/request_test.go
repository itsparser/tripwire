package utils

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
)

func TestBodyConstructor(t *testing.T) {
	type args struct {
		body interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    *bytes.Buffer
		wantErr bool
	}{
		{
			name: "TestBodyConstructor",
			args: args{
				body: "test",
			},
			want:    bytes.NewBufferString("\"test\""),
			wantErr: false,
		},
		{
			name: "TestBodyConstructor",
			args: args{
				body: nil,
			},
			want:    bytes.NewBufferString("null"),
			wantErr: false,
		},
		{
			name: "TestBodyConstructor",
			args: args{
				body: []byte("test"),
			},
			want:    bytes.NewBufferString("\"dGVzdA==\""),
			wantErr: false,
		},
		{
			name: "TestBodyConstructor",
			args: args{
				body: []string{"test"},
			},
			want:    bytes.NewBufferString("[\"test\"]"),
			wantErr: false,
		},
		{
			name: "TestBodyConstructor",
			args: args{
				body: map[string]interface{}{
					"test": "test",
				},
			},
			want:    bytes.NewBufferString("{\"test\":\"test\"}"),
			wantErr: false,
		},
		{
			name: "TestBodyConstructor",
			args: args{
				body: map[string]interface{}{
					"test": map[string]interface{}{
						"test": "test",
					},
				},
			},
			want:    bytes.NewBufferString("{\"test\":{\"test\":\"test\"}}"),
			wantErr: false,
		},
		{
			name: "TestBodyConstructor",
			args: args{
				body: map[string]interface{}{
					"test": map[string]interface{}{
						"test": map[string]interface{}{
							"test": "test",
						},
					},
				},
			},
			want:    bytes.NewBufferString("{\"test\":{\"test\":{\"test\":\"test\"}}}"),
			wantErr: false,
		},
		//{
		//	name: "TestBodyConstructor",
		//	args: args{
		//		body: map[string]interface{}{
		//			"test": map[string]interface{}{
		//				"test": map[string]interface{}{
		//					"test": map[string]interface{}{
		//						"test": "test",
		//					},
		//				},
		{
			name: "TestBodyConstructor",
			args: args{
				body: []interface{}{
					"test",
					"test2",
				},
			},
			want:    bytes.NewBufferString("[\"test\",\"test2\"]"),
			wantErr: false,
		},
		{
			name: "TestBodyConstructor",
			args: args{
				body: map[string]interface{}{
					"test": "test",
				},
			},
			want:    bytes.NewBufferString("{\"test\":\"test\"}"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BodyConstructor(tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("BodyConstructor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BodyConstructor() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWebRequest(t *testing.T) {
	tests := []struct {
		name    string
		args    Request
		want    string
		wantErr bool
	}{
		{
			name: "TestWebRequest",
			args: Request{
				URL:    "https://httpbin.org/get",
				Method: "GET",
			},
			want:    "{\n          \"args\": {}, \n          \"headers\": {\n            \"Accept-Encoding\": \"gzip\", \n            \"Content-Type\": \"application/json\", \n            \"Host\": \"httpbin.org\", \n            \"User-Agent\": \"Go-http-client/2.0\", \n            \"X-Amzn-Trace-Id\": \"Root=1-61ab0318-5bf3fdee23e546927a6948bd\"\n          }, \n          \"origin\": \"59.91.153.204\", \n          \"url\": \"https://httpbin.org/get\"\n        }",
			wantErr: false,
		},
		//{
		//	name: "TestWebRequest",
		//	args: args{
		//		url:     "http://localhost:8080/test",
		//		method:  "POST",
		//		headers: nil,
		//		body:    "test",
		//	},
		//	//want:    []byte("{\"test\":\"test\"}"),
		//	wantErr: false,
		//},
		//{
		//	name: "TestWebRequest",
		//	args: args{
		//		url:     "http://localhost:8080/test",
		//		method:  "POST",
		//		headers: nil,
		//		body: map[string]interface{}{
		//			"test": "test",
		//		},
		//	},
		//	//want:    []byte("{\"test\":\"test\"}"),
		//	wantErr: false,
		//},
		//{
		//	name: "TestWebRequest",
		//	args: args{
		//		url:     "http://localhost:8080/test",
		//		method:  "POST",
		//		headers: nil,
		//		body: map[string]interface{}{
		//			"test": map[string]interface{}{
		//				"test": "test",
		//			},
		//		},
		//	},
		//	//want:    []byte("{\"test\":{\"test\":\"test\"}}"),
		//	wantErr: false,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := WebRequest(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("WebRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			var iGot = make(map[string]interface{})
			if err := json.NewDecoder(got.Body).Decode(&iGot); err != nil {
				panic(err)
			}
			if !reflect.DeepEqual(iGot["url"], tt.args.URL) {
				t.Errorf("WebRequest() got = %v, want %v", iGot["url"], tt.args.URL)
			}
		})
	}
}
