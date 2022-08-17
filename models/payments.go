package models

type Payment struct {
	Payment_ID   int  `json:"paymentID,omitempty"`
	User_id      int  `json:"user_id,omitempty"`
	COD_Payments bool `json:"cod,omitempty"`
	// Payment_Made_At   time.Time `json:"paymentMadeAt"`
	// Payment_Update_At time.Time `json:"paymentsUpdateAt"`
}

type RazorOrder struct {
}

type PayOrders struct {
	UserName     string
	Email        string
	PhoneNumber  int
	ProductCount int
	ProductPrice float64
	TotalAmount  float64
}

type PageVariable struct {
	UserID      int
	OrderID     string
	UserName    string
	Email       string
	PhoneNumber int
	TotalAmount float64
}

type RzrPaySucess struct {
	OrderID   string
	PaymentID string
}
