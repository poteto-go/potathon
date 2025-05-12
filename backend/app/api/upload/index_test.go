package upload

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUploadApi(t *testing.T) {
	uploadApi := NewUploadApi(nil)

	t.Run("assert find route", func(t *testing.T) {
		t.Run("[POST] /upload/input", func(t *testing.T) {
			// Act
			found := uploadApi.Check(http.MethodPost, "/upload/input")

			// Assert
			assert.True(t, found)
		})
	})
}
