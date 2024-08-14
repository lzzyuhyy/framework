package jwt

import (
	"framework/nacos"
	"github.com/golang-jwt/jwt/v4"
	"gopkg.in/yaml.v2"
	"time"
)

type Config struct {
	Jwt
}

type Jwt struct {
	Secret string `yaml:"secret"`
	Dura   int64  `yaml:"dura"`
}

var conf Config

func getConfig() error {
	config, err := nacos.GetNacosConfig()
	if err != nil {
		return err
	}

	err = yaml.Unmarshal([]byte(config), &conf)
	if err != nil {
		return err
	}
	return nil
}

func GenerateJwtToken(payload string) (string, error) {
	err := getConfig()
	if err != nil {
		return "", err
	}
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Unix() + conf.Jwt.Dura
	claims["iat"] = time.Now().Unix()
	claims["payload"] = payload
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(conf.Jwt.Secret))
}

func ParseJwtToken(tokenString string, payload *string) bool {
	err := getConfig()
	if err != nil {
		return false
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.Jwt.Secret), nil
	})
	if err != nil {
		return false
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		*payload = claims["payload"].(string)
		return true
	}
	return false
}
