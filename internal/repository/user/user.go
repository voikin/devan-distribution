package user

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/voikin/devan-distribution/internal/entity"
)

type Repo struct {
	conn *pgx.Conn
}

func New(conn *pgx.Conn) *Repo {
	return &Repo{
		conn: conn,
	}
}

func (r *Repo) CreateUser(ctx context.Context, user entity.User) (int64, error) {
	var id int64
	row := r.conn.QueryRow(ctx, createUserQuery, user.Username, user.Password, user.Role.ID)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *Repo) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	rows, err := r.conn.Query(ctx, getUserByUsernameQuery, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user entity.User
	var role entity.Role
	permissionsMap := make(map[int64]entity.Permission)

	for rows.Next() {
		var permissionID *int64
		var permissionName *string

		err := rows.Scan(&user.ID, &user.Username, &user.Password, &role.ID, &role.Name, &permissionID, &permissionName)
		if err != nil {
			return nil, err
		}

		if permissionID != nil && permissionName != nil {
			if _, exists := permissionsMap[*permissionID]; !exists {
				permissionsMap[*permissionID] = entity.Permission{
					ID:   *permissionID,
					Name: *permissionName,
				}
			}
		}
	}

	role.Permissions = make([]entity.Permission, 0, len(permissionsMap))
	for _, perm := range permissionsMap {
		role.Permissions = append(role.Permissions, perm)
	}
	user.Role = role

	return &user, nil
}

func (r *Repo) GetUserByID(ctx context.Context, userId int64) (*entity.User, error) {
	rows, err := r.conn.Query(ctx, getUserByIDQuery, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user entity.User
	var role entity.Role
	permissionsMap := make(map[int64]entity.Permission)

	for rows.Next() {
		var permissionID *int64
		var permissionName *string

		err := rows.Scan(&user.ID, &user.Username, &user.Password, &role.ID, &role.Name, &permissionID, &permissionName)
		if err != nil {
			return nil, err
		}

		if permissionID != nil && permissionName != nil {
			if _, exists := permissionsMap[*permissionID]; !exists {
				permissionsMap[*permissionID] = entity.Permission{
					ID:   *permissionID,
					Name: *permissionName,
				}
			}
		}
	}

	role.Permissions = make([]entity.Permission, 0, len(permissionsMap))
	for _, perm := range permissionsMap {
		role.Permissions = append(role.Permissions, perm)
	}
	user.Role = role

	return &user, nil
}

func (r *Repo) UpdateUser(ctx context.Context, user entity.User) (bool, error) {
	commandTag, err := r.conn.Exec(ctx, updateUserQuery, user.Username, user.Password, user.Role.ID, user.ID)
	if err != nil {
		return false, err
	}

	return commandTag.RowsAffected() > 0, nil
}

func (r *Repo) DeleteUser(ctx context.Context, userId int64) (bool, error) {
	commandTag, err := r.conn.Exec(ctx, deleteUserQuery, userId)
	if err != nil {
		return false, err
	}

	return commandTag.RowsAffected() > 0, nil
}
