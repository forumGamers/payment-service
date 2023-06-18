package routes

import (
	"github.com/forumGamers/payment-service/cmd"
	md "github.com/forumGamers/payment-service/middlewares"
	"github.com/gin-gonic/gin"
)

func (r routes) transactionRoutes(rg *gin.RouterGroup){
	uri := rg.Group("/transaction")

	uri.
	Use(md.Authentication).
	POST("/",cmd.CreateTransaction)
}