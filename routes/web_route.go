package routes

import "simple-ecommerce/controllers"

func SetupWebRoutes() {
	app.Get("/order_list", controllers.OrderApiController.ListHtml)
	app.Get("/order/view/:token", controllers.OrderApiController.ByCustomerWantDetailHtml)
	app.Get("/order/complete/:token", controllers.OrderApiController.ByCustomerWantCompleteHtml)
}
