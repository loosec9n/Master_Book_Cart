package models

import "time"

type Session_Cart struct {
	Session_ID         uint      `json:"sessionId"`
	User_ID            User      `json:"userId"`
	Session_Created_At time.Time `json:"sessionCreatedAt"`
	Session_Updated_At time.Time `json:"sessionUpdatedAt"`
}

type Cart struct {
	Cart_ID uint `json:"cartId"`
	// Session_ID      Session_Cart `json:"session_cart"`
	User_ID         int       `json:"cartUserId"`
	Product_ID      Product   `json:"productId"`
	Product_Count   int       `json:"productCount"`
	Cart_Created_At time.Time `json:"cartCreatedAt"`
	Cart_Updated_At time.Time `json:"cartUpdatedAt"`
}
