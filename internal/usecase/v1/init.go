package v1

import (
	"github.com/sakshi29/getir-app/common/config"
	mongoDB "github.com/sakshi29/getir-app/internal/repository/mongo-db"
	"github.com/sakshi29/getir-app/internal/usecase"
)

func InitUseCaseV1(cfg *config.Config, mongoDBRepo mongoDB.MongoDBRepo) usecase.UseCaseV1 {

	return useCaseV1{
		inMemoryStore: make(map[string]string),
		cfg:           cfg,
		mongoDB:       mongoDBRepo,
	}
}
