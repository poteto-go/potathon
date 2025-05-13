package usecase

import "net/http"

const MaxUploadSize = 1024 * 1024 // 10MB

// check if req content is under max length
func IsUnderMaxLength(r *http.Request) bool {
	return r.ContentLength <= MaxUploadSize
}
