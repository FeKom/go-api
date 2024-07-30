package router

import "github.com/gin-gonic/gin"

func InitializeServer() {
	// Initializer router
	router := gin.Default()

	//
	initializerRoutes(router)

	//run the server
	router.Run(":3000")
}
