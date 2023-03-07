package middleware

import (
	"fmt"
	"testing"
)

func TestAuth(t *testing.T) {
	token, err := generateToken("test", "123456")
	fmt.Printf("token:%s, err:%+v\n", token, err)

	result, err := parseToken(token)
	fmt.Printf("claim:%+v, err:%+v\n", result, err)
}
