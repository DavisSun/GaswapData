package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PingHandler(context *gin.Context) {

	context.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
