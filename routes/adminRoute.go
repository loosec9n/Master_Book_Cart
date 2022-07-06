package routes

import (
	"Book_Cart_Project/controllers"
	"Book_Cart_Project/middleware"

	"github.com/go-chi/chi/v5"
)

func AdminRoute(routes chi.Router, Controller controllers.Controller) {

	routes.Get("/admin/login", Controller.AdminLoginIndex)
	routes.Post("/admin/login", Controller.AdminLogin())
	routes.Get("/admin/logout", Controller.AdminLogout)

	routes.Group(func(r chi.Router) {
		r.Use(middleware.TokenVerifyMiddleware)
		r.Post("/admin/addproduct", Controller.AdminProductAdd())
		r.Get("/admin/viewproduct", Controller.AdminProductView())
	})
}
