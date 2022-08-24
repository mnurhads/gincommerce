package main

import (
	"gincommerce/controllers"
	"gincommerce/database"
	"gincommerce/middleware"
	"gincommerce/routes"
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

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	// The authentication middleware is applied to all routes, including the /users/signup route. So nobody can actually use the application.
	router.Use(middleware.Authentication())
	// Your routes are inconsistent starting with and without '/'.
	router.GET("/add/cart", app.AddToCart())
	router.GET("/remove/item", app.RemoveItem())
	router.GET("/list/cart", controllers.GetItemFromCart())
	router.POST("/add/address", controllers.AddAddress())
	router.PUT("/edit/homeaddress", controllers.EditHomeAddress())
	router.PUT("/edit/workaddress", controllers.EditWorkAddress())
	router.GET("/delete/addresses", controllers.DeleteAddress())
	router.GET("/cart/checkout", app.BuyFromCart())
	router.GET("/instant/buy", app.InstantBuy())
	//router.GET("logout", controllers.Logout())
	//break :)

	// Log the error that the router can possibly return.
	log.Fatal(router.Run(":" + port))
}
