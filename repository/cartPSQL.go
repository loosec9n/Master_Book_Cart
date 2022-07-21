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
		product_count;`

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
	)

	return cart, err
}

func (r Repository) ViewCart(cart models.Cart) ([]models.Cart, error) {

	query := `SELECT 
	cart.cart_id, 
	cart.user_id, 
	product.product_id,
	product.product_name,
	product.product_description,
	product.product_price,
	product.product_rating,
	product_category.category_id,
	product_category.category_name,
	product_category.category_description,
	product_author.author_id,
	product_author.author_name,
	cart.product_count
	FROM product
	INNER JOIN 
	cart 
	ON 
	product.product_id = cart.product_id
	INNER JOIN 
	product_category 
	ON 
	product_category.category_id = product.product_category_id
	INNER JOIN
	product_author
	ON 
	product_author.author_id = product.product_author_id
	WHERE 
	user_id = $1;`

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
			&cart.Product_ID.Product_Description,
			&cart.Product_ID.Product_Price,
			&cart.Product_ID.Product_Rating,
			&cart.Product_ID.Product_Category.Category_ID,
			&cart.Product_ID.Product_Category.Category_Name,
			&cart.Product_ID.Product_Category.Category_Description,
			&cart.Product_ID.Product_Author.Author_ID,
			&cart.Product_ID.Product_Author.Author_Name,
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
