package go_jwt

import (
	"bytes"
	"encoding/json"
	"time"
)

func Verify(jwt string, secret []byte, payload interface{}) (*Headers, error) {
	base64Headers, base64Payload, base64Sign, err := parse(jwt)

	if err != nil {
		return nil, err
	}

	if !bytes.Equal(base64Sign, sign(base64Headers, base64Payload, secret)) {
		return nil, &InvalidError{}
	}

	headers := &Headers{}

	err = json.Unmarshal(decodeBase64(base64Headers), headers)

	if err != nil {
		return nil, err
	}

	if headers.Exp > 0 && time.Now().Unix() > headers.Exp {
		return headers, &ExpError{}
	}

	if headers.Nbf > 0 && time.Now().Unix() < headers.Nbf {
		return headers, &NbfError{}
	}

	err = json.Unmarshal(decodeBase64(base64Payload), payload)

	if err != nil {
		return headers, err
	}

	return headers, nil
}
