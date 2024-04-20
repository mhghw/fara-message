package api

type loginBody struct {
	Username string `json:username`
	Password string `json:password`
}

type HTTPError struct{
	Message string `json:message`
}
