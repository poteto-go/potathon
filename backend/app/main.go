package main

import (
	"net/http"

	"github.com/poteto-go/potathon-backend/api/upload"
	"github.com/poteto-go/poteto"
	"github.com/redis/go-redis/v9"
)

func main() {
	p := poteto.New()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "kvs:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
		PoolSize: 1000,
	})

	p.GET("/hello", func(ctx poteto.Context) error {
		return ctx.JSON(http.StatusOK, map[string]string{
			"hello": "poteto",
		})
	})

	p.AddApi(upload.NewUploadApi(rdb))

	p.Run("3000")
}
