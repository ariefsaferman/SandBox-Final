package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func logMiddleware(text string) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := "123455"
		received := time.Now()
		log.Println("Request received", received, requestId)
		c.Set("requestId", requestId)
		c.Next()
		log.Println("Response sent", time.Since(received))
	}
}

func errorMiddleware(c *gin.Context) {
	hasError := len(c.Errors) > 0
	if hasError {
		log.Println("something went wrong", c.Errors)
		firstErr := c.Errors[0]
		c.JSON(http.StatusNotFound, gin.H{
			"message": firstErr.Error(),
			"code":    "NOT_FOUND",
		})
	}
}

func main() {
	r := gin.Default()
	handler := func(c *gin.Context) {
		id := c.MustGet("requestId")

		log.Println("handler called", id)
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	}

	handlerErrorId := func(c *gin.Context) {
		id := c.Param("id")
		if id == "1" {
			errNotFound := errors.New("user not found")
			c.Error(errNotFound)
			c.Error(errors.New("ouch"))
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"name": "Dewa",
		})
	}

	v1 := r.Group("/v1")
	{
		v1.GET("/users", handlerErrorId, handler)

		v1.GET("/users/:id", logMiddleware("1"), handlerErrorId, errorMiddleware)
	}
	r.GET("/ping", logMiddleware("1"), handler)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
