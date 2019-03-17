package models

type AuthReqest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
