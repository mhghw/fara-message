package api

type loginBody struct {
	Username string `json:username`
	Password string `json:password`
}
