package entity

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     Role   `json:"role"`
}

type Role struct {
	ID          int64        `json:"id"`
	Name        string       `json:"name"`
	Permissions []Permission `json:"permissions"`
}

type Permission struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
