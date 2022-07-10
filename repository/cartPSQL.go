package repository

import "Book_Cart_Project/models"

func (r Repository) AddToCart(cart models.Cart) (models.Cart, error) {
	//writing query
	query := `INSERT INTO cart(
		user_id,
		product_id,
		product_count)
		VALUES ($1, $2, $3)
		RETURNING
		cart_id,
		user_id,
		product_id,
		product_count,
		cart_created_at;`

	err := r.DB.QueryRow(
		query,
		cart.User_ID,
		cart.Product_ID.Product_ID,
		cart.Product_Count,
	).Scan(
		&cart.Cart_ID,
		&cart.User_ID,
		&cart.Product_ID.Product_ID,
		&cart.Product_Count,
		&cart.Cart_Created_At,
	)

	return cart, err
}

func (r Repository) ViewCart(cart models.Cart) ([]models.Cart, error) {

	query := `SELECT 
		cart.cart_id, 
		cart.user_id, 
		product.product_id,
		product.product_name,
		cart.product_count
		FROM product
		INNER JOIN cart ON product.product_id=  cart.product_id
		WHERE user_id = $1;`

	rows, err := r.DB.Query(query,
		cart.User_ID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var carts []models.Cart

	for rows.Next() {
		err := rows.Scan(
			&cart.Cart_ID,
			&cart.User_ID,
			&cart.Product_ID.Product_ID,
			&cart.Product_ID.Product_Name,
			&cart.Product_Count,
		)
		if err != nil {
			return carts, err
		}
		carts = append(carts, cart)
	}
	if err = rows.Err(); err != nil {
		return carts, err
	}

	return carts, nil

}
