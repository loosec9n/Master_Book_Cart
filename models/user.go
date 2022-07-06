package models

import "time"

//User -> Struct for defining user
type User struct {
	User_ID       int           `json:"id"`
	First_Name    string        `json:"firstName"`
	Last_Name     string        `json:"lastName"`
	Password      string        `json:"password"`
	Email         string        `json:"email"`
	Phone_Number  int           `json:"phoneNumber"`
	IsAdmin       bool          `json:"isAdmin"`
	UserID        int           `json:"userID"`
	UserCart      []ProductUser `json:"userCart"`
	UserAddressID []Address     `json:"addressID"`
	OrderStatus   []Order       `json:"order"`
	CreatedAt     time.Time     `json:"userCreatedAt"`
	UpdatedAt     time.Time     `json:"userUpdatedAt"`
	DeletedAt     time.Time     `json:"userDeletedAt"`
}
