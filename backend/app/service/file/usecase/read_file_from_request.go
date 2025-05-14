package usecase

import (
	"bytes"
	"io"
	"net/http"

	"github.com/poteto-go/potathon-backend/infra"
)

func ReadFileFromRequest(r *http.Request) ([]byte, string, error) {
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		return []byte(""), "", infra.ErrNoRequestFormData
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return nil, "", infra.ErrFatalReadData
	}

	return buf.Bytes(), fileHeader.Filename, nil
}
