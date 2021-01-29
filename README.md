# Golang JWT

Implementation is simple.
Signing and verifying JWT via hmac and sha256.
Validate JWT by `Headers.Exp` and `Headers.Nbf` fields.
Use `base64.RawURLEncoding` to make JWT safe to use in web (urls).

## New JWT

```golang
package main

import (
	"fmt"
	"log"
	"time"
	jwt "github.com/SeregaSE/go-jwt"
)

type Payload struct {
	Id int
	Role string
}

var secret = []byte("qwerty123456")

func main() {
	token, err := jwt.New(
		&jwt.Headers{
			Exp: time.Now().Add(time.Hour).Unix(),
			Nbf: time.Now().Add(time.Second * 10).Unix(),
		},
		&Payload{1, "root"},
		secret,
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(token)
}
```

## Verify JWT

If JWT is ok

```golang
package main

import (
	"fmt"
	"log"
	"time"
	jwt "github.com/SeregaSE/go-jwt"
)

type Payload struct {
	Id int
	Role string
}

var secret = []byte("qwerty123456")

func main() {
	token, err := jwt.New(
		&jwt.Headers{ Exp: time.Now().Add(time.Hour).Unix() },
		&Payload{1, "root"},
		secret,
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(token)

	payload := &Payload{}

	_, err = jwt.Verify(token, secret, &payload)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", payload)
}
```

If JWT is not ok
```golang
package main

import (
	"fmt"
	"log"
	"time"
	jwt "github.com/SeregaSE/go-jwt"
)

type Payload struct {
	Id int
	Role string
}

var secret = []byte("qwerty123456")

func main() {
    token := "eyJleHAiOjE2MTE5MTg1ODh9.eyJJZCI6MSwiUm9sZSI6InJvb3QifQ.YPoqj6h_n6OyhHhIFMB3G_dIvxsUgK-PSJmcmjGIjbc"
    
    payload := &Payload{}
    
    _, err := jwt.Verify(token, secret, &payload)
    
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("%v\n", payload)
}
```

## TODO

- Doc verify errors
- Sign JWT via RSA alg
