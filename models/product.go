package models

import "time"

//Product -> product table schema
type Product struct {
	Product_ID           uint            `json:"productId"`
	Product_Name         string          `json:"productName"`
	Product_Description  string          `json:"produtDescription"`
	Product_Category_ID  ProductCategory `json:"categoryID"`
	Product_Author_ID    ProductAuthor   `json:"productAuthor"`
	Product_Price        float64         `json:"productPrice"`
	Product_Rating       float32         `json:"productRating"`
	Product_Discount_ID  Discount        `json:"discountID"`
	Product_Inventory_ID Inventory       `json:"inventoryID"`
	Product_Image        string          `json:"productImage"`
	Product_Created_At   time.Time       `json:"productCreatedAt"`
	Product_Updated_At   time.Time       `json:"productUpdatedAt"`
	Product_Deleted_At   time.Time       `json:"productDeletedAt"`
}
type ProductUser struct {
	ProductUser_ID          uint          `json:"productId"`
	ProductUser_Name        string        `json:"productName"`
	ProductUser_Count       int           `json:"productUsrCount"`
	ProductUser_Description string        `json:"produtDescription"`
	ProductUser_Author      ProductAuthor `json:"productAuthor"`
	ProductUser_Price       float64       `json:"productPrice"`
	ProductUser_Rating      float32       `json:"productRating"`
	ProductUser_Image       string        `json:"productImage"`
	ProductUser_Created_At  time.Time     `json:"productUsrCreatedAt"`
	ProductUser_Updated_At  time.Time     `json:"productUsrUpdatedAt"`
	ProductUser_Deleted_At  time.Time     `json:"productUsrDeletedAt"`
}

type ProductCategory struct {
	Category_ID          uint      `json:"categoryId"`
	Category_Name        string    `json:"categoryName"`
	Category_Description string    `json:"categoryDescription"`
	Category_Created_At  time.Time `json:"categoryCreatedAt"`
	Category_Updated_At  time.Time `json:"categoryUpdatedAt"`
	Category_Deleted_At  time.Time `json:"categoryDeletedAt"`
}
type ProductAuthor struct {
	Author_ID         uint      `json:"authorID"`
	Author_Name       string    `json:"authorName"`
	Author_Created_At time.Time `json:"authorCreatedAt"`
	Author_Updated_At time.Time `json:"authorUpdatedAt"`
	Author_Deleted_At time.Time `json:"authorDeletedAt"`
}
type Discount struct {
	Discount_ID          uint      `json:"discountId"`
	Discount_Name        string    `json:"discountName"`
	Discount_Description string    `json:"discountDescription"`
	Discount_Percentage  float32   `json:"discountPercentage"`
	Discount_Status      bool      `json:"discountStatus"`
	Discount_Created_At  time.Time `json:"discountCreatedAt"`
	Discount_Updated_At  time.Time `json:"discountUpdatedAt"`
	Discount_Deleted_At  time.Time `json:"discountDeletedAt"`
}
type Inventory struct {
	Inventory_ID         uint      `json:"inventoryID"`
	Inventory_Quantity   int       `json:"inventoryQuantity"`
	Inventory_Created_At time.Time `json:"innentoryCreatedAt"`
	Inventory_Updated_At time.Time `json:"inventoryUpdatedAt"`
	Inventory_Deleted_At time.Time `json:"inventoryDeletedAt"`
}
