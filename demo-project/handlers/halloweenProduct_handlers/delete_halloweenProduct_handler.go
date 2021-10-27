package halloweenproduct_handlers

import (
	"net/http"

	halloweenproduct "github.com/RN0311/demo-project/domain/halloweenProduct"
	"github.com/gin-gonic/gin"
)

var DeleteHandler = func(r *halloweenproduct.Repository) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		if len(id) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Id can not be empty!"})
			return
		}

		err := r.Delete(id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, "Successfully deleted!")
	}
}
