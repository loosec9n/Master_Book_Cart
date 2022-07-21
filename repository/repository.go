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
	BlockUser(models.User) (models.User, error)
	ViewUser() ([]models.User, error)
	AddToCart(models.Cart) (models.Cart, error)
	ViewCart(models.Cart) ([]models.Cart, error)
	CheckActiveProd(int) (bool, error)
	UserSearchProduct(int) (Prod, error)
	FindUserByEmail(models.User) (models.User, error)
	ForgetPasswordUpdate(models.User, string) (models.ForgotPasswordInput, error)
}

type ProductRepository interface {
	Addproduct(models.Product) (models.Product, error)
	ViewProduct(models.Filter) ([]Prod, models.Metadata, error)
	BlockProduct(models.Product) (models.Product, error)
	AddCategory(models.ProductCategory) (models.ProductCategory, error)
	ViewCategory() ([]models.ProductCategory, error)
	AddAuthor(models.ProductAuthor) (models.ProductAuthor, error)
	ViewAuthor() ([]models.ProductAuthor, error)
}
