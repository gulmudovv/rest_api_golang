package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	memoryStorage := NewMemoryStorage()
	fmt.Printf("memoryStorage:ТИП%T=======%v\n", memoryStorage, memoryStorage)
	handler := NewHandler(memoryStorage)
	fmt.Printf("handler:TYPE:%T========%v\n", handler, handler)

	router := gin.Default()

	router.POST("/employee", handler.CreateEmployee)
	router.GET("/employee/:id", handler.GetEmployee)
	//router.PUT("/employee/:id", handler.UpdateEmployee)
	//router.DELETE("/employee/:id", handler.DeleteEmployee)

	router.Run()
}
