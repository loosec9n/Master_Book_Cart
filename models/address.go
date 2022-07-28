package models

type Address struct {
	UserAddress_ID uint   `json:"useraddressID"`
	AddressID      int    `json:"addressID"`
	HouseName      string `json:"houseName"`
	StreetName     string `json:"streetName"`
	LandMark       string `json:"landMark"`
	City           string `json:"city"`
	State          string `json:"state"`
	Pincode        uint   `json:"pincode"`
}
