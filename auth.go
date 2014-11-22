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

type InvalidUserNameError struct {
}
type InvalidPasswordError struct {
}

func (u *InvalidUserNameError) Error() string {
	return "Invalid username."
}
func (u *InvalidPasswordError) Error() string {
	return "Invalid password."
}
