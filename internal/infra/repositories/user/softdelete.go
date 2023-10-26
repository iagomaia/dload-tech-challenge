package repositories

import (
	"context"
	"log"
	"net/http"

	"github.com/iagomaia/dload-tech-challenge/internal/domain/models/utils"
	factories "github.com/iagomaia/dload-tech-challenge/internal/factories/clients"
	repositoriesshared "github.com/iagomaia/dload-tech-challenge/internal/infra/repositories/shared"
	userservicesprotocols "github.com/iagomaia/dload-tech-challenge/internal/services/protocols/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	_ userservicesprotocols.ISoftDeleteUserRepository = (*SoftDeleteUserRepository)(nil)
)

type SoftDeleteUserRepository struct {
	collection *mongo.Collection
	session    mongo.Session
	ctx        context.Context
}

func (r *SoftDeleteUserRepository) Init() {
	session, collection, err := factories.GetMongoClient().GetCollection(UserCollection)
	if err != nil {
		log.Fatalf("Error connection to DB: %v\n", err)
	}
	defer session.EndSession(context.Background())
	r.session = session
	r.collection = collection
}

func (r *SoftDeleteUserRepository) WithCtx(ctx context.Context) userservicesprotocols.ISoftDeleteUserRepository {
	return &SoftDeleteUserRepository{
		collection: r.collection,
		session:    r.session,
		ctx:        ctx,
	}
}

func (r *SoftDeleteUserRepository) SoftDelete(id string) error {
	defer r.session.EndSession(r.ctx)
	cErr := utils.CustomError{
		Status:  http.StatusInternalServerError,
		Message: "Error deleting user by ID",
	}

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		cErr.OriginalError = err
		return cErr
	}

	payload := repositoriesshared.GetSoftDeletePayload()
	result, err := r.collection.UpdateByID(r.ctx, oid, payload)

	if err != nil {
		cErr.OriginalError = err
		return cErr
	}

	if result.MatchedCount == 1 {
		cErr = utils.CustomError{
			Status:  http.StatusNotFound,
			Message: "User not found",
		}
		return cErr
	}
	return nil
}
