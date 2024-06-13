package orders

const (
	createOrderQuery = `
INSERT INTO orders (
    id_pas, datetime, time3, time4, cat_pas, status_id, tpz, insp_sex_m, insp_sex_f, time_over, id_st1, id_st2
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
) RETURNING id;`

	getOrderQuery    = `SELECT * FROM orders WHERE id = $1;`
	updateOrderQuery = `
UPDATE orders SET
    id_pas = $1, datetime = $2, time3 = $3, time4 = $4, cat_pas = $5, status_id = $6, tpz = $7, insp_sex_m = $8, insp_sex_f = $9, time_over = $10, id_st1 = $11, id_st2 = $12
WHERE id = $13;`

	deleteOrderQuery = `DELETE FROM orders WHERE id = $1;`
	listOrdersQuery  = `SELECT * FROM orders;`
)

const (
	createStatusQuery = `INSERT INTO statuses (name) VALUES ($1) RETURNING id;`
	getStatusQuery    = `SELECT * FROM statuses WHERE id = $1;`
	listStatusesQuery = `SELECT * FROM statuses;`
)
