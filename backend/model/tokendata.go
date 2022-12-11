package model

type TokenUserData struct {
	User
	RoleName string `json:"role_name"`
}
