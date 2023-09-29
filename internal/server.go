package internal

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	*http.Server
	router *gin.Engine
	service *OrderService
	port int
}

func NewServer(service *OrderService, port int) *Server {
	router := gin.New()
	router.Use(gin.Recovery())

	router.POST("/orders", func(c *gin.Context) {
		var order Order
		if err := c.ShouldBindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := service.SaveOrder(c.Request.Context(), &order); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, order)
	})

	router.GET("/orders", func(c *gin.Context) {
		orders, err := service.GetOrders(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, orders)
	})

	router.GET("/orders/:id", func(c *gin.Context) {
		order, err := service.GetOrder(c.Request.Context(), c.Param("id"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, order)
	})

	return &Server{
		Server: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: router,
		},
		router: router,
		service: service,
		port: port,
	}
}

func (s *Server) Run() error {
	return s.Server.ListenAndServe()
}