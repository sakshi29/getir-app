package delivery

import (
	"github.com/sakshi29/getir-app/common/config"
	"github.com/sakshi29/getir-app/internal/usecase"
)

type (
	API struct {
		cfg       *config.Config
		useCaseV1 usecase.UseCaseV1
	}

	Response struct {
		Message string      `json:"msg"`
		Code    int         `json:"code"`
		Data    interface{} `json:"records"`
	}
)
