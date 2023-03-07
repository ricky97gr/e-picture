package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var JwtStr = []byte("这是jwt认证密钥")

type Claims struct {
	UserName string
	Phone    string
	Role     string
	jwt.StandardClaims
}

func Auth(ctx *gin.Context) {

}

func generateToken(userName, phone string) (string, error) {
	claim := &Claims{
		UserName: userName,
		Phone:    phone,
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString(JwtStr)
}

func parseToken(token string) (*Claims, error) {
	tokenClaim, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return JwtStr, nil
	})
	if _, ok := tokenClaim.Claims.(*Claims); ok {
		if tokenClaim.Claims.Valid() == nil {
			return tokenClaim.Claims.(*Claims), nil
		}
	}
	return nil, err
}
