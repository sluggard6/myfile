package service

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/sluggard/myfile/model"
)

// type TokenService interface {
// 	NewToken(user *model.User) (token string, err error)
// 	CheckToken(token string) (user *model.User, err error)
// }

// var tokenService = &tokenServiceImpl{
// 	tokenMap: make(map[string]tokenInfo),
// }

// func NewTokenService() TokenService {
// 	return tokenService
// }

// type tokenServiceImpl struct {
// 	tokenMap map[string]tokenInfo
// }

type TokenService struct {
	tokenMap map[string]tokenInfo
}

var tokenService = &TokenService{
	tokenMap: make(map[string]tokenInfo),
}

type tokenInfo struct {
	info   interface{}
	exTime time.Time
}

type TokenCheckError struct {
	message string
}

func (e *TokenCheckError) Error() string {
	return e.message
}

func (t *TokenService) NewToken(user *model.User) (token string, err error) {
	token = t.tokenValue()
	// t.tokenMap[token]
	t.tokenMap[token] = tokenInfo{
		info:   user,
		exTime: time.Now().Add(time.Minute * 30),
	}
	return
}
func (t *TokenService) CheckToken(token string) (user *model.User, err error) {
	ti, ok := t.tokenMap[token]
	if ok {
		if ti.exTime.After(time.Now()) {
			user = ti.info.(*model.User)
		} else {
			err = &TokenCheckError{"token expired"}
			delete(t.tokenMap, token)
		}
	} else {
		err = &TokenCheckError{"token not found"}
	}
	return
}

func (t *TokenService) tokenValue() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
