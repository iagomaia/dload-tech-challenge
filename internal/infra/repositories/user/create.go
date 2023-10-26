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
	_       userservicesprotocols.ICreateUserRepository = (*CreateUserRepository)(nil)
	indexes                                             = []mongo.IndexModel{
		{
			Keys:    bson.M{"email": 1},
			Options: options.Index().SetUnique(true),
		},
	}
)

type CreateUserRepository struct {
	collection *mongo.Collection
	session    mongo.Session
	ctx        context.Context
}

func (r *CreateUserRepository) Init() {
	session, collection, err := factories.GetMongoClient().GetCollection(UserCollection)
	if err != nil {
		log.Fatalf("Error connection to DB: %v\n", err)
	}
	defer session.EndSession(context.Background())
	_, err = collection.Indexes().CreateMany(context.Background(), indexes)
	if err != nil {
		log.Fatalf("Error creating indexes: %v\n", err)
	}
	r.session = session
	r.collection = collection
}

func (r *CreateUserRepository) WithCtx(ctx context.Context) userservicesprotocols.ICreateUserRepository {
	return &CreateUserRepository{
		collection: r.collection,
		session:    r.session,
		ctx:        ctx,
	}
}

func (r *CreateUserRepository) Create(dto *userservicesprotocols.CreateUserDto) (*usermodels.User, error) {
	defer r.session.EndSession(r.ctx)
	dbe := mapCreateUserDtoToDbe(dto)
	result, err := r.collection.InsertOne(r.ctx, dbe)
	if err != nil {
		cErr := utils.CustomError{
			Status:        http.StatusInternalServerError,
			Message:       "Error inserting user into DB",
			OriginalError: err,
		}
		if mongo.IsDuplicateKeyError(err) {
			cErr = utils.CustomError{
				Status:        http.StatusConflict,
				Message:       "Email already in use",
				OriginalError: err,
			}
		}
		return nil, cErr
	}
	id, _ := result.InsertedID.(primitive.ObjectID)
	dbe.Id = &id

	return mapDbeToModel(dbe), nil
}
