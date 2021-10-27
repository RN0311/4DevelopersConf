package halloweenproduct_handlers

import (
	"net/http"

	halloweenproduct "github.com/RN0311/demo-project/domain/halloweenProduct"
	"github.com/RN0311/demo-project/handlers/representation"
	"github.com/gin-gonic/gin"
)

var GetHalloweenProductById = func(r *halloweenproduct.Repository) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		if len(id) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Id can not be empty!"})
			return
		}

		com, err := r.GetById(id)
		if err != nil {
			panic(err)
		}

		if com == nil {
			c.String(http.StatusNotFound, "")
			return
		}

		representation := representation.HalloweenProduct{
			Id:      com.Id,
			BrandId: com.BrandId,
			Rate:    com.Rate,
		}

		c.JSON(http.StatusOK, representation)
	}
}
