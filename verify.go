package go_jwt

import (
	"bytes"
)

func Verify(jwt string, secret []byte) (bool, error) {
	rawHeaders, rawPayload, rawSign, err := parse(jwt)

	if err != nil {
		return false, err
	}

	return bytes.Equal(rawSign, sign(rawHeaders, rawPayload, secret)), nil
}
