package orders

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/voikin/devan-distribution/internal/entity"
	"github.com/voikin/devan-distribution/internal/errs"
)

type Repo struct {
	conn *pgx.Conn
}

func New(conn *pgx.Conn) *Repo {
	return &Repo{
		conn: conn,
	}
}

func (r *Repo) CreateOrder(ctx context.Context, order entity.Order) (int64, error) {
	var id int64
	err := r.conn.QueryRow(ctx, createOrderQuery, order.IDPas, order.DateTime, order.Time3, order.Time4, order.CatPas, order.StatusID, order.TPZ, order.InspSexM, order.InspSexF, order.TimeOver, order.IDSt1, order.IDSt2).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *Repo) GetOrderById(ctx context.Context, id int64) (*entity.Order, error) {
	var order entity.Order
	err := r.conn.QueryRow(ctx, getOrderQuery, id).Scan(&order.ID, &order.IDPas, &order.DateTime, &order.Time3, &order.Time4, &order.CatPas, &order.StatusID, &order.TPZ, &order.InspSexM, &order.InspSexF, &order.TimeOver, &order.IDSt1, &order.IDSt2)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, errs.NewErrorOrderNotFound(id)
		}
		return nil, err
	}
	return &order, nil
}

func (r *Repo) UpdateOrder(ctx context.Context, order entity.Order) error {
	_, err := r.conn.Exec(ctx, updateOrderQuery, order.IDPas, order.DateTime, order.Time3, order.Time4, order.CatPas, order.StatusID, order.TPZ, order.InspSexM, order.InspSexF, order.TimeOver, order.IDSt1, order.IDSt2, order.ID)
	return err
}

func (r *Repo) DeleteOrder(ctx context.Context, id int64) error {
	_, err := r.conn.Exec(ctx, deleteOrderQuery, id)
	return err
}

func (r *Repo) ListOrders(ctx context.Context) ([]entity.Order, error) {
	rows, err := r.conn.Query(ctx, listOrdersQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []entity.Order
	for rows.Next() {
		var order entity.Order
		err := rows.Scan(&order.ID, &order.IDPas, &order.DateTime, &order.Time3, &order.Time4, &order.CatPas, &order.StatusID, &order.TPZ, &order.InspSexM, &order.InspSexF, &order.TimeOver, &order.IDSt1, &order.IDSt2)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}
