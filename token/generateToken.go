package token

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type SignedDetails struct {
	User_ID      int
	First_Name   string
	Last_Name    string
	Email        string
	Phone_Number int
	Is_Admin     bool
	jwt.StandardClaims
}

// to generate token
func GenerateTokenAdmin(first_name string, last_name string, email string, phone_number int, user_id int, is_admin bool) (signedtoken string, signedrefreshtoken string) {

	var secret_key = os.Getenv("SECRET_ADMIN")

	//log.Println("jwt Secret key: ", secret_key)

	claims := &SignedDetails{
		User_ID:      user_id,
		First_Name:   first_name,
		Last_Name:    last_name,
		Email:        email,
		Phone_Number: phone_number,
		Is_Admin:     is_admin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshclaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	//Generating token
	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret_key))

	refreshtoken, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshclaims).SignedString([]byte(secret_key))

	return token, refreshtoken
}

//to verify token
func ValidateTokenAdmin(signedtoken string) (bool, error) {
	//validating token
	var secret_key = os.Getenv("SECRET_ADMIN")
	//var userID int

	token, err := jwt.ParseWithClaims(
		signedtoken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secret_key), nil
		},
	)

	if err != nil {
		log.Println("error parsing jwt: ", err)
		return false, err
	}

	claims, ok := token.Claims.(*SignedDetails)
	//log.Println(claims.User_ID)

	if !ok {
		log.Println("The Token is Invalid")
		return false, err
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		log.Println("Token Expired.")
		return false, err
	}
	return true, nil
}
