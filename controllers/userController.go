package controllers

import (
	"Book_Cart_Project/models"
	"Book_Cart_Project/token"
	"Book_Cart_Project/utils"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/thanhpk/randstr"
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

		//log.Println("Hashed Password: ", user.Password)

		//creating update time
		// user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

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

		//checking the user is active or not
		if !founduser.Is_Active {
			log.Println("User has been deactivated by Admin")
			w.WriteHeader(http.StatusUnauthorized)
			//utils.ResponseJSON(w, "This user is inactive, please contact the admin")
			json.NewEncoder(w).Encode(utils.PrepareResponse(true, "This user is inactive, please contact the admin", nil))
			return
		}

		log.Println("User exists.")

		//getting hashed password from database
		dbPassword := founduser.Password

		//verifying password
		ok := utils.VerifyPassword(requestPassword, dbPassword)

		if !ok {
			log.Println("Invalid user password.")
			//utils.ResponseJSON(w, "Inavlid Password")
			w.WriteHeader(http.StatusBadGateway)
			json.NewEncoder(w).Encode(utils.PrepareResponse(true, "Invalid user password", founduser))
			return
		}

		//generationg jwt token
		token, refresh_token := token.GenerateTokenUser(founduser.First_Name, founduser.Last_Name, founduser.Email, founduser.Phone_Number, founduser.User_ID, founduser.IsAdmin)

		jwt.Token = token
		jwt.Refresh_Token = refresh_token

		w.WriteHeader(http.StatusOK)
		log.Println("User login sucessful")
		json.NewEncoder(w).Encode(utils.PrepareResponse(true, "User login sucessful", jwt))

	}
}

func (c Controller) UserHomePage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//getting UserId form thr JWT payload

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
		//searchParam := chi.URLParam(r, "search")
		products, result, err := c.ProductRepo.ViewProduct(f)

		if err != nil {
			log.Println("Error - no query execution - Product View", err)
			w.WriteHeader(http.StatusNotImplemented)
			json.NewEncoder(w).Encode(utils.PrepareResponse(false, "Error - no query execution - Product View", err))
		}

		log.Println("Sucess in Viewing Products")
		json.NewEncoder(w).Encode(utils.PrepareResponse(true, "home page view products success", map[string]interface{}{
			"total products": result.TotalRecords,
			"current page":   result.CurrentPage,
			"total pages":    result.LastPage,
			"data":           &products,
		}))

		// log.Println("Logged in Successfully")
		// //utils.ResponseJSON(w, "Succesfully logged in")
		// json.NewEncoder(w).Encode(utils.PrepareResponse(true, "User Home page", "Sucessful login"))

	}
}

func (c Controller) SearchProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		param, err := strconv.Atoi(r.URL.Query().Get("search"))
		if err != nil || param < 1 {
			log.Println("Error converting string to int", err)
			w.WriteHeader(http.StatusNotImplemented)
			json.NewEncoder(w).Encode(utils.PrepareResponse(false, "Error converting string to int", err))
			return
		}

		//checking wheather product is active or not
		activeProd, err := c.UserRepo.CheckActiveProd(param)

		if !activeProd {
			log.Println("Product is not active")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(utils.PrepareResponse(false, "product is not active", err))
			return
		}

		product, err := c.UserRepo.UserSearchProduct(param)

		if err != nil {
			log.Println("error executing query for the user products search")
			w.WriteHeader(http.StatusNotImplemented)
			json.NewEncoder(w).Encode(utils.PrepareResponse(false, "error exectuing query for the user products search", err))
			return
		}

		log.Println("search product is visible")
		w.WriteHeader(http.StatusFound)
		json.NewEncoder(w).Encode(utils.PrepareResponse(true, "search product is visible", &product))

	}
}

func (c Controller) ForgetPassword() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userCred models.User

		json.NewDecoder(r.Body).Decode(&userCred)

		//finding the user by email
		user, err := c.UserRepo.FindUserByEmail(userCred)

		//checking is user sigedup or not
		if err != nil {
			log.Println("user is not present in the database")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(utils.PrepareResponse(false, "user is not present in the database", err))
			return
		}

		//checking the user is active in database or not
		if !user.Is_Active {
			log.Println("not an active user")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(utils.PrepareResponse(false, "not an active user", nil))
			return
		}

		//generating random code for resetting password
		resetToken := randstr.String(20)

		//encoding the random string
		//passwordResetToken := utils.Encode(resetToken)\
		passwordResetToken := utils.HashPassword(resetToken)

		//update the hashed password into the database
		modUser, err := c.UserRepo.ForgetPasswordUpdate(user, passwordResetToken)

		if err != nil {
			log.Println("not able to update the dummy password", err)
			w.WriteHeader(http.StatusNotImplemented)
			json.NewEncoder(w).Encode(utils.PrepareResponse(false, "not able to update the password in the user table", &modUser))
			return
		}

		log.Println("password hashed and updated in user - forget password ")
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(utils.PrepareResponse(true, "password hashed and updated in user - sending the email", map[string]interface{}{
			"user":  &modUser,
			"token": resetToken,
		}))

		//send the email for resetting the password
		sendEmailData := utils.EmailData{
			URL:       "localhost:8080/user/forget/password/" + resetToken,
			FirstName: user.First_Name,
			Subject:   "Your password reset link",
		}

		err = utils.SendEmail(user, &sendEmailData)
		if err != nil {
			log.Println("error sending the email", err)
			w.WriteHeader(http.StatusNotImplemented)
			json.NewEncoder(w).Encode(utils.PrepareResponse(false, "error sending the email", err))
			return
		}

		log.Println("email send for reseting the data")
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(utils.PrepareResponse(true, "email send for reseting the data", map[string]interface{}{
			"user":  &modUser,
			"token": resetToken,
		}))
	}
}

func (c Controller) ResetPassword() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//get resetToken from params
		var resetToken models.ResetPasswordInput
		var userCred models.User

		json.NewDecoder(r.Body).Decode(&resetToken)

		userCred.Email = resetToken.Email

		//checking the password match
		hashedResetPassword := resetToken.PasswordConfirm

		//verify the token with users in db
		user, err := c.UserRepo.FindUserByEmail(userCred)

		//checking the user is active in database or not
		if err != nil {
			log.Println("user is not present in the database")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(utils.PrepareResponse(false, "user is not present in the database", err))
			return
		}
		if !user.Is_Active {
			log.Println("not an active user")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(utils.PrepareResponse(false, "not an active user", nil))
			return
		}

		ok := utils.VerifyPassword(resetToken.TokenPassword, user.Password)

		if !ok {
			log.Println("Invalid user password.")
			w.WriteHeader(http.StatusBadGateway)
			json.NewEncoder(w).Encode(utils.PrepareResponse(true, "Invalid user password", &user.First_Name))
			return
		}

		log.Println("Token and password match : updating the new password")
		updateUser, err := c.UserRepo.ForgetPasswordUpdate(user, hashedResetPassword)
		if err != nil {
			log.Println("not able to update the password in the user table")
			w.WriteHeader(http.StatusNotImplemented)
			json.NewEncoder(w).Encode(utils.PrepareResponse(false, "ot able to update the password in the user table", &updateUser))
			return
		}

		log.Println("new password updated")
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(utils.PrepareResponse(true, "new password updated", &updateUser))

	}
}

func (c Controller) UserLogout(w http.ResponseWriter, r *http.Request) {
	//user logout function goes here ..
}

func (c Controller) HomePage(w http.ResponseWriter, r *http.Request) {
	//homepage will parse here ..
}
