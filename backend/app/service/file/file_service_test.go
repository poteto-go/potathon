package file

import (
	"net/http/httptest"
	"testing"

	"github.com/agiledragon/gomonkey"
	"github.com/poteto-go/potathon-backend/infra"
	"github.com/poteto-go/potathon-backend/service/file/usecase"
	"github.com/stretchr/testify/assert"
)

func TestNewFileService(t *testing.T) {
	// Act
	fs := NewFileService(nil)

	// Assert
	assert.NotNil(t, fs)
}

func TestFileService_ReadFile(t *testing.T) {

	t.Run("error case", func(t *testing.T) {
		t.Run("if over max length", func(t *testing.T) {
			patches := gomonkey.NewPatches()
			defer patches.Reset()

			// Arrange
			fs := NewFileService(nil)
			r := httptest.NewRequest("POST", "/test", nil)
			r.ContentLength = usecase.MaxUploadSize + 1

			// Act
			err := fs.UploadFile(r)

			// Assert
			assert.Error(t, err)
			assert.ErrorIs(t, err, infra.ErrOverMaxContentLength)
		})

		t.Run("if cannot read file from request => error", func(t *testing.T) {
			patches := gomonkey.NewPatches()
			defer patches.Reset()

			// Arrange
			fs := NewFileService(nil)
			r := httptest.NewRequest("POST", "/test", nil)
			r.ContentLength = usecase.MaxUploadSize - 1

			// Act
			err := fs.UploadFile(r)

			// Assert
			assert.Error(t, err)
		})
	})
}
