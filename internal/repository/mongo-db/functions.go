package mangoDB

import (
	"context"

	"github.com/sakshi29/getir-app/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

func (c Collection) FetchDocuments(ctx context.Context, filters model.DocumentFilters) ([]model.Record, error) {

	documents := make([]model.Record, 0)

	pipeline := mongo.Pipeline{
		{
			{
				"$project", bson.M{
					"_id":       "$_id",
					"key":       "$key",
					"createdAt": "$createdAt",
					"value":     "$value",
					"totalCount": bson.M{
						"$sum": "$counts",
					},
				},
			},
		},
		{
			{
				"$match", bson.M{
					"$expr": bson.M{
						"$and": []bson.M{
							{"$gte": []interface{}{"$totalCount", filters.MinCount}},
							{"$lte": []interface{}{"$totalCount", filters.MaxCount}},
							{"$gte": []interface{}{"$createdAt", filters.StartDate}},
							{"$lte": []interface{}{"$createdAt", filters.EndDate}},
						},
					},
				},
			},
		},
	}

	cursor, err := c.records.Aggregate(ctx, pipeline)
	if err != nil {
		return documents, err
	}

	if err = cursor.All(ctx, &documents); err != nil {
		return documents, err
	}

	return documents, nil
}
