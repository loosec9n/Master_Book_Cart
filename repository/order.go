package repository

import (
	"Book_Cart_Project/models"
	"context"
	"log"
	"time"
)

func (r Repository) CreateNewOrder(order models.OrderBody, orderIn models.Order) ([]models.Order, float64, error) {
	var orderOut []models.Order
	//var ordered models.Ordered

	ctx := context.Background()
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return orderOut, 0, err
	}

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

	rows, err := tx.Query(query, order.User_ID)
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
		log.Println("Create new order select query failed", err)
		tx.Rollback()
		return orderOut, totalAmount, err
	}

	query2 := `INSERT INTO
				ordered(user_id,total_amount)
				VALUES($1,$2);`
	_, err = tx.Exec(query2,
		order.User_ID,
		totalAmount,
	)
	if err != nil {
		log.Println("Create new order insert query failed", err)
		tx.Rollback()
		return orderOut, totalAmount, err
	}
	err = tx.Commit()
	if err != nil {
		log.Println("commit failed")
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

func (r Repository) OrderedProduct(usrId models.OrderBody, orderIn models.Cart) error {
	var orders []models.OrderConfirm
	//var total_price float64
	//begin transactions
	ctx := context.Background()
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// for i := 0; i <= len(); i++ {

	//execute queries against trasactions
	query1 := `SELECT
		users.user_id,
		users.user_address_id,
		product.product_id,
		product.product_price,
		product.product_inventory_id,
		cart.product_count,
		cart_id
		FROM cart
		INNER JOIN product ON cart.product_id = product.product_id
		INNER JOIN users ON cart.user_id = users.user_id
		WHERE cart.user_id =$1;`

	rows, err := tx.Query(
		query1,
		usrId.User_ID)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var order models.OrderConfirm
		err := rows.Scan(
			&order.User_id,
			&order.Address_id,
			&order.Product_id,
			&order.Pro_price,
			&order.Inventory_id,
			&order.Quantity,
			&order.Cart_id,
		)
		if err != nil {
			log.Println("error in rows.next", err)
			return err
		}
		order.Total_price += (order.Pro_price * float64(order.Quantity))

		orders = append(orders, order)
	}

	if err = rows.Err(); err != nil {
		log.Println("Select query have a problem", err)
		tx.Rollback()
		return err
	}
	log.Println("Select from cart worked")

	//looping with productId
	for _, v := range orders {
		//need to reduce product quantity in inventory
		_, err := tx.Exec(`UPDATE inventory
					SET inventory_quantity = inventory_quantity - $1
					WHERE inventory_id = $2;`,
			v.Product_id,
			v.Inventory_id,
		)
		if err != nil {
			log.Println("Update inventory failed", err)
			tx.Rollback()
			return err
		}
		log.Println("inventory updation worked")

		//inserting the products form cart into order confirm

		_, err = tx.Exec(`
		INSERT INTO user_order(
			user_id,
			address_id,
			product_id,
			inventory_id,
			quantity,
			product_price,
			total_price,
			order_at)
			VALUES($1,$2,$3,$4,$5,$6,$7,$8);`,
			v.User_id,
			v.Address_id,
			v.Product_id,
			v.Inventory_id,
			v.Quantity,
			v.Pro_price,
			v.Total_price,
			time.Now(),
		)
		if err != nil {
			log.Println("Insertion into the orderConfirm failed", err)
			tx.Rollback()
			return err
		}
		log.Println("Insertion into user order worked")
		//deleting the confirmed order from the cart

		_, err = tx.Exec(`DELETE FROM
		 					cart
		 					WHERE
		 					user_id = $1;`,
			usrId.User_ID)

		if err != nil {
			log.Println("deleting the cart failed", err)
			tx.Rollback()
			return err
		}
		log.Println("delete from cart worked")
	}
	err = tx.Commit()
	if err != nil {
		log.Println("commit failed")
		return err
	}
	log.Println("Commit sucessful")
	return nil

}

// func (r Repository) OrderedProduct(usrId models.OrderBody) error {
// 	//var orders models.OrderSelect

// 	//begin transactions
// 	ctx := context.Background()
// 	tx, err := r.DB.BeginTx(ctx, nil)
// 	if err != nil {
// 		return err
// 	}

// 	//execute queries against trasactions
// 	query1 := `INSERT INTO
// 				user_order (
// 					user_id,
// 					address_id,
// 					product_id,
// 					inventory_id,
// 					quantity,
// 					cart_id,
// 					total_price)
// 				SELECT
// 						users.user_id,
// 						users.user_address_id,
// 						product.product_id,
// 						product.product_inventory_id,
// 						cart.product_count,
// 						cart_id,
// 						product.product_price
// 						FROM cart
// 						INNER JOIN product ON cart.product_id = product.product_id
// 						INNER JOIN users ON cart.user_id = users.user_id
// 						WHERE cart.user_id =$1;`

// 	_, err = tx.Exec(query1, usrId.User_ID)

// 	//log.Println("Product ID :", orders.Product_id)

// 	if err != nil {
// 		log.Println("Select query have a problem", err)
// 		tx.Rollback()
// 	}

// 	//deleting the confirmed order from the cart
// 	// _, err = tx.Exec(`DELETE FROM
// 	// 					cart
// 	// 					WHERE
// 	// 					user_id = $1;`,
// 	// 	usrId.User_ID)

// 	// if err != nil {
// 	// 	log.Println("deleting the cart failed", err)
// 	// 	tx.Rollback()
// 	// }

// 	err = tx.Commit()
// 	if err != nil {
// 		log.Println("commit failed")
// 		return err
// 	}
// 	log.Println("Commit sucessful")
// 	return nil

// }
