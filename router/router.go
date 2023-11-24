package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//Initialize the router
	r := gin.Default()
	initializeRoutes(r)

	// Run the server
	r.Run(":8080")

}
