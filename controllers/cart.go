package controllers

import (
	"Book_Cart_Project/models"
	"Book_Cart_Project/utils"
	"encoding/json"
	"log"
	"net/http"
)

func (c Controller) AddToCart() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var cart models.Cart

		//fetching data
		json.NewDecoder(r.Body).Decode(&cart)

		cart, err := c.UserRepo.AddToCart(cart)

		if err != nil {
			log.Println("Error adding to cart")
			w.WriteHeader(http.StatusNotImplemented)
			json.NewEncoder(w).Encode(utils.PrepareResponse(false, "error adding to cart", err))
			return
		}

		log.Println("Product added to cart")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(utils.PrepareResponse(true, "sucessfully added to cart", &cart))
	}
}

func (c Controller) ViewCart() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var cart models.Cart

		json.NewDecoder(r.Body).Decode(&cart)

		carts, total, err := c.UserRepo.ViewCart(cart)

		if err != nil {
			log.Println("Error viewing cart")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(utils.PrepareResponse(false, "errro viewing cart", err))
			return
		}

		log.Println("Success viewing cart")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(utils.PrepareResponse(true, "sucess viewing cart", map[string]interface{}{
			"Total Price":      total,
			"Products in cart": &carts,
		}))
	}
}
