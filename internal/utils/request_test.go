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
	type args struct {
		url     string
		method  string
		headers map[string]string
		body    interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "TestWebRequest",
			args: args{
				url:     "https://httpbin.org/get",
				method:  "GET",
				headers: nil,
				body:    nil,
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
			got, err := WebRequest(tt.args.url, tt.args.method, tt.args.headers, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("WebRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			var _i_got = make(map[string]interface{})
			//var _i_want = make(map[string]interface{})
			if err := json.Unmarshal(got, &_i_got); err != nil {
				panic(err)
			}
			//if err := json.Unmarshal([]byte(tt.want), &_i_want); err != nil {
			//	panic(err)
			//}
			if !reflect.DeepEqual(_i_got["url"], tt.args.url) {
				t.Errorf("WebRequest() got = %v, want %v", _i_got["url"], tt.args.url)
			}
		})
	}
}
