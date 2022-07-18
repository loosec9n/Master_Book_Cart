package routes

import (
	"Book_Cart_Project/controllers"
	"Book_Cart_Project/middleware"

	"github.com/go-chi/chi/v5"
)

func UserRoute(routes chi.Router, Controller controllers.Controller) {
	//routes.Get("/user/signup", Controller.UserSignUpIndex)
	routes.Post("/user/signup", Controller.UserSignUp())
	//routes.Get("/user/login", Controller.UserLoginIndex)
	routes.Post("/user/login", Controller.UserLogin())
	routes.Get("/", Controller.HomePage)

	routes.Group(func(r chi.Router) {
		r.Use(middleware.TokenVerifyMiddleware)
		//r.Get("/user/logout", Controller.UserLogout)
		r.Post("/homepage/product", Controller.SearchProduct())
		r.Get("/homepage", Controller.UserHomePage())
		r.Post("/user/add/cart", Controller.AddToCart())
		r.Post("/user/view/cart", Controller.ViewCart())
	})

}
