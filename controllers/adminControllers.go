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
			log.Println("Entered wrong admin details")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(utils.PrepareResponse(false, "Incorrect admin details", nil))
			return
		}

		//getting hashed password from database
		dbPassword := admin.Password

		//verifying password
		passwordMatch := utils.VerifyPassword(requestPassword, dbPassword)

		if !passwordMatch {
			log.Println("Invalid Admin Password.")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(utils.PrepareResponse(false, "Invalid Admin Password", nil))
			return
		}

		token, refresh_token := token.GenerateToken(admin.First_Name, admin.Last_Name, admin.Email, admin.Phone_Number, admin.User_ID)

		jwt.Token = token
		jwt.Refresh_Token = refresh_token

		w.WriteHeader(http.StatusOK)
		log.Println("Successfull Admin Login")
		json.NewEncoder(w).Encode(utils.PrepareResponse(true, "Admin Login Sucessful", jwt))

	}
}

func (c Controller) AdminProductView() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var products []models.Product

		products, err := c.ProductRepo.ViewProduct()

		if err != nil {
			log.Println("Error in Executing Query for Product View:", err)
			w.WriteHeader(http.StatusNotImplemented)
			//	utils.ResponseJSON(w, "Error in Executing Query for Product View.")
			json.NewEncoder(w).Encode(utils.PrepareResponse(false, "Error in Executing Query for Product View:", err))
			return
		}

		//utils.ResponseJSON(w, products)
		log.Println("Products are visible to the admin")
		json.NewEncoder(w).Encode(utils.PrepareResponse(true, "Admin can view the products", products))
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
			//	utils.ResponseJSON(w, "Failed to add product.")
			json.NewEncoder(w).Encode(utils.PrepareResponse(false, "Failed to add products", nil))
			return
		}
		log.Println("Product added by admin")
		// utils.ResponseJSON(w, "Product added.")
		// utils.ResponseJSON(w, &pro)
		json.NewEncoder(w).Encode(utils.PrepareResponse(true, "Product added by admin", &pro))

	}
}

//AdmiBlockUser toggles the active status of the User.
func (c Controller) AdminBlockUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var blockuser models.User

		json.NewDecoder(r.Body).Decode(&blockuser)

		blockedUser, err := c.UserRepo.BlockUser(blockuser)

		if err != nil {
			w.WriteHeader(http.StatusNotModified)
			//utils.ResponseJSON(w, "Failed to block user")
			json.NewEncoder(w).Encode(utils.PrepareResponse(false, "unable to update user", nil))
		}
		log.Println("Updated the User Active Status")
		// utils.ResponseJSON(w, "Updates the User status")
		// utils.ResponseJSON(w, &blockedUser)
		json.NewEncoder(w).Encode(utils.PrepareResponse(true, "updated the user status by admin", &blockedUser))

	}
}

func (c Controller) AdminViewUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		viewUser, err := c.UserRepo.ViewUser()

		if err != nil {
			w.WriteHeader((http.StatusNotFound))
			//utils.ResponseJSON(w, "No User Found")
			json.NewEncoder(w).Encode(utils.PrepareResponse(false, "No user found", nil))
		}

		//utils.ResponseJSON(w, &viewUser)
		json.NewEncoder(w).Encode(utils.PrepareResponse(true, "User found", &viewUser))
	}
}

func (c Controller) AdminLogout(w http.ResponseWriter, r *http.Request) {
	//logout function for the admin goes here.
}
