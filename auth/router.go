package auth

import (
	"github.com/julienschmidt/httprouter"
	
)

func NewRouter(customerController controller.CustomerController, productController controller.ProductController, orderController controller.OrderController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/register", customerController.Register)
	router.POST("/api/login", customerController.Login)
	router.PUT("/api/users/:customerId", customerController.Update)
	router.DELETE("/api/users/:customerId", customerController.Delete)
	router.GET("/api/users/:customerId", customerController.FindById)

	router.GET("/api/products", productController.GetAll)
	router.POST("/api/products", productController.Create)
	router.GET("/api/products/:productId", productController.GetById)
	router.PUT("/api/products/:productId", productController.Update)
	router.DELETE("/api/products/:productId", productController.Delete)

	router.POST("/api/orders", orderController.Create)
	router.GET("/api/orders/:orderId", orderController.GetById)
	router.PUT("/api/orders/:orderId", orderController.Update)
	router.DELETE("/api/orders/:orderId", orderController.Delete)

	router.PanicHandler = exception.ErrorHandle

	return router
}