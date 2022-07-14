package routes

import (
	"Book_Cart_Project/controllers"
	"Book_Cart_Project/middleware"

	"github.com/go-chi/chi/v5"
)

func AdminRoute(routes chi.Router, Controller controllers.Controller) {

	//routes.Get("/admin/login", Controller.AdminLoginIndex)
	routes.Post("/admin/login", Controller.AdminLogin())

	routes.Group(func(r chi.Router) {
		r.Use(middleware.TokenVerifyMiddleware)
		//r.Get("/admin/logout", Controller.AdminLogout)
		r.Post("/admin/addproduct", Controller.AdminProductAdd())
		r.Get("/admin/viewproduct", Controller.AdminProductView())
		r.Post("/admin/blockuser", Controller.AdminBlockUser())
		r.Get("/admin/viewuser", Controller.AdminViewUser())
		r.Post("/admin/add/category", Controller.AddCategory())
		r.Get("/admin/view/category", Controller.ViewCategory())
		r.Post("/admin/add/author", Controller.AddAuthor())
		r.Get("/admin/view/author", Controller.ViewAuthor())
	})
}
