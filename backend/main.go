package main

import (
	"fmt"
	"os"

	"github.com/alanpv92/events/database"
	"github.com/alanpv92/events/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("could not load env varibles")
	}
	database.Init()
	server := gin.Default()
	port := fmt.Sprintf(":%v", os.Getenv("PORT"))
    routes.RegisterRoutes(server)
	fmt.Printf("Server is running at %v", port)
	err = server.Run(port)
	if err != nil {
		panic("could not start server")
	}

}
