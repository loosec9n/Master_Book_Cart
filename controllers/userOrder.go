package controllers

import (
	"Book_Cart_Project/models"
	"Book_Cart_Project/utils"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/razorpay/razorpay-go"
	"github.com/thanhpk/randstr"
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
func (c Controller) UserRazorIndex(w http.ResponseWriter, r *http.Request) {
	//user template needed to parse here

	userID, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		log.Println("Error converting string to int", err)
		w.WriteHeader(http.StatusNotImplemented)
		json.NewEncoder(w).Encode(utils.PrepareResponse(false, "Error converting string to int", err))
		return
	}

	client := razorpay.NewClient("rzp_test_UPVZaukIo1yTq7", "FWAxiPyt6guT3ZM3g9zcg5eS")

	data := map[string]interface{}{
		"amount":   "1000",
		"currency": "INR",
		"receipt":  randstr.String(5),
	}
	body, err := client.Order.Create(data, nil)
	if err != nil {
		log.Println("Order creation failed", err)
		return
	}

	//save the orderid from the body
	value := body["id"]

	str := value.(string)

	//query to get the details form cart
	rzrPay, err := c.UserRepo.PaymentMod(userID)
	if err != nil {
		log.Println("no data from PaymentMod", err)
	}

	//log.Println("PageVariable struct:", rzrPay)

	HomePageVars := models.PageVariable{
		OrderID:     str,
		UserName:    rzrPay.UserName,
		Email:       rzrPay.Email,
		PhoneNumber: rzrPay.PhoneNumber,
		TotalAmount: rzrPay.TotalAmount,
	}
	//log.Println("total amount :", rzrPay.TotalAmount)
	t, err := template.ParseFiles("static/app.html")

	if err != nil {
		log.Println("template parsing error", err)
	}

	err = t.Execute(w, HomePageVars)
	if err != nil {
		log.Println("template executing error ", err)
	}
}

func (c Controller) UserRazorPayment() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (c Controller) OrderPlaced() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var orderBody models.OrderBody
		var orders models.Cart

		json.NewDecoder(r.Body).Decode(&orderBody)

		err := c.UserRepo.OrderedProduct(orderBody, orders)

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
