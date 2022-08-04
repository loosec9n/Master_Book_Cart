package models

type Discount struct {
	Discount_id    int    `json:"discountID,omitempty"`
	Discount_code  string `json:"discountCode,omitempty"`
	Discount_price int    `json:"discountPercentage,omitempty"`
}
