package main

func AuthenticateUserNamePassword(userName string, password string) (AuthResponse, error) {
	authResp := AuthResponse{}
	authResp.GUID = ""
	authResp.LoggedIn = false
	authResp.UserName = ""
	return authResp, nil
}
