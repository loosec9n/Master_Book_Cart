package repository

import (
	"Book_Cart_Project/models"
	"fmt"
	"log"
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

	log.Println("product from addproduct", product)

	query := `INSERT INTO product (
		product_name, 
		product_description, 
		product_price,
		product_category_id,
		product_author_id)
		VALUES($1, $2, $3, $4, $5) 
		RETURNING 
		product_id, 
		product_name, 
		product_description, 
		product_price,
		product_category_id,
		product_author_id;`

	err := r.DB.QueryRow(query,
		product.Product_Name,
		product.Product_Description,
		product.Product_Price,
		product.Product_Category.Category_ID,
		product.Product_Author.Author_ID,
	).Scan(
		&product.Product_ID,
		&product.Product_Name,
		&product.Product_Description,
		&product.Product_Price,
		&product.Product_Category.Category_ID,
		&product.Product_Author.Author_ID,
	)

	log.Println("add product error:", err)
	return product, err

}

func (r Repository) ViewProduct(filter models.Filter) ([]Prod, models.Metadata, error) {
	var products []Prod
	// value := 1
	// var arg []interface{}

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

	// if searchParam.Product != "" {
	// 	query = query + `WHERE product_name iLIKE $` + fmt.Sprintf(`%d`, value) + `,`
	// 	arg = append(arg, searchParam.Product)
	// 	value++
	// }

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

func (r Repository) UserSearchProduct(searchParam models.SearchParm) ([]Prod, error) {
	var usps []Prod
	var arg []interface{}
	var flag bool
	var Oflag bool
	i := 1

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
	`
	//product_id = $1;`
	// if searchParam.Product != " " {
	// products := fmt.Sprintf("%s WHERE product_name iLike '%%%d%%'", query, searchParam.P)
	// }

	// if (searchParam.Product || searchParam.Categorty || searchParam.Author) != ""{
	// 	query = query +`WHERE`
	// }

	if searchParam.Product != "" {
		if !flag {
			query = query + `WHERE `
			flag = true
		}
		query = query + `product_name iLIKE $` + fmt.Sprintf("%d", i) //+ " AND "
		arg = append(arg, fmt.Sprint("%", searchParam.Product, "%"))
		i++
	}

	if searchParam.Categorty != "" {
		if !flag {
			query = query + `WHERE `
			flag = true
		} else {
			query = query + ` AND `
		}
		query = query + `category_name iLike $` + fmt.Sprintf("%d", i)
		arg = append(arg, fmt.Sprint("%", searchParam.Categorty, "%"))
		i++
	}

	if searchParam.Author != "" {
		if !flag {
			query = query + `WHERE `
			flag = true
		} else {
			query = query + ` AND `
		}
		query = query + `author_name iLike $` + fmt.Sprintf("%d", i)
		arg = append(arg, fmt.Sprint("%", searchParam.Author, "%"))
		i++
	}

	//ordering the output accouding to user preference
	if searchParam.OrderBY != "" {
		if !Oflag {
			query = query + `ORDER BY `
			Oflag = true
		} else {
			query = query + ` , `
		}

		if searchParam.OrderBY == "asc" {
			query = query + `product_name`
		} else {
			query = query + `product_name DESC`
		}
	}

	if searchParam.Oprice != "" {
		if !Oflag {
			query = query + `ORDER BY `
			Oflag = true
		} else {
			query = query + ` , `
		}
		if searchParam.Oprice == "asc" {
			query = query + `product_price`
		} else {
			query = query + `product_price DESC`
		}
	}

	log.Println("query", query)
	log.Println("arg", arg)

	stmt, err := r.DB.Prepare(query)
	if err != nil {
		log.Println("Preparing the query failed", err)
		return usps, err
	}

	row, err := stmt.Query(arg...)
	if err != nil {
		log.Println("Search product query failed", err)
		return usps, err
	}
	defer row.Close()

	for row.Next() {
		var usp Prod
		err := row.Scan(
			&usp.Is_Active,
			&usp.Id,
			&usp.Name,
			&usp.Description,
			&usp.Category,
			&usp.Author,
			&usp.Price,
		)
		if err != nil {
			log.Println("error scaning the search product")
			return usps, err
		}
		if usp.Is_Active {
			usps = append(usps, usp)
		}
	}
	if err = row.Err(); err != nil {
		return usps, err
	}

	return usps, err
}

func (r Repository) CheckActiveProd(product_id int) (bool, error) {

	var activeProduct Prod

	query := `SELECT is_active
		FROM product
		WHERE product_id = $1;`

	err := r.DB.QueryRow(query, product_id).Scan(&activeProduct.Is_Active)

	return activeProduct.Is_Active, err
}

func (r Repository) AddInventory(inventory models.Inventory) (models.Inventory, error) {
	query := `INSERT INTO inventory(
		inventory_id,
		inventory_quantity)
		VALUES($1,$2)
		RETURNING
		inventory_id,
		inventory_quantity;`

	err := r.DB.QueryRow(query,
		inventory.Inventory_ID,
		inventory.Inventory_Quantity).Scan(
		&inventory.Inventory_ID,
		&inventory.Inventory_Quantity,
	)
	return inventory, err
}
