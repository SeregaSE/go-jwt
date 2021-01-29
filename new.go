package go_jwt

import (
	"encoding/json"
	"fmt"
)

func New(headers *Headers, payload interface{}, secret []byte) (string, error) {
	rawHeaders, err := json.Marshal(headers)

	if err != nil {
		return "", err
	}

	base64Headers := encodeBase64(rawHeaders)

	rawPayload, err := json.Marshal(payload)

	if err != nil {
		return "", err
	}

	base64Payload := encodeBase64(rawPayload)

	return fmt.Sprintf("%s.%s.%s", base64Headers, base64Payload, sign(base64Headers, base64Payload, secret)), nil
}
