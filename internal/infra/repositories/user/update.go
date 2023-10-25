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
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	_ userservicesprotocols.IUpdateUserRepository = (*UpdateUserRepository)(nil)
)

type UpdateUserRepository struct {
	collection *mongo.Collection
	session    mongo.Session
	ctx        context.Context
}

func (r *UpdateUserRepository) Init() {
	session, collection, err := factories.GetMongoClient().GetCollection(UserCollection)
	if err != nil {
		log.Fatalf("Error connection to DB: %v\n", err)
	}
	defer session.EndSession(context.Background())
	r.session = session
	r.collection = collection
}

func (r *UpdateUserRepository) WithCtx(ctx context.Context) userservicesprotocols.IUpdateUserRepository {
	return &UpdateUserRepository{
		collection: r.collection,
		session:    r.session,
		ctx:        ctx,
	}
}

func (r *UpdateUserRepository) Update(dto *userservicesprotocols.UpdateUserDto) (*usermodels.User, error) {
	defer r.session.EndSession(r.ctx)
	cErr := utils.CustomError{
		Status:  http.StatusInternalServerError,
		Message: "Error updating user by ID",
	}

	oid, err := primitive.ObjectIDFromHex(dto.ID)
	if err != nil {
		cErr.OriginalError = err
		return nil, cErr
	}

	payload := mapUpdateUserDtoToPayload(dto)
	query := bson.M{
		"_id": oid,
	}

	returnDoc := options.After
	options := &options.FindOneAndUpdateOptions{
		ReturnDocument: &returnDoc,
	}
	result := r.collection.FindOneAndUpdate(r.ctx, query, payload, options)

	if result.Err() != nil {
		cErr.OriginalError = result.Err()
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
