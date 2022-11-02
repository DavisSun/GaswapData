package main

import (
	MainRouter "GaswapData/Route"
	runner "GaswapData/Runner"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	MainRouter.Setup(r)

	err := r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		log.Println(err)
	}

	runner.Start()
}
