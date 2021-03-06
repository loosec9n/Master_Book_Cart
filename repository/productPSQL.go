package repository

import (
	"Book_Cart_Project/models"
	"time"
)

type Prod struct {
	Id          int       `json:"productId"`
	Is_Active   bool      `json:"isActive"`
	Name        string    `json:"productName"`
	Description string    `json:"productDescription"`
	Author      string    `json:"productAuthor"`
	Category    string    `josn:"productCategory"`
	Price       float64   `json:"productPrice"`
	Created_At  time.Time `json:"created_at"`
}

func (r Repository) Addproduct(product models.Product) (models.Product, error) {

	query := `INSERT INTO product (
		product_name, 
		product_description, 
		product_price,
		product_category_id,
		product_author_id,
		prduct_inentory_id)
		VALUES($1, $2, $3, $4, $5) 
		RETURNING 
		product_id, 
		product_name, 
		product_description, 
		product_price,
		product_author_id,
		product_category_id;`

	err := r.DB.QueryRow(query,
		product.Product_Name,
		product.Product_Description,
		product.Product_Price,
		product.Product_Category.Category_ID,
		product.Product_Author.Author_ID,
		product.Product_Inventory.Inventory_ID,
	).Scan(
		&product.Product_ID,
		&product.Product_Name,
		&product.Product_Description,
		&product.Product_Price,
		&product.Product_Author.Author_ID,
		&product.Product_Category.Category_ID)

	return product, err

}

func (r Repository) ViewProduct(filter models.Filter) ([]Prod, models.Metadata, error) {
	var products []Prod

	//Writing and executing query
	query := `SELECT COUNT(*) OVER(),
		product.product_id, 
		product.product_name, 
		product.product_description, 
		product_author.author_name,
		product_category.category_name,
		product.product_price
		FROM 
		product
		INNER JOIN 
		product_category 
		ON 
		product_category.category_id = product.product_category_id
		INNER JOIN
		product_author
		ON 
		product_author.author_id = product.product_author_id
		LIMIT $1 OFFSET $2;`

	rows, err := r.DB.Query(query, filter.Limit(), filter.Offset())

	if err != nil {
		return nil, models.Metadata{}, err
	}
	defer rows.Close()

	// if searchParam := r.DB.Query("search"); searchParam != " " {
	// 	products = fmt.Sprintf("%s WHERE product_name iLike '%%%search%%'", query, searchParam)
	// }

	var toatalRecords int

	// Loop through rows, using scan to assign column data to struct fields
	for rows.Next() {
		var product Prod
		if err := rows.Scan(&toatalRecords,
			&product.Id,
			&product.Name,
			&product.Description,
			&product.Category,
			&product.Author,
			&product.Price); err != nil {
			return products, models.Metadata{}, err
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return products, models.Metadata{}, err
	}

	return products, models.ComputeMetadata(toatalRecords, filter.Page, filter.PageSize), nil

}

func (r Repository) BlockProduct(product models.Product) (models.Product, error) {
	query := `UPDATE product
		SET is_active = $1
		WHERE product_id = $2
		RETURNING
		product_id,
		is_active,
		product_name,
		product_description,
		product_price;`
	err := r.DB.QueryRow(query,
		product.Is_Active,
		product.Product_ID,
	).Scan(
		&product.Product_ID,
		&product.Is_Active,
		&product.Product_Name,
		&product.Product_Description,
		&product.Product_Price,
	)
	return product, err

}

func (r Repository) UserSearchProduct(product_id int) (Prod, error) {
	var usp Prod

	query := `SELECT 
	product.is_active,
	product.product_id,
	product.product_name,
	product.product_description,
	product_category.category_name,
	product_author.author_name,
	product.product_price
	FROM product
	INNER JOIN 
	product_category 
	ON 
	product_category.category_id = product.product_category_id
	INNER JOIN 
	product_author
	ON 
	product_author.author_id = product.product_author_id
	WHERE 
	product_id = $1;`

	err := r.DB.QueryRow(query,
		product_id).Scan(
		&usp.Is_Active,
		&usp.Id,
		&usp.Name,
		&usp.Description,
		&usp.Category,
		&usp.Author,
		&usp.Price,
	)

	return usp, err
}

func (r Repository) CheckActiveProd(product_id int) (bool, error) {

	var activeProduct Prod

	query := `SELECT is_active
		FROM product
		WHERE product_id = $1;`

	err := r.DB.QueryRow(query, product_id).Scan(&activeProduct.Is_Active)

	return activeProduct.Is_Active, err
}
