package Router

import (
	priceRouter "GaswapData/route/price"
	"github.com/gin-gonic/gin"
)

func Setup(rt *gin.Engine) {
	// price Router
	priceRouter.GetPriceRouter(rt)
}
