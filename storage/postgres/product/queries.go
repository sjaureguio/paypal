package product

const (
	query     = "SELECT * FROM products"
	queryAll  = " ORDER BY name"
	queryByID = " WHERE id = $1"
)
