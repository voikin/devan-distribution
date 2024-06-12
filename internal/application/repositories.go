package application

import (
	"github.com/jackc/pgx/v5"
	userRepo "github.com/voikin/devan-distribution/internal/repository/auth"
)

type repositories struct {
	userRepository *userRepo.Repo
}

func createRepository(conn *pgx.Conn) *repositories {
	return &repositories{
		userRepository: userRepo.New(conn),
	}
}
