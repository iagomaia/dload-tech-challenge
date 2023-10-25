package usercontrollers

import (
	"net/http"

	contracts "github.com/iagomaia/dload-tech-challenge/internal/domain/contracts/user"
	usecases "github.com/iagomaia/dload-tech-challenge/internal/domain/usecases/user"
	presentation "github.com/iagomaia/dload-tech-challenge/internal/presentation/protocols"
)

type CreateUserController struct {
	UseCase usecases.ISignupUser
}

func (c *CreateUserController) Handle(req *presentation.HttpRequest) (*presentation.HttpResponse, error) {
	reqBody := req.Body.(*contracts.SignupRequest)

	dto := &usecases.SignupUserDto{
		Name:     reqBody.Name,
		Email:    reqBody.Email,
		Password: reqBody.Password,
	}

	_, err := c.UseCase.WithCtx(req.Ctx).Signup(dto)
	if err != nil {
		return nil, err
	}

	resp := &presentation.HttpResponse{
		Status: http.StatusCreated,
		Body: &contracts.SignupResponse{
			Message: "User created successfuly",
		},
	}
	return resp, nil
}
