package repository

import (
	"Book_Cart_Project/models"
	"context"
	"log"
)

func (r Repository) CreateNewOrder(order models.OrderBody, orderIn models.Order) ([]models.Order, float64, error) {
	var orderOut []models.Order

	query := `SELECT 
	product.product_name, 
	cart.product_count, 
	users.first_name,
	users.last_name, 
	users.user_address_id, 
	product.product_price
	FROM cart
	INNER JOIN product ON cart.product_id = product.product_id
	INNER JOIN users ON cart.user_id = users.user_id
	WHERE cart.user_id = $1;`

	rows, err := r.DB.Query(query, order.User_ID)
	if err != nil {
		return orderOut, 0, err
	}
	defer rows.Close()
	var totalAmount float64
	for rows.Next() {
		err := rows.Scan(
			&orderIn.Products.Product_Name,
			&orderIn.Cart_ID.Product_Count,
			&orderIn.User_ID.First_Name,
			&orderIn.User_ID.Last_Name,
			&orderIn.Address_ID.AddressID,
			&orderIn.Products.Product_Price,
		)
		if err != nil {
			return orderOut, totalAmount, err
		}
		totalAmount += (orderIn.Products.Product_Price * (float64(orderIn.Cart_ID.Product_Count)))
		orderOut = append(orderOut, orderIn)
	}
	if err = rows.Err(); err != nil {
		return orderOut, totalAmount, err
	}

	return orderOut, totalAmount, err
}

func (r Repository) OrderPayments(pay models.Payment) error {
	query := `INSERT INTO 
				user_payment(
					user_id,
					cod_payment)
				Values ($1,$2);`

	_, err := r.DB.Exec(
		query,
		pay.User_id,
		pay.COD_Payments,
	)
	if err != nil {
		log.Println("update paymnets failed", err)
		return err
	}
	return nil
}

// func (r Repository) OrderedProduct(usrId models.OrderBody, orderIn models.Cart) error {
// 	var orders models.OrderConfirm
// 	//begin transactions
// 	ctx := context.Background()
// 	tx, err := r.DB.BeginTx(ctx, nil)
// 	if err != nil {
// 		return err
// 	}

// 	for i := 0; i <= len(); i++ {

// 		//execute queries against trasactions
// 		query1 := `SELECT
// 		users.user_id,
// 		users.user_address_id,
// 		product.product_id,
// 		product.product_inventory_id,
// 		cart.product_count,
// 		cart_id
// 		FROM cart
// 		INNER JOIN product ON cart.product_id = product.product_id
// 		INNER JOIN users ON cart.user_id = users.user_id
// 		WHERE cart.user_id =$1;`

// 		err = tx.QueryRow(
// 			query1,
// 			usrId.User_ID).Scan(
// 			&orders.User_id,
// 			&orders.Address_id,
// 			&orders.Product_id[i],
// 			&orders.Inventory_id[i],
// 			&orders.Quantity[i],
// 			&orders.Cart_id,
// 		)
// 		//log.Println("User ID :", orders.User_id)
// 		if err != nil {
// 			log.Println("Select query have a problem", err)
// 			tx.Rollback()
// 			return err
// 		}

// 	}

// 	//looping with productId
// 	for idx := range orders.Product_id {
// 		//need to reduce product quantity in inventory
// 		_, err := tx.Exec(`UPDATE inventory
// 					SET inventory_quantity = invetory_quantity = $1
// 					WHERE inventory_id = $2;`,
// 			orders.Quantity[idx],
// 			orders.Inventory_id[idx],
// 		)
// 		if err != nil {
// 			log.Println("Update inventory failed")
// 			tx.Rollback()
// 			return err
// 		}

// 		//inserting the products form cart into order confirm

// 		_, err = tx.Exec(`
// 		INSERT INTO user_order(
// 			user_id,
// 			address_id,
// 			product_id,
// 			inventory_id,
// 			quantity,
// 			cart_id,
// 			order_at)
// 			VALUES($1,$2,$3,$4,$5,$6,$7);`,
// 			orders.User_id,
// 			orders.Address_id,
// 			orders.Product_id,
// 			orders.Inventory_id,
// 			orders.Quantity,
// 			orders.Cart_id,
// 			time.Now(),
// 		)
// 		if err != nil {
// 			log.Println("Insertion into the orderConfirm failed")
// 			tx.Rollback()
// 			return err
// 		}

// 		//deleting the confirmed order from the cart
// 		_, err = tx.Exec(`TRUNCATE TABLE cart;`)
// 		if err != nil {
// 			log.Println("deleting the cart failed")
// 			tx.Rollback()
// 			return err
// 		}
// 	}
// 	err = tx.Commit()
// 	if err != nil {
// 		log.Println("commit failed")
// 		return err
// 	}
// 	log.Println("Commit sucessful")
// 	return nil

// }

func (r Repository) OrderedProduct(usrId models.OrderBody) error {
	//var orders models.OrderSelect

	//begin transactions
	ctx := context.Background()
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	//execute queries against trasactions
	query1 := `INSERT INTO 
				user_order (
					user_id, 
					address_id, 
					product_id, 
					inventory_id, 
					quantity, 
					cart_id, 
					total_price)
				SELECT
						users.user_id,
						users.user_address_id,
						product.product_id,
						product.product_inventory_id,
						cart.product_count,
						cart_id,
						product.product_price
						FROM cart
						INNER JOIN product ON cart.product_id = product.product_id
						INNER JOIN users ON cart.user_id = users.user_id
						WHERE cart.user_id =$1;`

	_, err = tx.Exec(query1, usrId.User_ID)

	//log.Println("Product ID :", orders.Product_id)

	if err != nil {
		log.Println("Select query have a problem", err)
		tx.Rollback()
	}

	//deleting the confirmed order from the cart
	// _, err = tx.Exec(`DELETE FROM
	// 					cart
	// 					WHERE
	// 					user_id = $1;`,
	// 	usrId.User_ID)

	// if err != nil {
	// 	log.Println("deleting the cart failed", err)
	// 	tx.Rollback()
	// }

	err = tx.Commit()
	if err != nil {
		log.Println("commit failed")
		return err
	}
	log.Println("Commit sucessful")
	return nil

}
