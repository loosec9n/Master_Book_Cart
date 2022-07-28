package models

//User -> Struct for defining user
type User struct {
	User_ID      int    `json:"id"`
	Is_Active    bool   `json:"isActive"`
	First_Name   string `json:"firstName"`
	Last_Name    string `json:"lastName"`
	Password     string `json:"password,omitempty"`
	Email        string `json:"email"`
	Phone_Number int    `json:"phoneNumber"`
	IsAdmin      bool   `json:"isAdmin"`
	UserID       int    `json:"userID,omitempty"`
	//UserCart      []ProductUser `json:"userCart"`
	UserAddressID []Address `json:"addressID,omitempty"`
	OrderStatus   []Order   `json:"order,omitempty"`
	//CreatedAt     time.Time `json:"userCreatedAt,omitempty"`
	// UpdatedAt     time.Time     `json:"userUpdatedAt"`
	// DeletedAt     time.Time     `json:"userDeletedAt"`
}
