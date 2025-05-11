package web

type Response struct {
	Username string `json:"username"`
	FullName string `json:"fullname"`
	Token    string `json:"token"`
}
