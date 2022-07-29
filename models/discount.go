package models

type Discount struct {
	Discount_id    int    `json:"discountID"`
	Discount_code  string `json:"discountCode"`
	Discount_price int    `json:"discountPercentage"`
}
