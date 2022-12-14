package model

type TokenUserData struct {
	Username string  `json:"username"`
	Issuer   string  `json:"issuer"`
	RoleID   float64 `json:"role_id"`
	RoleName string  `json:"role_name"`
}
