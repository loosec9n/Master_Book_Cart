package models

type Address struct {
	UserAddress_ID uint   `json:"useraddressID,omitempty"`
	AddressID      int    `json:"addressID,omitempty"`
	HouseName      string `json:"houseName,omitempty"`
	StreetName     string `json:"streetName,omitempty"`
	LandMark       string `json:"landMark,omitempty"`
	City           string `json:"city,omitempty"`
	State          string `json:"state,omitempty"`
	Pincode        uint   `json:"pincode,omitempty"`
}
