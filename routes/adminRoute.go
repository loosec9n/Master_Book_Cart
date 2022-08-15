package routes

import (
	"Book_Cart_Project/controllers"
	"Book_Cart_Project/middleware"

	"github.com/go-chi/chi/v5"
)

func AdminRoute(routes chi.Router, Controller controllers.Controller) {

	//routes.Get("/admin/login", Controller.AdminLoginIndex)
	routes.Post("/admin/login", Controller.AdminLogin())
	routes.Patch("/admin/delivery/status", Controller.AdminEditOrderStatus())

	routes.Group(func(r chi.Router) {
		r.Use(middleware.AdminVerifyMiddleware)
		//r.Get("/admin/logout", Controller.AdminLogout)
		r.Post("/admin/add/product", Controller.AdminProductAdd())
		r.Get("/admin/view/product", Controller.AdminProductView())
		r.Post("/admin/block/product", Controller.AdminBlockProduct())
		r.Post("/admin/blockuser", Controller.AdminBlockUser())
		r.Get("/admin/viewuser", Controller.AdminViewUser())
		r.Post("/admin/add/category", Controller.AddCategory())
		r.Get("/admin/view/category", Controller.ViewCategory())
		r.Post("/admin/add/author", Controller.AddAuthor())
		r.Get("/admin/view/author", Controller.ViewAuthor())
		r.Post("/admin/add/inventory", Controller.AdminAddInventory())
		r.Get("/admin/report/month", Controller.AdminReport())
		r.Patch("/admin/delivery/status", Controller.AdminEditOrderStatus())
	})
}
