package main

import (
	"go-first/config"
	"go-first/controller"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.DBInit()
	inDB := &controller.InDB{DB: db}

	router := gin.Default()

	router.POST("/orders", inDB.CreateOrder)
	router.GET("/orders", inDB.GetOrders)
	router.GET("/order/:id", inDB.GetOrder)
	router.PUT("/order/:id", inDB.UpdateOrder)
	router.DELETE("/order/:id", inDB.DeleteOrder)
	router.Run(":3000")
}
