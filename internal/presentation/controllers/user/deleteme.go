package usercontrollers

import (
	"net/http"

	usecases "github.com/iagomaia/dload-tech-challenge/internal/domain/usecases/user"
	presentation "github.com/iagomaia/dload-tech-challenge/internal/presentation/protocols"
)

type DeleteMeController struct {
	UseCase usecases.IDeleteMe
}

func (c *DeleteMeController) Handle(req *presentation.HttpRequest) (*presentation.HttpResponse, error) {
	err := c.UseCase.WithCtx(req.Ctx).Delete()
	if err != nil {
		return nil, err
	}

	resp := &presentation.HttpResponse{
		Status: http.StatusNoContent,
	}
	return resp, nil
}
