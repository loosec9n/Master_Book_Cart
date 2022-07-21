package repository

import (
	"Book_Cart_Project/models"
	"log"
)

func (r Repository) AddCategory(category models.ProductCategory) (models.ProductCategory, error) {

	query := `INSERT INTO product_category(
		category_id,
		category_name,
		category_description)
		VALUES($1,$2,$3)
		RETURNING
		category_id,
		category_name,
		category_description;`

	err := r.DB.QueryRow(query,
		category.Category_ID,
		category.Category_Name,
		category.Category_Description).Scan(
		&category.Category_ID,
		&category.Category_Name,
		&category.Category_Description,
	)

	return category, err
}

func (r Repository) ViewCategory() ([]models.ProductCategory, error) {
	var categories []models.ProductCategory

	//Query for selecting the categories
	query := `SELECT 
		category_id, 
		category_name,
		category_description
		FROM 
		product_category;`

	rows, err := r.DB.Query(query)

	if err != nil {
		log.Println("Category was not selected form Database")
		return nil, err
	}
	defer rows.Close()

	//Looping throu rows -> assigning data to the struct
	for rows.Next() {
		var category models.ProductCategory
		if err := rows.Scan(
			&category.Category_ID,
			&category.Category_Name,
			&category.Category_Description,
		); err != nil {
			return categories, err
		}
		categories = append(categories, category)
	}

	if err = rows.Err(); err != nil {
		return categories, err
	}
	return categories, nil
}

func (r Repository) AddAuthor(author models.ProductAuthor) (models.ProductAuthor, error) {

	//query for inserting the author into the author table
	query := `INSERT INTO product_author(
		author_id,	
		author_name)
			VALUES($1,$2)
			RETURNING
			author_id,
			author_name;`
	err := r.DB.QueryRow(query,
		author.Author_ID,
		author.Author_Name).Scan(
		&author.Author_ID,
		&author.Author_Name,
	)

	return author, err
}

func (r Repository) ViewAuthor() ([]models.ProductAuthor, error) {

	//Query to select the authors form author table
	query := `SELECT 
		author_id,
		author_name,
		author_created_at
		FROM product_author;`

	rows, err := r.DB.Query(query)
	if err != nil {
		log.Println("Was not able to SELECT form Author Table")
		return nil, err
	}
	defer rows.Close()

	var authors []models.ProductAuthor

	//Looping through the author table
	for rows.Next() {
		var author models.ProductAuthor
		err := rows.Scan(
			&author.Author_ID,
			&author.Author_Name,
		)
		if err != nil {
			log.Println("Was not able to Scan in Author table")
			return nil, err
		}
		authors = append(authors, author)
	}

	if err := rows.Err(); err != nil {
		return authors, err
	}
	return authors, nil
}
