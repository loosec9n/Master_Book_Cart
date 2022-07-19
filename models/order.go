package models

import "time"

type Order struct {
	Order_ID uint `json:"orderId"`
	//Order_Cart     []ProductUser `json:"orderCart"`
	Order_At       time.Time `json:"orderAt"`
	Order_Price    uint      `json:"price"`
	Order_Discount int       `json:"discount"`
	Payment_Method Payment   `json:"paymentMethod"`
}
