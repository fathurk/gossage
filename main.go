package main

import (
	"gossage/handlers"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	s := &http.Server{
		Addr:           ":3030",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			user := v1.Group("/user")
			{
				user.POST("/create", handlers.CreateUser)
				user.GET("/:email", handlers.Finduser)
			}

			contact := v1.Group("/contact")
			{
				contact.POST("/add", handlers.AddContact)
			}
		}
	}
	s.ListenAndServe()
}
