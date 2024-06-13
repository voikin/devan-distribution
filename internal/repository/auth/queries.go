package auth

const (
	getUserQuery = `select id from users where username = $1 and password_hash = $2`
)
