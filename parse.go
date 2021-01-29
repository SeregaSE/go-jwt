package go_jwt

import (
	"errors"
	"strings"
)

func parse(jwt string) ([]byte, []byte, []byte, error) {
	parts := strings.Split(jwt, ".")

	if len(parts) != 3 {
		return []byte(""), []byte(""), []byte(""), errors.New("failed to parse jwt token")
	}

	return []byte(parts[0]), []byte(parts[1]), []byte(parts[2]), nil
}
