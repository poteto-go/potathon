package main

import (
	"net/http"

	"github.com/poteto-go/poteto"
)

func main() {
	p := poteto.New()

	p.GET("/hello", func(ctx poteto.Context) error {
		return ctx.JSON(http.StatusOK, map[string]string{
			"hello": "poteto",
		})
	})

	p.Run("3000")
}
