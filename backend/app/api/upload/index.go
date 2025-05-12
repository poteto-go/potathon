package upload

import "github.com/poteto-go/poteto"

func NewUploadApi() poteto.Poteto {
	uploadApi := poteto.New()

	uploadApi.Leaf("/upload", func(api poteto.Leaf) {
		api.POST("/input", func(ctx poteto.Context) error {
			return nil
		})
	})

	return uploadApi
}
