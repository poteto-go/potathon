package upload

import (
	"github.com/poteto-go/potathon-backend/api/upload/controller"
	"github.com/poteto-go/poteto"
	"github.com/redis/go-redis/v9"
)

func NewUploadApi(rdb *redis.Client) poteto.Poteto {
	uploadApi := poteto.New()

	uploadApi.Leaf("/upload", func(api poteto.Leaf) {
		uploadController := controller.NewUploadController(rdb)

		api.POST("/input", uploadController.UploadInput)
	})

	return uploadApi
}
