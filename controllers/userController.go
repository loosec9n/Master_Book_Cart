package controllers

import (
	"Book_Cart_Project/models"
	"Book_Cart_Project/token"
	"Book_Cart_Project/utils"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func (c Controller) UserSignUpIndex(w http.ResponseWriter, r *http.Request) {
	//user template needed to parse here
}

func (c Controller) UserSignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User

		//getting values from json
		json.NewDecoder(r.Body).Decode(&user)

		log.Println("Checking whether user exists.")
		err := c.UserRepo.DoesUserExists(user)

		if err {
			log.Println("User Already Exists.")
			w.WriteHeader(http.StatusBadRequest)
			//json.NewEncoder(w).Encode("User already exists.")
			utils.ResponseJSON(w, "User already exists.")
			return
		}

		// hashing password using bcrypt
		hashedPassword := utils.HashPassword(user.Password)

		// assigning hashed password to usermodels
		user.Password = hashedPassword

		log.Println("Hashed Password: ", user.Password)

		//creating update time
		user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		//Registering user to the database
		user = c.UserRepo.UserSignup(user)

		//writing userdata to the response
		log.Println("Signed Up Successfully")
		utils.ResponseJSON(w, &user)
	}
}

func (c Controller) UserLoginIndex(w http.ResponseWriter, r *http.Request) {
	//user login template will parse here
}

func (c Controller) UserLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var founduser models.User
		var jwt models.JWT

		//Decoding from request body
		json.NewDecoder(r.Body).Decode(&founduser)

		requestPassword := founduser.Password

		log.Println("Checking whether User exists.")

		founduser, err := c.UserRepo.UserLogin(founduser)

		if err != nil {
			if err == sql.ErrNoRows {
				log.Println("The user does not exists.")
				w.WriteHeader(http.StatusBadGateway)
				utils.ResponseJSON(w, "Please Check the Username")
			} else {
				log.Fatal(err)
			}
		}

		log.Println("User exists.")

		//getting hashed password from database
		dbPassword := founduser.Password

		//verifying password
		passwordMatch := utils.VerifyPassword(requestPassword, dbPassword)

		if !passwordMatch {
			log.Println("Invalid Password.")
			utils.ResponseJSON(w, "Inavlid Password")
			w.WriteHeader(http.StatusBadGateway)
			return
		}

		//generationg jwt token
		token, refresh_token := token.GenerateToken(founduser.First_Name, founduser.Last_Name, founduser.Email, founduser.Phone_Number, founduser.UserID)

		jwt.Token = token
		jwt.Refresh_Token = refresh_token

		w.WriteHeader(http.StatusOK)
		utils.ResponseJSON(w, jwt)

	}
}

func (c Controller) UserHomePage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Logged in Successfully")
		// json.NewEncoder(w).Encode("Successfully logged in")
		utils.ResponseJSON(w, "Succesfully logged in")
	}
}

func (c Controller) UserLogout(w http.ResponseWriter, r *http.Request) {
	//user logout function goes here ..
}

func (c Controller) HomePage(w http.ResponseWriter, r *http.Request) {
	//homepage will parse here ..
}
