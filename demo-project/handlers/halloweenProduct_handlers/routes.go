package halloweenproduct_handlers

import (
	halloweenproduct "github.com/RN0311/demo-project/domain/halloweenProduct"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(engine *gin.Engine, repo *halloweenproduct.Repository) {
	engine.GET("halloweenproduct/id/:id", GetHalloweenProductById(repo))
	engine.POST("halloweenproduct", PostHalloweenProductHandler(repo))
	engine.DELETE("halloweenproduct/id/:id", DeleteHandler(repo))
}
