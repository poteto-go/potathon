package usecase

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsUnderMaxLength(t *testing.T) {
	t.Run("true case", func(t *testing.T) {
		t.Run("< maxUploadSize", func(t *testing.T) {
			// Arrange
			r := httptest.NewRequest("POST", "/test", nil)
			r.ContentLength = MaxUploadSize - 1

			// Act
			result := IsUnderMaxLength(r)

			// Assert
			assert.True(t, result)
		})

		t.Run("== MaxUploadSize", func(t *testing.T) {
			// Arrange
			r := httptest.NewRequest("POST", "/test", nil)
			r.ContentLength = MaxUploadSize

			// Act
			result := IsUnderMaxLength(r)

			// Assert
			assert.True(t, result)
		})
	})

	t.Run("false case", func(t *testing.T) {
		t.Run("> maxUploadSize", func(t *testing.T) {
			// Arrange
			r := httptest.NewRequest("POST", "/test", nil)
			r.ContentLength = MaxUploadSize + 1

			// Act
			result := IsUnderMaxLength(r)

			// Assert
			assert.False(t, result)
		})
	})
}
