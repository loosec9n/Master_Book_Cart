package models

// type for jwt
type JWT struct {
	Token         string `json:"jwt_token"`
	Refresh_Token string `json:"jwt_refresh_token"`
}
