package models

import "time"

type Order struct {
	Order_ID   uint    `json:"orderId,omitempty"`
	User_ID    User    `json:"userID,omitempty"`
	Address_ID Address `json:"addressID,omitempty"`
	Products   Product `json:"products,omitempty"`
	Cart_ID    Cart    `json:"cartID,omitempty"`
	TotalPrice float64 `json:"toatlPrice"`
	//Inventory_ID Inventory      `json:"inentoryID,omitempty"`
	// Order_status bool      `json:"orderStatus,omitempty"`
	// Created_at   time.Time `json:"orderAt"`
	// Canceled_at  time.Time `json:"canceledAt,omitempty"`
	// Deleted_at   time.Time `json:"deletedAt,omitempty"`
}

//order body from json
type OrderBody struct {
	User_ID uint `json:"id_buyer"`
}

type Wishlist struct {
	Wishlist_ID   int     `json:"whishlistId,omitempty"`
	User_ID       User    `json:"userID,omitempty"`
	Product_ID    Product `json:"productID,omitempty"`
	Product_Image string  `json:"productImage,omitempty"`
	Product_Price Product `josn:"productPrice,omitempty"`
	// Created_at  time.Time `json:"createdAt"`
	// Updated_at  time.Time `json:"updatedAt,omitempty"`
	// Deleted_at  time.Time `json:"deletedAt,omitempty"`
}

type OrderConfirm struct {
	Order_id     uint
	User_id      uint    `json:"user_id"`
	Address_id   uint    `json:"address_id"`
	Product_id   int     `josn:"product_id"`
	Inventory_id int     `json:"inventory_id"`
	Quantity     int     `json:"quantity"`
	Cart_id      int     `json:"cart_id"`
	Pro_price    float64 `json:"product_price"`
	Total_price  float64 `json:"total_price"`
	//Payment_id   string  `json:"payment_id"`
}
type OrderSelect struct {
	User_id      uint    `json:"user_id"`
	Address_id   uint    `json:"address_id"`
	Product_id   int     `josn:"product_id"`
	Inventory_id int     `json:"inventory_id"`
	Quantity     int     `json:"quantity"`
	Cart_id      int     `json:"cart_id"`
	TotalPrice   float64 `json:"total_price"`
	//Payment_id   string  `json:"payment_id"`
}

type OrderReport struct {
	Order_month time.Time `json:"order_month"`
	Total_price float64   `json:"total_price"`
	Order_count int       `json:"order_count"`
}

type ReportIn struct {
	Order_status string `json:"order_status"`
}
type ChangeOrder struct {
	Order_status string `json:"order_status"`
	Order_id     int    `json:"order_id"`
}
type Ordered struct {
	Order_id     int
	User_id      int
	Total_amount float64
	Payed        bool
}
