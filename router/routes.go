package router

import (
	h "github.com/fekom/go-api/handlers"
	"github.com/gin-gonic/gin"
)

func initializerRoutes(router *gin.Engine) {
	router.GET("/search/:search_team", h.GetValues)
}
