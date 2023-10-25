package usercontrollers

import (
	"net/http"

	contracts "github.com/iagomaia/dload-tech-challenge/internal/domain/contracts/user"
	usecases "github.com/iagomaia/dload-tech-challenge/internal/domain/usecases/user"
	presentation "github.com/iagomaia/dload-tech-challenge/internal/presentation/protocols"
)

type MeController struct {
	UseCase usecases.IGetMe
}

func (c *MeController) Handle(req *presentation.HttpRequest) (*presentation.HttpResponse, error) {
	user, err := c.UseCase.WithCtx(req.Ctx).Get()
	if err != nil {
		return nil, err
	}

	resp := &presentation.HttpResponse{
		Status: http.StatusOK,
		Body: &contracts.MeResponse{
			User: *mapUserModelToObject(user),
		},
	}
	return resp, nil
}
