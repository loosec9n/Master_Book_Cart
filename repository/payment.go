package repository

import (
	"Book_Cart_Project/models"
	"context"
	"log"
)

func (r Repository) PaymentMod(usrId int) (models.PageVariable, error) {

	// var payOrder []models.PayOrders
	// var totalAmount float64

	var payOrder models.PageVariable

	ctx := context.Background()
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return payOrder, err
	}

	query1 := `SELECT
				first_name, 
				email, 
				phone_number 
				FROM 
				users 
				WHERE 
				user_id = $1;`

	err = tx.QueryRow(query1, usrId).Scan(
		&payOrder.UserName,
		&payOrder.Email,
		&payOrder.PhoneNumber,
	)
	if err != nil {
		log.Println("Payment mod query 1 failed", err)
		tx.Rollback()
		return payOrder, err
	}

	query2 := `SELECT 
				total_amount 
				FROM ordered 
				WHERE 
				user_id = $1;`

	err = tx.QueryRow(query2, usrId).Scan(
		&payOrder.TotalAmount,
	)
	if err != nil {
		log.Println("Payment mod query 2 failed", err)
		tx.Rollback()
		return payOrder, err
	}

	err = tx.Commit()
	if err != nil {
		log.Println("commit failed")
		return payOrder, err
	}
	return payOrder, err
}

func (r Repository) SucessPayment(payment models.RzrPaySucess, usrID int) error {

	query := `INSERT INTO
				user_payment(
					razorpay_payment_id,
					razorpay_order_id,
					user_id)
				VALUES($1,$2,$3)`

	_, err := r.DB.Exec(
		query,
		payment.PaymentID,
		payment.OrderID,
		usrID)
	if err != nil {
		log.Println("insert into the order failed from razorpay", err)
		return err
	}

	return nil

}
