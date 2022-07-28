package models

type Order struct {
	Order_ID     uint      `json:"orderId"`
	User_ID      User      `json:"userID"`
	Address_ID   Address   `json:"addressID"`
	Product_ID   Product   `json:"productID"`
	Payment_ID   Payment   `json:"paymentID"`
	Inventory_ID Inventory `json:"inentoryID"`
	Order_status bool      `json:"orderStatus"`
	Cart_ID      Cart      `json:"cartID"`
	// Created_at   time.Time `json:"orderAt"`
	// Canceled_at  time.Time `json:"canceledAt,omitempty"`
	// Deleted_at   time.Time `json:"deletedAt,omitempty"`
}

type Wishlist struct {
	Wishlist_ID   int     `js0n:"whishlistId"`
	User_ID       User    `json:"userID"`
	Product_ID    Product `json:"productID"`
	Product_Image string  `json:"productImage"`
	Product_Price Product `josn:"productPrice"`
	// Created_at  time.Time `json:"createdAt"`
	// Updated_at  time.Time `json:"updatedAt,omitempty"`
	// Deleted_at  time.Time `json:"deletedAt,omitempty"`
}
