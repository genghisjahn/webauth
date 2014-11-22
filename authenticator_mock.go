package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
)

type HashSaltHolder struct {
	Hash string
	Salt []byte
}

var UserAuthData map[string]HashSaltHolder

func (a *Authenticate) AuthenticateUserNamePassword(userName string, password string) (AuthResponse, error) {
	setHashSaltHolder()
	authResp := AuthResponse{}
	salt := make([]byte, 0, 64)
	pw := []byte(password)
	if storedHash, ok := UserAuthData[userName]; ok {
		hash := HashPasswordWithSalt(pw, salt)
		if hash == storedHash.Hash {
			authResp.GUID = "1001001SOS"
			authResp.LoggedIn = true
			authResp.UserName = userName
			authResp.Code = AUTH_SUCCESS
		} else {
			authResp.Code = AUTH_INVALID_PASSWORD
			errPW := new(InvalidPasswordError)
			return authResp, errPW
		}
	} else {
		authResp.Code = AUTH_INVALID_USERNAME
		errUN := new(InvalidPasswordError)
		return authResp, errUN
	}
	return authResp, nil
}

func setHashSaltHolder() {
	UserAuthData = make(map[string]HashSaltHolder)
	salt := make([]byte, 0, 64)
	password := []byte("password")
	holder := HashSaltHolder{}
	holder.Salt = salt
	holder.Hash = HashPasswordWithSalt(password, salt)
	UserAuthData["test@test.com"] = holder
}

func HashPasswordWithSalt(password []byte, salt []byte) string {
	h := hmac.New(sha512.New, salt)
	h.Write(password)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
