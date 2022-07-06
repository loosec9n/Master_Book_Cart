package repository

import (
	"Book_Cart_Project/models"
	"database/sql"
)

type Repository struct {
	DB *sql.DB
}

//Dependency Injection using interface

type UserRepository interface {
	UserSignup(models.User) models.User
	DoesUserExists(models.User) bool
	UserLogin(models.User) (models.User, error)
	AdminLogin(models.User) (models.User, error)
}

type ProductRepository interface {
	Addproduct(models.Product) (models.Product, error)
	ViewProduct() ([]models.Product, error)
}
