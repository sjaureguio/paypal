package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sjaureguio/paypal/router"
	"github.com/sjaureguio/paypal/storage/postgres"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	loadEnv()
	startServerWithEcho()
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Print("no se pudieron cargar las variables de entorno")
		panic(err)
	}
}

func startServerWithEcho() {
	db, err := postgres.New()
	if err != nil {
		log.Print("error de conexion con la base de datos")
		panic(err)
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	router.Product(e, db)
	router.Order(e, db)
	router.Subscription(e, db)
	router.Invoice(e, db)
	router.PayPal(e, db)

	port := os.Getenv("HTTP_PORT")

	if port == "" {
		port = "8080"
	}

	// Start server
	log.Print(e.Start(port))

	//log.Print(
	//	e.StartTLS(
	//		port,
	//		os.Getenv("SSL_PUBLIC_KEY"),
	//		os.Getenv("SSL_PRIVATE_KEY"),
	//	),
	//)
}
