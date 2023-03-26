package middleware

import (
	"my-admin/global"
	"my-admin/pkg/response"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var JwtStr = []byte("这是jwt认证密钥")

const (
	expiration = 60 * time.Minute
)

type Claims struct {
	UserName string
	Phone    string
	Role     string
	jwt.StandardClaims
}

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token == "" {
			response.Failed(ctx, response.ErrAuth)
			return
		}
		if !isTokenExist(token) {
			response.Failed(ctx, response.ErrAuth)
			return
		}
		if restoreToken(token) != nil {
			response.Failed(ctx, response.ErrRedis)
		}
		claims, err := parseToken(token)
		if err != nil {
			global.Logger.Errorf("failed to parse token, err:%+v\n", err)
		}
		ctx.Request.Header.Set("userName", claims.UserName)
		ctx.Request.Header.Set("role", claims.Role)
		ctx.Request.Header.Set("userID", claims.Phone)
		global.Logger.Infof("user:%s, auth successfully", claims.UserName)
		ctx.Next()

	}
}

func GenerateToken(userName, phone string) (string, error) {
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

func StoreToken(token string) error {
	return global.RedisClient.Set(token, nil, expiration).Err()
}

func restoreToken(token string) error {
	return StoreToken(token)
}

func isTokenExist(token string) bool {
	return global.RedisClient.Get(token).Err() == nil
}
