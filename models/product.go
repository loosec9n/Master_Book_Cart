package models

//Product -> product table schema
type Product struct {
	Product_ID          int             `json:"productId,omitempty"`
	Is_Active           bool            `json:"isActive,omitempty"`
	Product_Name        string          `json:"productName,omitempty"`
	Product_Description string          `json:"produtDescription,omitempty"`
	Product_Author      ProductAuthor   `json:"productAuthor,omitempty"`
	Product_Category    ProductCategory `json:"productCategory,omitempty"`
	Product_Price       float64         `json:"productPrice,omitempty"`
	Product_Rating      float32         `json:"productRating,omitempty"`
	Product_Discount_ID Discount        `json:"discountID,omitempty"`
	Product_Inventory   Inventory       `json:"inventoryID,omitempty"`
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
	Category_ID          uint   `json:"categoryID,omitempty"`
	Category_Name        string `json:"categoryName,omitempty"`
	Category_Description string `json:"categoryDescription,omitempty"`
	// Category_Created_At  time.Time `json:"categoryCreatedAt,omitempty"`
	// Category_Updated_At  time.Time `json:"categoryUpdatedAt,omitempty"`
	// Category_Deleted_At  time.Time `json:"categoryDeletedAt,omitempty"`
}
type ProductAuthor struct {
	Author_ID   uint   `json:"authorID,omitempty"`
	Author_Name string `json:"authorName,omitempty"`
	// Author_Created_At time.Time `json:"authorCreatedAt,omitempty"`
	// Author_Updated_At time.Time `json:"authorUpdatedAt,omitempty"`
	// Author_Deleted_At time.Time `json:"authorDeletedAt,omitempty"`
}

type Inventory struct {
	Inventory_ID       uint `json:"inventoryID,omitempty"`
	Inventory_Quantity int  `json:"inventoryQuantity,omitempty"`
	// Inventory_Created_At time.Time `json:"inventoryCreatedAt,omitempty"`
	// Inventory_Updated_At time.Time `json:"inventoryUpdatedAt,omitempty"`
	// Inventory_Deleted_At time.Time `json:"inventoryDeletedAt,omitempty"`
}
