package file

import (
	"context"
	"net/http"

	"github.com/poteto-go/potathon-backend/infra"
	"github.com/poteto-go/potathon-backend/service/file/usecase"
	"github.com/redis/go-redis/v9"
)

type IFileService interface {
	UploadFile(r *http.Request) error
}

type FileService struct {
	rdb *redis.Client
}

func NewFileService(rdb *redis.Client) IFileService {
	return &FileService{
		rdb: rdb,
	}
}

func (fs *FileService) UploadFile(r *http.Request) error {
	if !usecase.IsUnderMaxLength(r) {
		return infra.ErrOverMaxContentLength
	}

	data, fileName, err := usecase.ReadFileFromRequest(r)
	if err != nil {
		return err
	}

	var ctx = context.Background()
	if err := fs.rdb.Set(ctx, fileName, data, 0).Err(); err != nil {
		return infra.ErrFatalRedisSet
	}

	return nil
}
