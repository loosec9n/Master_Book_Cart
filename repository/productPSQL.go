package repository

import "Book_Cart_Project/models"

func (r Repository) Addproduct(product models.Product) (models.Product, error) {

	query := `INSERT INTO product (
		product_name, product_description, product_price)
		VALUES($1, $2, $3) RETURNING product_id, product_name, product_description, product_price;`
	err := r.DB.QueryRow(query, product.Product_Name, product.Product_Description, product.Product_Price).Scan(
		&product.Product_ID, &product.Product_Name, &product.Product_Description, &product.Product_Price)

	return product, err

}

func (r Repository) ViewProduct() ([]models.Product, error) {
	var products []models.Product

	//Writing and executing query
	query := `SELECT product_id, product_name, product_description, product_price FROM product;`
	rows, err := r.DB.Query(query)

	if err != nil {
		return nil, err
	}

	// Loop through rows, using scan to assign column data to struct fields
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.Product_ID, &product.Product_Name, &product.Product_Description, &product.Product_Price); err != nil {
			return products, err
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return products, err
	}
	return products, nil

}

// func (r Repository) Category(category models.ProductCategory) (models.ProductCategory, error) {

// 	//Query for category
// 	query := `SELECT category_id FROM product_category
// 				WHERE category_id = $1`
// 	err := r.DB.QueryRow(query, category.Category_ID).Scan(&category.Category_ID)

// 	if err != nil {
// 		logFatal(err)
// 	}
// 	return category, err
// }
