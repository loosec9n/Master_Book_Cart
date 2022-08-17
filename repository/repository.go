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
	ViewCart(models.Cart) ([]models.Cart, float64, error)
	CheckActiveProd(int) (bool, error)
	UserSearchProduct(models.SearchParm) ([]Prod, error)
	FindUserByEmail(models.User) (models.User, error)
	ForgetPasswordUpdate(models.User, string) (models.ForgotPasswordInput, error)
	AddAddress(models.Address) (models.Address, error)
	AddWishlist(models.Wishlist) (models.Wishlist, error)
	ViewWishlist(models.Filter, models.Wishlist) ([]wishList, models.Metadata, error)
	DeleteProductWishlist(models.Wishlist) (int64, error)
	CreateNewOrder(models.OrderBody, models.Order) ([]models.Order, float64, error)
	OrderPayments(models.Payment) error
	OrderedProduct(models.OrderBody, models.Cart) error
	PaymentMod(int) (models.PageVariable, error)
	SucessPayment(models.RzrPaySucess, int) error
}

type ProductRepository interface {
	Addproduct(models.Product) (models.Product, error)
	ViewProduct(models.Filter) ([]Prod, models.Metadata, error)
	BlockProduct(models.Product) (models.Product, error)
	AddCategory(models.ProductCategory) (models.ProductCategory, error)
	ViewCategory() ([]models.ProductCategory, error)
	AddAuthor(models.ProductAuthor) (models.ProductAuthor, error)
	ViewAuthor() ([]models.ProductAuthor, error)
	AddInventory(models.Inventory) (models.Inventory, error)
	AdminReport(models.ReportIn) ([]models.OrderReport, error)
	EditOrderStatus(models.ChangeOrder) (models.ChangeOrder, error)
}
