package token

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

var ErrMissingHeader = errors.New(`header "authorization" is missing`)

type Config struct{
	key string
	identityKey string
}

var(
	once sync.Once
	config = Config{key: "", identityKey: ""}
)

func Init(key string, identityKey string) {
	once.Do(func(){
		config.key = key
		config.identityKey = identityKey
	})
}

func Parse(tokenString string, key string) (string, error) {
	t, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error){
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(key), nil
	})
	if err != nil {
		return "", err
	}
	
	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		return claims[config.identityKey].(string), nil
	}
	
	return "", nil
}

func ParseRequest(c *gin.Context) (string, error) {
	a := c.Request.Header.Get("Authorization")
	if a=="" {
		return "", ErrMissingHeader
	}

	var t string
	fmt.Sscanf(a, "Bear %s", &t)
	return Parse(t, config.key)
}

func Sign(identity string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		config.identityKey: identity,
		"nbf": time.Now(),
		"iat": time.Now(),
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	signed, err := t.SignedString([]byte(config.key))
	return signed, err
}