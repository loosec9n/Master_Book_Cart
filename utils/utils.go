package utils

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/http"
)

func HashPassword(password string) string {
	data := []byte(password)
	password = fmt.Sprintf("%x", md5.Sum(data))
	return password
}

func VerifyPassword(requestedPassword, dbPassword string) bool {
	requestedPassword = fmt.Sprintf("%x", md5.Sum([]byte(requestedPassword)))
	// log.Println("Requested Password", requestedPassword)
	// log.Println("DB Password", dbPassword)
	return requestedPassword == dbPassword
}

func ResponseJSON(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}
