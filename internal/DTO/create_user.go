package DTO

type CreateUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	RoleID   int    `json:"role"`
}
