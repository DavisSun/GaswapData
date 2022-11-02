package Price

import (
	controllers "GaswapData/Controllers"
	"github.com/gin-gonic/gin"
)

func GetPriceRouter(rt *gin.Engine) {

	priceRouter := rt.Group("/price")

	priceRouter.GET("/ping", controllers.PingHandler)

}
