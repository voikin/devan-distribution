package user

const (
	createUserQuery        = `INSERT INTO users (username, password_hash, role_id) VALUES ($1, $2, $3) RETURNING id`
	getUserByUsernameQuery = `
		SELECT u.id, u.username, u.password_hash, r.id, r.name, p.id, p.name
		FROM users u
		LEFT JOIN roles r ON u.role_id = r.id
		LEFT JOIN role_permissions rp ON r.id = rp.role_id
		LEFT JOIN permissions p ON rp.permission_id = p.id
		WHERE u.username = $1`
	getUserByIDQuery = `
		SELECT u.id, u.username, u.password_hash, r.id, r.name, p.id, p.name
		FROM users u
		LEFT JOIN roles r ON u.role_id = r.id
		LEFT JOIN role_permissions rp ON r.id = rp.role_id
		LEFT JOIN permissions p ON rp.permission_id = p.id
		WHERE u.id = $1`
	updateUserQuery = `UPDATE users SET username=$1, password_hash=$2, role_id=$3 WHERE id=$4`
	deleteUserQuery = `DELETE FROM users WHERE id=$1`
	selectRolesSQL  = `SELECT id, name FROM roles;`
)
