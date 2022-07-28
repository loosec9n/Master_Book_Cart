package models

type Payment struct {
	Payment_ID int `json:"paymentID"`
	//user_id
	COD_Payments   bool `json:"cod"`
	Payment_Status bool `json:"payment_status"`
	// Payment_Made_At   time.Time `json:"paymentMadeAt"`
	// Payment_Update_At time.Time `json:"paymentsUpdateAt"`
}
