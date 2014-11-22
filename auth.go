package main

type Authenticator interface {
	AuthenticateUserNamePassword(userName string, password string) (AuthResponse, error)
}

type Authenticate struct {
}

func ProcessAuthentication(userName string, password string) (AuthResponse, error) {
	auth := Authenticate{}
	return auth.AuthenticateUserNamePassword(userName, password)
}
