package models

type Payment struct {
	Payment_ID   int  `json:"paymentID,omitempty"`
	User_id      int  `json:"user_id,omitempty"`
	COD_Payments bool `json:"cod,omitempty"`
	// Payment_Made_At   time.Time `json:"paymentMadeAt"`
	// Payment_Update_At time.Time `json:"paymentsUpdateAt"`
}
