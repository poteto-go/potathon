package usecase

import (
	"net/http"

	"github.com/poteto-go/potathon-backend/infra"
)

func ReadFileFromRequest(r *http.Request) ([]byte, error) {
	file, _, err := r.FormFile("file")
	if err != nil {
		return []byte(""), infra.ErrNoRequestFormData
	}
	defer file.Close()

	var data []byte
	if _, err := file.Read(data); err != nil {
		return []byte(""), infra.ErrFatalReadData
	}

	return data, nil
}
