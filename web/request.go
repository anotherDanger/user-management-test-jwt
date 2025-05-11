package web

type Request struct {
	Username string `json:"username"`
	FullName string `json:"fullname"`
	Password string `json:"password"`
}
