package controller

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey"
	"github.com/poteto-go/potathon-backend/service/file"
	"github.com/poteto-go/poteto"
	"github.com/stretchr/testify/assert"
)

func TestNewUploadController(t *testing.T) {
	// Act
	cont := NewUploadController(nil)

	// Assert
	assert.NotNil(t, cont)
}

func TestUploadController_UploadInput(t *testing.T) {
	// Arrange
	cont := NewUploadController(nil).(*UploadController)

	t.Run("normal case", func(t *testing.T) {
		t.Run("success upload file", func(t *testing.T) {
			patches := gomonkey.NewPatches()
			defer patches.Reset()

			// Arrange
			w := httptest.NewRecorder()
			ctx := poteto.NewContext(w, nil)

			// mock
			patches.ApplyMethod(
				reflect.TypeOf(cont.fileService),
				"UploadFile",
				func(_ *file.FileService, _ *http.Request) error {
					return nil
				},
			)

			// Act
			err := cont.UploadInput(ctx)

			// Assert
			assert.NoError(t, err)
			assert.Equal(t, http.StatusCreated, w.Result().StatusCode)
			assert.Contains(t, w.Body.String(), `"result":"success"`)
		})
	})

	t.Run("error case", func(t *testing.T) {
		t.Run("if fatal upload file => error", func(t *testing.T) {
			patches := gomonkey.NewPatches()
			defer patches.Reset()

			// Arrange
			w := httptest.NewRecorder()
			ctx := poteto.NewContext(w, nil)

			// mock
			patches.ApplyMethod(
				reflect.TypeOf(cont.fileService),
				"UploadFile",
				func(_ *file.FileService, _ *http.Request) error {
					return errors.New("error")
				},
			)

			// Act
			err := cont.UploadInput(ctx)

			// Assert
			assert.Error(t, err)
			assert.ErrorContains(t, err, "error")
		})
	})
}
