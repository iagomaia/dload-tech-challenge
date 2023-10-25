package usercontrollers

import (
	"net/http"

	contracts "github.com/iagomaia/dload-tech-challenge/internal/domain/contracts/user"
	usecases "github.com/iagomaia/dload-tech-challenge/internal/domain/usecases/user"
	presentation "github.com/iagomaia/dload-tech-challenge/internal/presentation/protocols"
)

type LoginController struct {
	UseCase usecases.ILogin
}

func (c *LoginController) Handle(req *presentation.HttpRequest) (*presentation.HttpResponse, error) {
	reqBody := req.Body.(*contracts.LoginRequest)

	dto := &usecases.LoginDto{
		Email:    reqBody.Email,
		Password: reqBody.Password,
	}

	token, err := c.UseCase.WithCtx(req.Ctx).Login(dto)
	if err != nil {
		return nil, err
	}

	resp := &presentation.HttpResponse{
		Status: http.StatusOK,
		Body: &contracts.LoginResponse{
			AccessToken: token,
		},
	}
	return resp, nil
}
