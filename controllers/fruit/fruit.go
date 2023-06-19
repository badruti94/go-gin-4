package fruit

import (
	"fmt"
	"go-gin-4/db"
	"go-gin-4/models/fruit"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllFruit(ctx *gin.Context) {
	db := db.Connect()

	var fruits []fruit.Fruit

	db.Find(&fruits)

	ctx.JSON(http.StatusOK, gin.H{
		"data": fruits,
	})

}

func GetFruitById(ctx *gin.Context) {
	db := db.Connect()

	var fruit fruit.Fruit

	db.First(&fruit, ctx.Param("id"))

	ctx.JSON(http.StatusOK, gin.H{
		"data": fruit,
	})
}

func CreateFruit(ctx *gin.Context) {
	db := db.Connect()

	var fruit fruit.Fruit

	ctx.Bind(&fruit)

	fmt.Println(fruit)
	db.Create(&fruit)

	ctx.JSON(http.StatusOK, gin.H{
		"data": fruit,
	})

}

func UpdateFruit(ctx *gin.Context) {
	db := db.Connect()

	var fruit fruit.Fruit

	db.First(&fruit, ctx.Param("id"))

	ctx.Bind(&fruit)

	db.Save(&fruit)

	ctx.JSON(http.StatusOK, gin.H{
		"data": fruit,
	})
}

func DeleteFruit(ctx *gin.Context) {
	db := db.Connect()

	db.Delete(&fruit.Fruit{}, ctx.Param("id"))

	ctx.JSON(http.StatusOK, gin.H{
		"data": "success",
	})
}
