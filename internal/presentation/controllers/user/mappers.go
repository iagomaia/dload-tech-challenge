package usercontrollers

import (
	usercontracts "github.com/iagomaia/dload-tech-challenge/internal/domain/contracts/user"
	usermodels "github.com/iagomaia/dload-tech-challenge/internal/domain/models/user"
	userusecases "github.com/iagomaia/dload-tech-challenge/internal/domain/usecases/user"
)

func mapUserModelToObject(user *usermodels.User) *usercontracts.UserObject {
	return &usercontracts.UserObject{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func mapUpdateUserContractToDto(payload *usercontracts.UpdateUserRequest) *userusecases.UpdateUserDto {
	return &userusecases.UpdateUserDto{
		Name: payload.Name,
	}
}
