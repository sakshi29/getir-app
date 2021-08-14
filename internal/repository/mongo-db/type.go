package mangoDB

import "go.mongodb.org/mongo-driver/mongo"

type Collection struct {
	records *mongo.Collection
}
