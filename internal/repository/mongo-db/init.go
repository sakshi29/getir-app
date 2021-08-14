package mangoDB

import (
	"context"
	"log"

	"github.com/sakshi29/getir-app/common/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init(cfg config.MongoDBConfig) (MongoDBRepo, error) {

	// Set client options
	clientOptions := options.Client().ApplyURI(cfg.Uri)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	log.Println("successfully connected to MongoDB!")

	collection := client.Database(cfg.DatabaseName).Collection(cfg.Collection)

	return Collection{
		records: collection,
	}, nil
}
