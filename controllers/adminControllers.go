package controllers

import (
	"Book_Cart_Project/models"
	"Book_Cart_Project/repository"
	"Book_Cart_Project/token"
	"Book_Cart_Project/utils"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
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
		admin, err := c.UserRepo.AdminLogin(admin)

		if err != nil {
			if err == sql.ErrNoRows {
				log.Println("Please enter correct admin details")
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(utils.PrepareResponse(false, "Invalid admin email", err))
				return
			} else {
				log.Println("Please enter correct admin details")
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(utils.PrepareResponse(false, "Please enter correct admin details", err))
				return
			}

		}

		if !admin.IsAdmin {
			log.Println("Not a admin")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(utils.PrepareResponse(false, "Not an admin", nil))
			return
		}

		//getting hashed password from database
		dbPassword := admin.Password
		//fmt.Println("DB password", dbPassword)
		//verifying password
		passwordMatch := utils.VerifyPassword(requestPassword, dbPassword)

		if !passwordMatch {
			log.Println("Invalid Admin Password.")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(utils.PrepareResponse(false, "Invalid Admin Password", nil))
			return
		}

		token, refresh_token := token.GenerateTokenAdmin(admin.First_Name, admin.Last_Name, admin.Email, admin.Phone_Number, admin.User_ID, admin.IsAdmin)

		jwt.Token = token
		jwt.Refresh_Token = refresh_token

		w.WriteHeader(http.StatusOK)
		log.Println("Successfull Admin Login")
		json.NewEncoder(w).Encode(utils.PrepareResponse(true, "Admin Login Sucessful", jwt))

	}
}

func (c Controller) AdminProductView() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil || params < 1 {
			log.Println("Error converting string to int", err)
			w.WriteHeader(http.StatusNotImplemented)
			json.NewEncoder(w).Encode(utils.PrepareResponse(false, "Error converting page to params for page number", err))
			return
		}

		f := models.Filter{
			PageSize: 5,
			Page:     params,
		}

		result := models.Metadata{}

		products, result, err := c.ProductRepo.ViewProduct(f)

		if err != nil {
			log.Println("Error in Executing Query for Product View:", err)
			w.WriteHeader(http.StatusNotImplemented)
			json.NewEncoder(w).Encode(utils.PrepareResponse(false, "Error in Executing Query for Product View:", err))
			return
		}

		log.Println("Products are visible to the admin")
		json.NewEncoder(w).Encode(utils.PrepareResponse(true, "Admin can view the products", map[string]interface{}{
			"data":     &products,
			"total":    result.TotalRecords,
			"page":     result.CurrentPage,
			"lastpage": result.LastPage,
		}))
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
			json.NewEncoder(w).Encode(utils.PrepareResponse(false, "Failed to add products", err))
			return
		}
		log.Println("Product added by admin")

		json.NewEncoder(w).Encode(utils.PrepareResponse(true, "Product added by admin", &pro))

	}
}

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

func (c Controller) AdminBlockProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var blockProduct models.Product

		json.NewDecoder(r.Body).Decode(&blockProduct)

		blockProduct, err := c.ProductRepo.BlockProduct(blockProduct)

		if err != nil {
			log.Println("unable to update product")
			w.WriteHeader(http.StatusNotModified)
			json.NewEncoder(w).Encode(utils.PrepareResponse(false, "unable to update product", err))
			return
		}

		log.Println("updated the product active status")
		json.NewEncoder(w).Encode(utils.PrepareResponse(true, "updated the product status by admin", &blockProduct))

	}
}

func (c Controller) AdminLogout(w http.ResponseWriter, r *http.Request) {
	//logout function for the admin goes here.
}
