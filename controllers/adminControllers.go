package controllers

import (
	"Book_Cart_Project/models"
	"Book_Cart_Project/repository"
	"Book_Cart_Project/token"
	"Book_Cart_Project/utils"
	"encoding/json"
	"log"
	"net/http"
)

type Controller struct {
	ProductRepo repository.ProductRepository
	UserRepo    repository.UserRepository
}

func (c Controller) AdminLoginIndex(w http.ResponseWriter, r *http.Request) {

}

func (c Controller) AdminLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var admin models.User
		var jwt models.JWT

		json.NewDecoder(r.Body).Decode(&admin)

		// checking whether the Login Credentials is of admin

		requestPassword := admin.Password

		log.Println("Checking whether Admin exists.")
		admin, _ = c.UserRepo.AdminLogin(admin)

		if !admin.IsAdmin {
			log.Println("Please enter correct admin details")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Entered admin id/email not valid.")
			return
		}

		//getting hashed password from database
		dbPassword := admin.Password

		//verifying password
		passwordMatch := utils.VerifyPassword(requestPassword, dbPassword)

		if !passwordMatch {
			log.Println("Invalid Password.")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Passowrd Invalid")
			return
		}

		token, refresh_token := token.GenerateToken(admin.First_Name, admin.Last_Name, admin.Email, admin.Phone_Number, admin.User_ID)

		jwt.Token = token
		jwt.Refresh_Token = refresh_token

		w.WriteHeader(http.StatusOK)
		utils.ResponseJSON(w, jwt)

	}
}

func (c Controller) AdminProductView() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var products []models.Product

		products, err := c.ProductRepo.ViewProduct()

		if err != nil {
			log.Println("Error in Executing Query for Product View:", err)
			w.WriteHeader(http.StatusNotImplemented)
			utils.ResponseJSON(w, "Error in Executing Query for Product View.")
			return
		}

		utils.ResponseJSON(w, products)
	}
}

func (c Controller) AdminProductAdd() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//declaring variable product
		var product models.Product
		//fetching data from fron end into product
		json.NewDecoder(r.Body).Decode(&product)

		pro, err := c.ProductRepo.Addproduct(product)

		if err != nil {
			w.WriteHeader(http.StatusNotImplemented)
			utils.ResponseJSON(w, "Failed to add product.")
			return
		}
		log.Println("Product added.")
		utils.ResponseJSON(w, "Product added.")
		utils.ResponseJSON(w, &pro)

	}
}

func (c Controller) AdminLogout(w http.ResponseWriter, r *http.Request) {

}
