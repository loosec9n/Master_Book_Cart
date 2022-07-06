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
			json.NewEncoder(w).Encode(utils.PrepareResponse(false, "User already exists", err))
			//utils.ResponseJSON(w, "User already exists.")
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
		//utils.ResponseJSON(w, &user)
		json.NewEncoder(w).Encode(utils.PrepareResponse(true, "new user created", &user))
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

		//checking the user is active or not
		if !founduser.Is_Active {
			log.Println("User has been deactivated by Admin")
			w.WriteHeader(http.StatusUnauthorized)
			//utils.ResponseJSON(w, "This user is inactive, please contact the admin")
			json.NewEncoder(w).Encode(utils.PrepareResponse(true, "This user is inactive, please contact the admin", founduser))
			return
		}

		//checking the error in the sql query
		if err != nil {
			if err == sql.ErrNoRows {
				log.Println("The user does not exists.")
				w.WriteHeader(http.StatusBadGateway)
				//utils.ResponseJSON(w, "Please Check the Username")
				json.NewEncoder(w).Encode(utils.PrepareResponse(true, "User does not exists", err))
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
			log.Println("Invalid user password.")
			//utils.ResponseJSON(w, "Inavlid Password")
			w.WriteHeader(http.StatusBadGateway)
			json.NewEncoder(w).Encode(utils.PrepareResponse(true, "Invalid user password", founduser))
			return
		}

		//generationg jwt token
		token, refresh_token := token.GenerateToken(founduser.First_Name, founduser.Last_Name, founduser.Email, founduser.Phone_Number, founduser.UserID)

		jwt.Token = token
		jwt.Refresh_Token = refresh_token

		w.WriteHeader(http.StatusOK)
		//	utils.ResponseJSON(w, jwt)
		log.Println("User login sucessful")
		json.NewEncoder(w).Encode(utils.PrepareResponse(true, "User login sucessful", jwt))

	}
}

func (c Controller) UserHomePage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Logged in Successfully")
		//utils.ResponseJSON(w, "Succesfully logged in")
		json.NewEncoder(w).Encode(utils.PrepareResponse(true, "User Home page", "Sucessful login"))

	}
}

func (c Controller) UserLogout(w http.ResponseWriter, r *http.Request) {
	//user logout function goes here ..
}

func (c Controller) HomePage(w http.ResponseWriter, r *http.Request) {
	//homepage will parse here ..
}
