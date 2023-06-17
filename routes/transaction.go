package routes

import "github.com/gin-gonic/gin"

func (r routes) transactionRoutes(rg *gin.RouterGroup){
	uri := rg.Group("/transaction")

	uri.GET("/")
}