package Router

import (
	priceRouter "GaswapData/Route/Price"
	"github.com/gin-gonic/gin"
)

func Setup(rt *gin.Engine) {
	// Price Router
	priceRouter.GetPriceRouter(rt)
}
