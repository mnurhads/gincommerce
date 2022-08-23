package main

import (
	"ecommerce/controllers"
	"ecommerce/database"
	"ecommerce/middleware"
	"ecommerce/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// But you should do this for all config: mongodb (credentials, database, collections), SECRET_KEY in tokengen.go.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
}
