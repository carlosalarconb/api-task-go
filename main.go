package main

import (
	router "api-tasks/routers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	port := ":8000"
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	fmt.Printf("Running on port" + port)

	r.GET("/task", router.GetAll)
	r.GET("/task/:id", router.GetOne)

	r.Run(port)
}
