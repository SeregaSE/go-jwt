package go_jwt

import (
	"testing"
)

func TestNew(t *testing.T) {
	h := &Headers{Exp: 1611918588}

	payload := struct {
		Id   int
		Role string
	}{1, "user"}

	jwt, err := New(h, payload, []byte("qwerty123456"))

	if err != nil {
		t.Error(err)
	}

	want := "eyJleHAiOjE2MTE5MTg1ODh9.eyJJZCI6MSwiUm9sZSI6InVzZXIifQ.LXSsKbXT28aHeHRIxYHLVdHMgtTpbtwyuxi_cTNOKOY"

	if jwt != want {
		t.Errorf("expect: %s, got: %s", want, jwt)
	}
}
