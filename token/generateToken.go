package token

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var SECRET_KEY = os.Getenv("SECRET")

type SignedDetails struct {
	User_ID      int
	First_Name   string
	Last_Name    string
	Email        string
	Phone_Number int
	jwt.StandardClaims
}

// to generate token
func GenerateToken(first_name string, last_name string, email string, phone_number int, user_id int) (signedtoken string, signedrefreshtoken string) {

	claims := &SignedDetails{
		User_ID:      user_id,
		First_Name:   first_name,
		Last_Name:    last_name,
		Email:        email,
		Phone_Number: phone_number,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshclaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	//Generating token
	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))

	refreshtoken, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshclaims).SignedString([]byte(SECRET_KEY))

	return token, refreshtoken
}

//to verify token
func ValidateToken(signedtoken string) bool {
	//validating token
	token, err := jwt.ParseWithClaims(
		signedtoken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		log.Fatalln(err)
	}

	claims, ok := token.Claims.(*SignedDetails)

	if !ok {
		log.Println("The Token is Invalid")
		return false
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		log.Println("Token Expired.")
		return false
	}
	return true
}
