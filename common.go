package go_jwt

var separator = "."

/**
https://tools.ietf.org/html/rfc7519#section-4.1
*/
type Headers struct {
	Exp int64  `json:"exp,omitempty"`
	Iat int64  `json:"iat,omitempty"`
	Nbf int64  `json:"nbf,omitempty"`
	Alg string `json:"alg,omitempty"`
	Aud string `json:"aud,omitempty"`
	Iss string `json:"iss,omitempty"`
	Jti string `json:"jti,omitempty"`
	Sub string `json:"sub,omitempty"`
}
