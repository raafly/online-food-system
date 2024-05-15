package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	
)

// func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
// 	return &http.Server{
// 		Addr:    "localhost:3000",
// 		Handler: authMiddleware,
// 	}
// }

func main() {

	db := config.NewDB()
	validate := validator.New()
	
	customerRepository := repository.NewCustomerRepository()
	customerService := service.NewCustomerService(customerRepository, db, validate)
	customerController := controller.NewCustomerController(customerService)

	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository, db, validate)
	productController := controller.NewProductController(productService)

	orderRepository := repository.NewOrderRepositoryImpl()
	orderService := service.NewOrderServiceImpl(orderRepository, db, validate)
	ordderController := controller.NewOrderController(orderService)
	

	router := route.NewRouter(customerController, productController, ordderController)
	authMiddleware := middleware.NewAuthMiddleware(router)

	
	server := http.Server{
		Addr: "localhost:3000",
		Handler: authMiddleware,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}