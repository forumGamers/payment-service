package routes

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type routes struct {
	router *gin.Engine
}

func Routes(){

	if err := godotenv.Load() ; err != nil {
		panic(err.Error())
	}

	r := routes { router: gin.Default() }

	r.router.Use(func(c *gin.Context) {
		c.Set("start",time.Now())

		c.Next()
	})

	c := cors.New(cors.Config{
		AllowOrigins: []string{os.Getenv("CORSLIST")},
		AllowMethods: []string{"GET","POST","PUT","DELETE","PATCH","OPTIONS"},
		AllowHeaders: []string{"Content-Type","Authorization"},
		AllowCredentials: true,
	})

	r.router.Use(func (c *gin.Context){
		if c.Request.Method != "OPTIONS" {
			origin := c.Request.Header.Get("Origin")
			if origin == "" {
                c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "Forbidden"})
                return
            }
		}
		c.Next()
	})

	r.router.Use(c)

	r.router.Use(logger.SetLogger())

	//testing connection
	r.router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{"message":"pong"})
	})

	groupRoutes := r.router.Group("/api")

	r.transactionRoutes(groupRoutes)

	port := os.Getenv("PORT")

	if port == "" {
		port = "4300"
	}

	r.router.Run(":"+port)
}