package utils

import (
	"crypto/md5"
	"fmt"
)

func HashPassword(password string) string {
	data := []byte(password)
	password = fmt.Sprintf("%x", md5.Sum(data))
	return password
}

func VerifyPassword(requestedPassword, dbPassword string) bool {
	requestedPassword = fmt.Sprintf("%x", md5.Sum([]byte(requestedPassword)))
	return requestedPassword == dbPassword
}
