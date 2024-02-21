package api

type loginBody struct {
	Username string `json:username`
	Password string `json:password`
}

type errorStruct struct{
	Message string `json:message`
}
