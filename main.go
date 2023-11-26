package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Gorilla")
}

func main() {
	memoryStorage := NewMemoryStorage()
	handler := NewHandler(memoryStorage)

	router := gin.Default()
	router.POST("/employee", handler.CreateEmployee)
	router.GET("/get/:id", handler.GetEmployee)

	router.Run()
}
