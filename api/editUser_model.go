package api

type UsernameStruct struct {
	Username string `json:"username"`
}

type Information struct {
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
