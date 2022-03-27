package main

import (
	"github.com/gin-gonic/gin"
	"swapnildaddikar07/request-decryption-middleware/controller"
	"swapnildaddikar07/request-decryption-middleware/middleware"
)

func main() {

	router := gin.Default()

	credentialChangeController := controller.NewCredentialChangeController()

	routerGroup := router.Group("/api")
	routerGroup.POST("/change-password", middleware.DecryptRequest, credentialChangeController.CredentialChange)

	router.Run(":8080")
}
