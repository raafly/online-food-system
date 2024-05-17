package server

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/raafly/online-food-service/auth"
	"github.com/raafly/online-food-service/config"
	"github.com/raafly/online-food-service/database"
	"github.com/raafly/online-food-service/listing"
	"github.com/raafly/online-food-service/order"
)

type Server struct {
	App  *fiber.App
	conf *config.AppConfig
}

func NewServer() *Server {
	app := fiber.New()
	conf, err := config.NewAppConfig()
	if err != nil {
		log.Printf("failed initialized config %v", err)
	}

	return &Server{
		App:  app,
		conf: conf,
	}
}

func (s *Server) Run() {
	db := database.NewDB(s.conf)

	s.App.Use(cors.New(cors.Config{
        AllowMethods: "GET,POST,PUT,DELETE",
        AllowHeaders: "Content-Type",
        AllowCredentials: false, // Ubah menjadi "false" jika tidak diperlukan
	}))

	listing.NewProductRoutes(s.App, db)
	auth.NewCustomerRoutes(s.App, db)
	order.NewOrderRoutes(s.App, db)

	s.GracefulShutdown(s.conf.Fiber.Port)
}

func (s *Server) GracefulShutdown(port string) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := s.App.Listen(":" + port); err != nil {
			log.Fatalf("error when listening to :%s, %s", port, err)
		}
	}()

	log.Printf("server is running on :%s", port)

	<-stop

	log.Println("server gracefully shutdown")

	if err := s.App.Shutdown(); err != nil {
		log.Fatalf("error when shutting down the server, %s", err)
	}

	log.Println("process clean up...")
}
