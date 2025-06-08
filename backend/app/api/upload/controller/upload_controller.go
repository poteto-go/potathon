package controller

import (
	"net/http"
	"sync"

	"github.com/poteto-go/potathon-backend/service/file"
	"github.com/poteto-go/poteto"
	"github.com/redis/go-redis/v9"
)

type IUploadController interface {
	UploadInput(ctx poteto.Context) error
}

type UploadController struct {
	lock        sync.Mutex
	fileService file.IFileService
}

func NewUploadController(rdb *redis.Client) IUploadController {
	return &UploadController{
		fileService: file.NewFileService(rdb),
	}
}

func (uc *UploadController) UploadInput(ctx poteto.Context) error {
	uc.lock.Lock()
	defer uc.lock.Unlock()

	if err := uc.fileService.UploadFile(ctx.GetRequest()); err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, map[string]string{
		"result": "success",
	})
}
