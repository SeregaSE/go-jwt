# Golang JWT

Implementation is simple. Signing and verifying JWT via hmac and sha256. Validate JWT by `Headeers.Exp` and `Headers.Nbf` fields.

## New JWT

```golang
type Payload struct {
	Id   int
	Role string
}

secret := []byte("qwerty123456")

headers := &Headers{
    Exp: time.Now().Add(time.Hour).Unix(),
    Nbf: time.Now().Add(time.Minute).Unix(),
}

payload := &Payload{1, "root"}

jwt, err := New(headers, payload, secret)

if err != nil {
	fmt.Println(err)
}

fmt.Println(jwt)
```

## Verify JWT

If JWT is ok

```golang
type Payload struct {
	Id   int
	Role string
}

secret := []byte("qwerty123456")

jwt := "eyJleHAiOjE2MTE5MTg1ODh9.eyJJZCI6MSwiUm9sZSI6InJvb3QifQ.YPoqj6h_n6OyhHhIFMB3G_dIvxsUgK-PSJmcmjGIjbc"

payload := &Payload{}

_, err := Verify(jwt, secret, &payload)

if err != nil {
	fmt.Println(err)
}

// payload.Id === 1

```

## TODO

- Doc verify errors
- Sign JWT via RSA alg
