package main

import (
	"go-gin-4/controllers/fruit"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// r.GET("/fruits", )
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello")
	})

	r.GET("/fruits", fruit.GetAllFruit)
	r.GET("/fruits/:id", fruit.GetFruitById)
	r.POST("/fruits", fruit.CreateFruit)
	r.PUT("/fruits/:id", fruit.UpdateFruit)
	r.DELETE("/fruits/:id", fruit.DeleteFruit)

	r.Run(":8080")
}
