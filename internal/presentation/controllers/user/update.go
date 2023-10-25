package usercontrollers

import (
	"net/http"

	contracts "github.com/iagomaia/dload-tech-challenge/internal/domain/contracts/user"
	"github.com/iagomaia/dload-tech-challenge/internal/domain/models/utils"
	usecases "github.com/iagomaia/dload-tech-challenge/internal/domain/usecases/user"
	presentation "github.com/iagomaia/dload-tech-challenge/internal/presentation/protocols"
)

type UpdateUserController struct {
	UseCase usecases.IUpdateUser
}

func (c *UpdateUserController) Handle(req *presentation.HttpRequest) (*presentation.HttpResponse, error) {
	body := req.Body.(*contracts.UpdateUserRequest)
	paramID := req.Params["userId"]
	userId := req.Ctx.Value(utils.UserIDContextKey).(string)

	if paramID != userId {
		cErr := utils.CustomError{
			Status:  http.StatusForbidden,
			Message: "You don't have permission to change this user",
		}
		return nil, cErr
	}

	dto := mapUpdateUserContractToDto(body)
	dto.ID = userId

	user, err := c.UseCase.WithCtx(req.Ctx).Update(dto)
	if err != nil {
		return nil, err
	}

	resp := &presentation.HttpResponse{
		Status: http.StatusOK,
		Body: &contracts.UpdateUserResponse{
			UpdatedUser: *mapUserModelToObject(user),
		},
	}
	return resp, nil
}
