package models

//User -> Struct for defining user
type User struct {
	User_ID      int    `json:"id,omitempty"`
	Is_Active    bool   `json:"isActive,omitempty"`
	First_Name   string `json:"firstName,omitempty"`
	Last_Name    string `json:"lastName,omitempty"`
	Password     string `json:"password,omitempty"`
	Email        string `json:"email,omitempty"`
	Phone_Number int    `json:"phoneNumber,omitempty"`
	IsAdmin      bool   `json:"isAdmin,omitempty"`
	UserID       int    `json:"userID,omitempty"`
	//UserCart      []ProductUser `json:"userCart"`
	UserAddressID []Address `json:"addressID,omitempty"`
	OrderStatus   []Order   `json:"order,omitempty"`
	//CreatedAt     time.Time `json:"userCreatedAt,omitempty"`
	// UpdatedAt     time.Time     `json:"userUpdatedAt"`
	// DeletedAt     time.Time     `json:"userDeletedAt"`
}
