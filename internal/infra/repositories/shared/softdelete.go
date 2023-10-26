package repositoriesshared

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func GetSoftDeletePayload() bson.M {
	return bson.M{
		"$set": bson.M{
			"deletedAt": time.Now(),
		},
	}
}
