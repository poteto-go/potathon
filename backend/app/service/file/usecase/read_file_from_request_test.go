package usecase

import (
	"errors"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey"
	"github.com/poteto-go/potathon-backend/infra"
	"github.com/stretchr/testify/assert"
)

func TestReadFileFromRequest(t *testing.T) {
	t.Run("error case", func(t *testing.T) {
		t.Run("if not found file from formData => ErrNoRequestFormData", func(t *testing.T) {
			patches := gomonkey.NewPatches()
			defer patches.Reset()

			// Arrange
			r := httptest.NewRequest("POST", "/test", nil)

			// Mock
			patches.ApplyMethod(
				reflect.TypeOf(r),
				"FormFile",
				func(_ *http.Request, _ string) (multipart.File, *multipart.FileHeader, error) {
					return nil, nil, errors.New("error")
				},
			)

			// Act
			_, _, err := ReadFileFromRequest(r)

			// Assert
			assert.Error(t, err)
			assert.ErrorIs(t, err, infra.ErrNoRequestFormData)
		})
	})
}
