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
	routes.Post("/user/signup/address", Controller.UserAddress())
	routes.Post("/user/login", Controller.UserLogin())
	routes.Get("/", Controller.HomePage)
	routes.Post("/user/forget/password", Controller.ForgetPassword())
	routes.Patch("/user/forget/password/reset", Controller.ResetPassword())
	routes.Get("/user/payment/razorpay", Controller.UserRazorIndex)
	routes.Get("/user/payment/success", Controller.RazorSuccessIndex)
	routes.Get("/user/payment-success", Controller.UserRazorPaySucess)

	routes.Group(func(r chi.Router) {
		r.Use(middleware.UserVerifyMiddleware)
		//r.Get("/user/logout", Controller.UserLogout)
		r.Get("/homepage/product", Controller.SearchProduct())
		r.Get("/homepage", Controller.UserHomePage())
		r.Post("/user/add/cart", Controller.AddToCart())
		r.Post("/user/view/cart", Controller.ViewCart())
		r.Post("/user/add/wishlist", Controller.UserWishlist())
		r.Get("/user/view/wishlist", Controller.UserViewWishlist())
		r.Delete("/user/remove/wishlist", Controller.UserDeleteWishlist())
		r.Get("/user/cart/checkout", Controller.CreateOrder())
		r.Post("/user/paymnet/cod", Controller.UserOrderPaymnet())
		r.Post("/user/order/confirm", Controller.OrderPlaced())

	})

}
