package auth

const (
	createUserQuery = `insert into users (username,password_hash) values ($1, $2) returning id`
	getUserQuery    = `select id from users where username = $1 and password_hash = $2`
)
