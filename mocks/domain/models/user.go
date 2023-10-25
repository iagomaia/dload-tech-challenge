package mocks

import (
	"time"

	models "github.com/iagomaia/dload-tech-challenge/internal/domain/models/user"
)

func GetUserModelMock() *models.User {
	return &models.User{
		Id:        "some-id",
		Name:      "some-name",
		Email:     "some-email",
		Password:  "some-password",
		CreatedAt: time.Now(),
		UpdatedAt: nil,
		DeletedAt: nil,
	}
}
