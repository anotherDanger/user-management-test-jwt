package domain

type Domain struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	FullName string `json:"fullname"`
	Password string `json:"password"`
}
