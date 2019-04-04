package util

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/astaxie/beego"
	jwt "github.com/dgrijalva/jwt-go"
)

func GetToken(name string) (string, error) {
	addtime, err := beego.AppConfig.Int("token_inter")
	if err != nil {
		addtime = 1
	}
	key := beego.AppConfig.String("token_key")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	nowtime := time.Now()
	claims["iss"] = "wujq"
	claims["iat"] = nowtime.Unix()
	claims["exp"] = nowtime.Add(time.Hour * time.Duration(addtime)).Unix()
	claims["sub"] = name
	claims["nbf"] = nowtime.Unix()
	claims["jti"] = GetGernate()
	token.Claims = claims
	tokenstring, err := token.SignedString([]byte(key))
	return tokenstring, err

}

func ParseToken(token string) (bool, string) {
	key := beego.AppConfig.String("token_key")
	tok, err := jwt.Parse(token, func(tk *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return false, ""
	}
	if claims, ok := tok.Claims.(jwt.MapClaims); ok && tok.Valid {

		return true, claims["sub"].(string)
	}
	return false, ""

}

func Md5Password(pd string) string {
	hash := md5.New()
	hash.Write([]byte(pd))
	return hex.EncodeToString(hash.Sum(nil))

}
