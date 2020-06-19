package main

import (
	"fmt"
	"go-workshop/src"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()
	health := engine.Group("/")

	userRoute := engine.Group("/user")
	userRoute.GET("/", src.GetUsers) // /user
	userRoute.GET("/:id", src.GetUser)
	userRoute.PUT("/:id", src.PutUser)
	userRoute.POST("/", src.PostUser)
	userRoute.DELETE("/:id", src.DeleteUser)

	debtsRoute := engine.Group("/debt")
	debtsRoute.POST("/", src.PostDebt)
	debtsRoute.GET("/", src.GetDebts)
	debtsRoute.GET("/:id", src.GetDebt)
	debtsRoute.PUT("/:id", src.PutDebt)
	debtsRoute.DELETE("/:id", src.DeleteDebt)
	src.AutoMigration()

	health.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Status": "Go healthy",
		})
	})

	engine.Run(fmt.Sprintf(":%v", os.Getenv("PORT")))
}
