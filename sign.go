package go_jwt

import (
	"crypto/hmac"
	"crypto/sha256"
)

func sign(h []byte, p []byte, secret []byte) []byte {
	jwtHmacHash := hmac.New(sha256.New, secret)
	jwtHmacHash.Write(h)
	jwtHmacHash.Write([]byte(separator))
	jwtHmacHash.Write(p)

	return encodeBase64(jwtHmacHash.Sum(nil))
}
