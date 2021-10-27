package halloweenproduct_handlers

import (
	"net/http"

	halloweenproduct "github.com/RN0311/demo-project/domain/halloweenProduct"
	"github.com/RN0311/demo-project/handlers/request_models"
	"github.com/gin-gonic/gin"
)

var PostHalloweenProductHandler = func(r *halloweenproduct.Repository) func(c *gin.Context) {
	return func(c *gin.Context) {
		var rm *request_models.CreateHalloweenProductModel
		err := c.ShouldBindJSON(&rm)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": " Bad Request"})
			return
		}

		com := &halloweenproduct.HalloweenProduct{
			Id:      rm.Id,
			BrandId: rm.BrandId,
			Rate:    rm.Rate,
		}

		err = r.Insert(com)

		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusAccepted, "Your halloween product is on it's way!")
	}
}
