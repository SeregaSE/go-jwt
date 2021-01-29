package go_jwt

import (
	"testing"
	"time"
)

func TestVerifyParseError(t *testing.T) {
	secret := []byte("qwerty123456")
	jwt := "hello world!"

	_, err := Verify(jwt, secret, nil)

	if _, ok := err.(*ParseError); !ok {
		t.Error(err)
	}
}

func TestVerifyInvalidErrorChangeJWT(t *testing.T) {
	secret := []byte("qwerty123456")
	jwt := "eyJleHAiOjE2MTE5MTg1ODh9.aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.YPoqj6h_n6OyhHhIFMB3G_dIvxsUgK-PSJmcmjGIjbc"

	_, err := Verify(jwt, secret, nil)

	if _, ok := err.(*InvalidError); !ok {
		t.Error(err)
	}
}

func TestVerifyInvalidErrorChangeSecret(t *testing.T) {
	secret := []byte("qwerty777777")
	jwt := "eyJleHAiOjE2MTE5MTg1ODh9.eyJJZCI6MSwiUm9sZSI7InJvb3QifQ.YPoqj6h_n6OyhHhIFMB3G_dIvxsUgK-PSJmcmjGIjbc"

	_, err := Verify(jwt, secret, nil)

	if _, ok := err.(*InvalidError); !ok {
		t.Error(err)
	}
}

func TestVerifyExpError(t *testing.T) {
	secret := []byte("qwerty123456")
	jwt := "eyJleHAiOjE2MTE5MTg1ODh9.eyJJZCI6MSwiUm9sZSI6InJvb3QifQ.YPoqj6h_n6OyhHhIFMB3G_dIvxsUgK-PSJmcmjGIjbc"

	_, err := Verify(jwt, secret, nil)

	if _, ok := err.(*ExpError); !ok {
		t.Error(err)
	}
}

func TestVerifyNbfError(t *testing.T) {
	type Payload struct {
		Id   int
		Role string
	}

	secret := []byte("qwerty123456")

	headers := &Headers{
		Exp: time.Now().Add(time.Minute).Unix(),
		Nbf: time.Now().Add(time.Second * 10).Unix(),
	}

	jwt, err := New(headers, &Payload{1, "root"}, secret)

	if err != nil {
		t.Error(err)
	}

	_, err = Verify(jwt, secret, nil)

	if _, ok := err.(*NbfError); !ok {
		t.Error(err)
	}
}

func TestVerifySuccess(t *testing.T) {
	type Payload struct {
		Id   int
		Role string
	}

	secret := []byte("qwerty123456")

	jwt, err := New(&Headers{Exp: time.Now().Add(time.Minute).Unix()}, &Payload{1, "root"}, secret)

	if err != nil {
		t.Error(err)
	}

	payload := &Payload{}

	_, err = Verify(jwt, secret, payload)

	if err != nil {
		t.Error(err)
	}

	if payload.Id != 1 {
		t.Error("user id must be 1")
	}

	if payload.Role != "root" {
		t.Error("user role must be root")
	}
}
