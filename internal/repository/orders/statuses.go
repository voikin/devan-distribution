package orders

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/voikin/devan-distribution/internal/entity"
	"github.com/voikin/devan-distribution/internal/errs"
)

func (r *Repo) CreateStatus(ctx context.Context, status entity.Status) (int64, error) {
	var id int64
	err := r.conn.QueryRow(ctx, createStatusQuery, status.Name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *Repo) GetStatusById(ctx context.Context, id int64) (*entity.Status, error) {
	var status entity.Status
	err := r.conn.QueryRow(ctx, getStatusQuery, id).Scan(&status.ID, &status.Name)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errs.NewErrorStatusNotFound(id)
		}
		return nil, err
	}
	return &status, nil
}

func (r *Repo) ListStatuses(ctx context.Context) ([]entity.Status, error) {
	rows, err := r.conn.Query(ctx, listStatusesQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var statuses []entity.Status
	for rows.Next() {
		var status entity.Status
		err := rows.Scan(&status.ID, &status.Name)
		if err != nil {
			return nil, err
		}
		statuses = append(statuses, status)
	}
	return statuses, nil
}
