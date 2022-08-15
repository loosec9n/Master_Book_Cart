package main

import (
	"log"
	"net/http"
	"os"

	"Book_Cart_Project/controllers"
	"Book_Cart_Project/database"
	"Book_Cart_Project/repository"
	"Book_Cart_Project/routes"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"

	//razorpay "github.com/razorpay/razorpay-go"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {
	file, err := os.OpenFile("loggingData.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)

	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.Println("New Run!")

	//checking if env is avaialble if not using the default value
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	//Connecting to database

	dataBase := database.ConnectDB()

	//initializing interface in controllers
	controller := &controllers.Controller{
		ProductRepo: repository.Repository{
			DB: dataBase,
		},
		UserRepo: repository.Repository{
			DB: dataBase,
		},
	}

	//Creating an instance of Chi router
	router := chi.NewRouter()
	//using the logger from chi router
	router.Use(middleware.Logger)

	//Admin and User routes defined and injecting db.
	routes.AdminRoute(router, *controller)
	routes.UserRoute(router, *controller)

	//Server intiaised
	log.Println("API is listening in the port: ", port)

	http.ListenAndServe(":"+port, router)

}
