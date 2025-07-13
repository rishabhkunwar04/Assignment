package main

import (
	"log"
	"targeting-engine/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode) // production-ready mode
	r := gin.New()
	r.Use(gin.Recovery())

	handler.RegisterRoutes(r)
	log.Println("Targeting Engine running on http://localhost:8080")
	log.Fatal(r.Run(":8080"))
}
