package usercontracts

import sharedcontracts "github.com/iagomaia/dload-tech-challenge/internal/domain/contracts/shared"

type SignupRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupResponse sharedcontracts.MessageResponse
