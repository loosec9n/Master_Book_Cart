package models

import "time"

type Session_Cart struct {
	Session_ID uint `json:"sessionId"`
	User_ID    User `json:"userId"`
	// Session_Created_At time.Time `json:"sessionCreatedAt"`
	// Session_Updated_At time.Time `json:"sessionUpdatedAt"`
}

type Cart struct {
	Cart_ID uint `json:"cartId,omitempty"`
	// Session_ID      Session_Cart `json:"session_cart"`
	User_ID         int       `json:"cartUserId,omitempty"`
	Product_ID      Product   `json:"productId,omitempty"`
	Product_Count   int       `json:"productCount,omitempty"`
	Inventory_ID    Inventory `json:"inventoryID,omitempty"`
	Cart_Created_At time.Time `json:"cartCreatedAt,omitempty"`
	Cart_Updated_At time.Time `json:"cartUpdatedAt,omitempty"`
}
