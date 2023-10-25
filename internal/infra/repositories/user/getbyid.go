package repositories

import (
	"context"
	"log"
	"net/http"

	usermodels "github.com/iagomaia/dload-tech-challenge/internal/domain/models/user"
	"github.com/iagomaia/dload-tech-challenge/internal/domain/models/utils"
	factories "github.com/iagomaia/dload-tech-challenge/internal/factories/clients"
	userservicesprotocols "github.com/iagomaia/dload-tech-challenge/internal/services/protocols/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	_ userservicesprotocols.IGetUserByIDRepository = (*GetUserByIDRepository)(nil)
)

type GetUserByIDRepository struct {
	collection *mongo.Collection
	session    mongo.Session
	ctx        context.Context
}

func (r *GetUserByIDRepository) Init() {
	session, collection, err := factories.GetMongoClient().GetCollection(UserCollection)
	if err != nil {
		log.Fatalf("Error connection to DB: %v\n", err)
	}
	defer session.EndSession(context.Background())
	r.session = session
	r.collection = collection
}

func (r *GetUserByIDRepository) WithCtx(ctx context.Context) userservicesprotocols.IGetUserByIDRepository {
	return &GetUserByIDRepository{
		collection: r.collection,
		session:    r.session,
		ctx:        ctx,
	}
}

func (r *GetUserByIDRepository) Get(id string) (*usermodels.User, error) {
	cErr := utils.CustomError{
		Status:  http.StatusInternalServerError,
		Message: "Error finding user by ID",
	}
	defer r.session.EndSession(r.ctx)
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		cErr.OriginalError = err
		return nil, cErr
	}

	query := bson.M{
		"_id": oid,
	}
	result := r.collection.FindOne(r.ctx, query)
	if result.Err() != nil {
		return nil, cErr
	}
	userDbe := &UserDbe{}
	err = result.Decode(userDbe)
	if err != nil {
		cErr.OriginalError = err
		return nil, cErr
	}

	return mapDbeToModel(userDbe), nil
}
