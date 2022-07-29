package repository

import (
	"Book_Cart_Project/models"
	"context"
	"database/sql"
	"log"
)

func (r Repository) AddToCart(cart models.Cart) (models.Cart, error) {

	//------------------implementing sql transactions for updating or adding products into cart------------------

	//Begin a transaction
	ctx := context.Background()
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return cart, err
	}

	//execute sql queries against transaction
	//checking for the product present int the cart, if not inserting and if yes updating
	query1 := `SELECT cart_id 
				FROM cart
				WHERE 
				user_id = $1 
				AND 
				product_id =$2;`
	err = tx.QueryRow(
		query1,
		cart.User_ID,
		cart.Product_ID.Product_ID,
	).Scan(
		&cart.Cart_ID,
	)

	if err == sql.ErrNoRows && err != nil { //no iteam found in the cart
		//inserting the new iteam into the cart
		query3 := `INSERT INTO cart(
			user_id,
			product_id,
			product_count)
			VALUES ($1, $2, $3)
			RETURNING
			cart_id,
			user_id,
			product_id,
			product_count;`

		err := tx.QueryRow(
			query3,
			cart.User_ID,
			cart.Product_ID.Product_ID,
			cart.Product_Count,
		).Scan(
			&cart.Cart_ID,
			&cart.User_ID,
			&cart.Product_ID.Product_ID,
			&cart.Product_Count,
		)
		log.Println("updated the query")
		if err != nil {
			log.Println("add cart insert failed")
			return cart, err
		}

	} else { //if the product is available in the cart table
		//updating the cart
		query2 := `UPDATE cart
					SET product_count = product_count + $1
					WHERE cart_id = $2
					RETURNING 
					product_count,
					cart_id;`

		err = tx.QueryRow(
			query2,
			cart.Product_Count,
			cart.Cart_ID,
		).Scan(
			&cart.Product_Count,
			&cart.Cart_ID,
		)

		if err != nil {
			tx.Rollback()
			return cart, err
		}
	}

	//queries worked prefectly and either updating or inserting the product
	err = tx.Commit() //commit will make the changes in the cart table
	if err != nil {
		log.Println("commit failed")
		return cart, err
	}
	log.Println("commit sucess")
	return cart, err

}

func (r Repository) ViewCart(cart models.Cart) ([]models.Cart, float64, error) {

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
		return nil, 0, err
	}

	defer rows.Close()

	var carts []models.Cart
	var totalPrice float64

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
			return carts, totalPrice, err
		}
		totalPrice += (cart.Product_ID.Product_Price * float64(cart.Product_Count))
		carts = append(carts, cart)
	}
	if err = rows.Err(); err != nil {
		return carts, totalPrice, err
	}

	return carts, totalPrice, nil

}

//need to create a discount function
