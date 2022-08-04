package controllers

import (
	"Book_Cart_Project/models"
	"Book_Cart_Project/utils"
	"encoding/json"
	"log"
	"net/http"
)

func (c Controller) CreateOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var orderBody models.OrderBody
		var orders models.Order
		json.NewDecoder(r.Body).Decode(&orderBody)

		order, totalAmount, err := c.UserRepo.CreateNewOrder(orderBody, orders)
		//fmt.Println("i m here")
		if err != nil {
			log.Println("Error running query: Order was not created")
			w.WriteHeader(http.StatusNotImplemented)
			json.NewEncoder(w).Encode(utils.PrepareResponse(false, "error is :", err))
			return
		}

		log.Println("User order placed")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(utils.PrepareResponse(true, "user order placed :", map[string]interface{}{
			"Total Amount": totalAmount,
			"Checkout":     &order,
		}))
	}
}
func (c Controller) UserOrderPaymnet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var payment models.Payment
		json.NewDecoder(r.Body).Decode(&payment)

		err := c.UserRepo.OrderPayments(payment)
		if err != nil {
			log.Println("Error running query")
			w.WriteHeader(http.StatusNotImplemented)
			json.NewEncoder(w).Encode(utils.PrepareResponse(false, "error is :", err))
			return
		}

		log.Println("Paymnet selected")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(utils.PrepareResponse(true, "cod payment updates", nil))
	}
}

func (c Controller) OrderPlaced() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var orderBody models.OrderBody

		json.NewDecoder(r.Body).Decode(&orderBody)

		err := c.UserRepo.OrderedProduct(orderBody)

		if err != nil {
			log.Println("Error running query: order was not confirmed")
			w.WriteHeader(http.StatusNotImplemented)
			json.NewEncoder(w).Encode(utils.PrepareResponse(false, "error is :", err))
			return
		}

		log.Println("Order confirmed")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(utils.PrepareResponse(true, "Order confirmed", nil))
	}
}
