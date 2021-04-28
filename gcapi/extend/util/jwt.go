package util

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWT struct {
}

func (j *JWT) Make(name,pwd string,expire time.Duration) (string, error) {
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":  name,
		"exp":  time.Now().Add(expire * time.Second).Unix(),
	})
	token, err := at.SignedString([]byte(pwd))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (j *JWT) Check(token,name,pwd string) error {
	claim, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(pwd), nil
	})
	if err != nil {
		return  err
	}
	tname := claim.Claims.(jwt.MapClaims)["name"].(string)
	if tname == name {
		return nil
	}
	return errors.New("invalid account")
}
