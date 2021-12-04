package binance

import (
	"reflect"
	"testing"
)

func TestNewClient(t *testing.T) {
	type args struct {
		apiKey    string
		secretKey string
	}
	tests := []struct {
		name string
		args args
		want *Client
	}{
		{
			name: "NewClient",
			args: args{
				apiKey:    "",
				secretKey: "",
			},
		},
		{
			name: "NewClient",
			args: args{
				apiKey:    "apiKey",
				secretKey: "secretKey",
			},
		},
		{
			name: "NewClient",
			args: args{
				apiKey:    "rnsgAhtXR4vsXBCykL3ObO6dsnBRzEmGfdDbih9Y6cwIIhgki9YbmfYU9BCfD585",
				secretKey: "76q4sw5w4L5q8GE9UVhKh0F1PtqtYh4f3vn4QY0yc35MSewGE7ve9gE4AT8VWlLe",
			},
		},
		{
			name: "NewClient",
			args: args{
				apiKey:    "",
				secretKey: "secretKey",
			},
		},
		{
			name: "NewClient",
			args: args{
				apiKey:    "apiKey",
				secretKey: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(tt.args.apiKey, tt.args.secretKey); !reflect.DeepEqual(got.APIKey, tt.args.apiKey) || !reflect.DeepEqual(got.SecretKey, tt.args.secretKey) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
