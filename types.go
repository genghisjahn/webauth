package main

const AUTH_SUCCESS = 0
const AUTH_INVALID_USERNAME = 1
const AUTH_INVALID_PASSWORD = 2

type AuthResponse struct {
	GUID     string
	UserName string
	LoggedIn bool
	Code     int
}
