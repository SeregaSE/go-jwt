package go_jwt

import (
	"testing"
)

func TestNew(t *testing.T) {
	secret := []byte("qwerty123456")
	want := "eyJleHAiOjE2MTE5MTg1ODh9.eyJJZCI6MSwiUm9sZSI6InJvb3QifQ.YPoqj6h_n6OyhHhIFMB3G_dIvxsUgK-PSJmcmjGIjbc"

	type Payload struct {
		Id   int
		Role string
	}

	headers := &Headers{Exp: 1611918588}
	payload := &Payload{1, "root"}

	jwt, err := New(headers, payload, secret)

	if err != nil {
		t.Error(err)
	}

	if jwt != want {
		t.Errorf("expect: %s, got: %s", want, jwt)
	}
}
