package repositories

import (
	"time"

	usermodels "github.com/iagomaia/dload-tech-challenge/internal/domain/models/user"
	userservicesprotocols "github.com/iagomaia/dload-tech-challenge/internal/services/protocols/user"
	"go.mongodb.org/mongo-driver/bson"
)

func mapCreateUserDtoToDbe(dto *userservicesprotocols.CreateUserDto) *UserDbe {
	return &UserDbe{
		Id:        nil,
		Name:      dto.Name,
		Email:     dto.Email,
		Password:  dto.Password,
		CreatedAt: time.Now(),
		UpdatedAt: nil,
		DeletedAt: nil,
	}
}

func mapDbeToModel(dbe *UserDbe) *usermodels.User {
	return &usermodels.User{
		Id:        dbe.Id.Hex(),
		Name:      dbe.Name,
		Email:     dbe.Email,
		Password:  dbe.Password,
		CreatedAt: dbe.CreatedAt,
		UpdatedAt: dbe.UpdatedAt,
		DeletedAt: dbe.DeletedAt,
	}
}

func mapUpdateUserDtoToPayload(dto *userservicesprotocols.UpdateUserDto) bson.M {
	now := time.Now()
	return bson.M{
		"$set": &UpdateUserDBPayload{
			Name:      dto.Name,
			UpdatedAt: &now,
		},
	}
}
