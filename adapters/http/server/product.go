package server

import (
	"encoding/json"
	"github.com/chjoaquim/go-hexagonal-arch/adapters/http/server/domain"
	"github.com/chjoaquim/go-hexagonal-arch/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type errorMessage struct {
	Message string `json:"message"`
}

func RunService(service application.ProductServiceInterface) {
	r := gin.Default()

	r.GET("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		product, err := service.Get(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		if product.GetID() == "" {
			c.JSON(http.StatusNotFound, errorMessage{
				Message: "not_found",
			})
			return
		}

		c.JSON(http.StatusOK, product)
	})

	r.POST("/products", func(c *gin.Context) {
		var product server.CreateRequest
		decoder := json.NewDecoder(c.Request.Body)
		err := decoder.Decode(&product)
		if err != nil {
			c.JSON(http.StatusBadRequest, errorMessage{
				Message: err.Error(),
			})
			return
		}

		res, err := service.Create(product.Name, product.Price)
		if err != nil {
			c.JSON(http.StatusInternalServerError, errorMessage{
				Message: err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, res)
	})

	if err := r.Run(":9000"); err != nil {
		panic(err)
	}
}
