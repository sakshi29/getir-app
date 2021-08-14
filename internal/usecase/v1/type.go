package v1

import (
	"github.com/sakshi29/getir-app/common/config"
	mongoDB "github.com/sakshi29/getir-app/internal/repository/mongo-db"
)

type useCaseV1 struct {
	inMemoryStore map[string]string
	cfg           *config.Config
	mongoDB       mongoDB.MongoDBRepo
}
