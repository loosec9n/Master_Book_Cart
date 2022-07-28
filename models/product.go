package models

//Product -> product table schema
type Product struct {
	Product_ID          uint            `json:"productId"`
	Is_Active           bool            `json:"isActive"`
	Product_Name        string          `json:"productName"`
	Product_Description string          `json:"produtDescription"`
	Product_Author      ProductAuthor   `json:"productAuthor"`
	Product_Category    ProductCategory `json:"productCategory"`
	Product_Price       float64         `json:"productPrice"`
	Product_Rating      float32         `json:"productRating"`
	Product_Discount_ID Discount        `json:"discountID"`
	Product_Inventory   Inventory       `json:"inventoryID"`
	Product_Image       string          `json:"productImage,omitempty"`
	// Product_Created_At  time.Time       `json:"productCreatedAt,omitempty"`
	// Product_Updated_At  time.Time       `json:"productUpdatedAt,omitempty"`
	// Product_Deleted_At  time.Time       `json:"productDeletedAt,omitempty"`
}

// type ProductUser struct {
// 	ProductUser_ID          uint          `json:"productId"`
// 	ProductUser_Name        string        `json:"productName"`
// 	ProductUser_Count       int           `json:"productUsrCount"`
// 	ProductUser_Description string        `json:"produtDescription"`
// 	ProductUser_Author      ProductAuthor `json:"productAuthor"`
// 	ProductUser_Price       float64       `json:"productPrice"`
// 	ProductUser_Rating      float32       `json:"productRating"`
// 	ProductUser_Image       string        `json:"productImage"`
// 	// ProductUser_Created_At  time.Time     `json:"productUsrCreatedAt"`
// 	// ProductUser_Updated_At  time.Time     `json:"productUsrUpdatedAt"`
// 	// ProductUser_Deleted_At  time.Time     `json:"productUsrDeletedAt"`
// }

type ProductCategory struct {
	Category_ID          uint   `json:"categoryID"`
	Category_Name        string `json:"categoryName"`
	Category_Description string `json:"categoryDescription"`
	// Category_Created_At  time.Time `json:"categoryCreatedAt,omitempty"`
	// Category_Updated_At  time.Time `json:"categoryUpdatedAt,omitempty"`
	// Category_Deleted_At  time.Time `json:"categoryDeletedAt,omitempty"`
}
type ProductAuthor struct {
	Author_ID   uint   `json:"authorID"`
	Author_Name string `json:"authorName"`
	// Author_Created_At time.Time `json:"authorCreatedAt,omitempty"`
	// Author_Updated_At time.Time `json:"authorUpdatedAt,omitempty"`
	// Author_Deleted_At time.Time `json:"authorDeletedAt,omitempty"`
}
type Discount struct {
	Discount_ID          uint    `json:"discountId"`
	Discount_Name        string  `json:"discountName"`
	Discount_Description string  `json:"discountDescription"`
	Discount_Percentage  float32 `json:"discountPercentage"`
	Discount_Status      bool    `json:"discountStatus"`
	// Discount_Created_At  time.Time `json:"discountCreatedAt,omitempty"`
	// Discount_Updated_At  time.Time `json:"discountUpdatedAt,omitempty"`
	// Discount_Deleted_At  time.Time `json:"discountDeletedAt,omitempty"`
}
type Inventory struct {
	Inventory_ID       uint `json:"inventoryID"`
	Inventory_Quantity int  `json:"inventoryQuantity"`
	// Inventory_Created_At time.Time `json:"inventoryCreatedAt,omitempty"`
	// Inventory_Updated_At time.Time `json:"inventoryUpdatedAt,omitempty"`
	// Inventory_Deleted_At time.Time `json:"inventoryDeletedAt,omitempty"`
}
