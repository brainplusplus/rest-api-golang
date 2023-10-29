package routes

import (
	"simple-ecommerce/controllers"
	"simple-ecommerce/middlewares"
)

func SetupApiRoutes() {
	app.Get("/", controllers.AuthApiController.WelcomeGuest)
	app.Post("/api/auth/login", controllers.AuthApiController.LoginCustomer)
	//register customer
	app.Post("/api/auth/register", controllers.AuthApiController.Register)
	app.Post("/api/auth/login_admin", controllers.AuthApiController.LoginAdmin)

	apiAuth := app.Group("/api/customer", middlewares.JWTMiddleware)
	apiAuth.Get("/product", controllers.ProductApiController.List)
	apiAuth.Get("/product/:id", controllers.ProductApiController.Get)
	//Task no 2
	apiAuth.Post("/order", controllers.OrderApiController.ByCustomerWantSave)
	//End Task no 2

	//Task no 3
	apiAuth.Get("/order", controllers.OrderApiController.ByCustomerWantList)
	apiAuth.Get("/order/:id", controllers.OrderApiController.ByCustomerWantGet)
	//End Task no 3

	apiAdmin := app.Group("/api/admin", middlewares.JWTMiddleware, middlewares.IsAdminMiddleware)

	//Task no 4
	apiAdmin.Get("/order", controllers.OrderApiController.List)
	//End Task no 4

	apiAdmin.Get("/customer", controllers.CustomerApiController.List)
	apiAdmin.Get("/customer/:id", controllers.CustomerApiController.Get)
	apiAdmin.Post("/customer", controllers.CustomerApiController.Save)
	apiAdmin.Put("/customer/:id", controllers.CustomerApiController.Update)
	apiAdmin.Delete("/customer/:id", controllers.CustomerApiController.Delete)

	apiAdmin.Get("/admin_user", controllers.AdminApiController.List)
	apiAdmin.Get("/admin_user/:id", controllers.AdminApiController.Get)
	apiAdmin.Post("/admin_user", controllers.AdminApiController.Save)
	apiAdmin.Put("/admin_user/:id", controllers.AdminApiController.Update)
	apiAdmin.Delete("/admin_user/:id", controllers.AdminApiController.Delete)

	apiAdmin.Get("/product", controllers.ProductApiController.List)
	apiAdmin.Get("/product/:id", controllers.ProductApiController.Get)
	apiAdmin.Post("/product", controllers.ProductApiController.Save)
	apiAdmin.Put("/product/:id", controllers.ProductApiController.Update)
	apiAdmin.Delete("/product/:id", controllers.ProductApiController.Delete)
}
