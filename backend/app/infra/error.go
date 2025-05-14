package infra

import "errors"

var (
	ErrFatalRedisSet        = errors.New("fatal-redis-set")
	ErrOverMaxContentLength = errors.New("over-max-content-size")
	ErrNoRequestFormData    = errors.New("no-request-form-data")
	ErrFatalReadData        = errors.New("fatal-read-data")
)
