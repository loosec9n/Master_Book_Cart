package middleware

import (
	"Book_Cart_Project/token"
	"Book_Cart_Project/utils"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func TokenVerifyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//getting token from header
		autheader := r.Header.Get("Authorization")
		bearerToken := strings.Split(autheader, " ")

		if len(bearerToken) == 2 {
			authToken := bearerToken[1]

			err := token.ValidateToken(authToken)

			if !err {
				log.Println("Token Invalid.")
				w.WriteHeader(http.StatusUnauthorized)

			} else if err {
				next.ServeHTTP(w, r)
			}

		} else {
			log.Println("Invalid Token")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(utils.PrepareResponse(false, "Inavalid Token", nil))
		}
	})
}
