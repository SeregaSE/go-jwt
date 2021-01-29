package go_jwt

import "encoding/base64"

func encodeBase64(src []byte) []byte {
	dst := make([]byte, base64.RawURLEncoding.EncodedLen(len(src)))
	base64.RawURLEncoding.Encode(dst, src)
	return dst
}

func decodeBase64(src []byte) []byte {
	dst := make([]byte, base64.RawURLEncoding.DecodedLen(len(src)))
	base64.RawURLEncoding.Decode(dst, src)
	return dst
}
