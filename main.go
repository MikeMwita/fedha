package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	log.Println("...starting our app")
	r := gin.Default()

	r.Run(":5000")
}
