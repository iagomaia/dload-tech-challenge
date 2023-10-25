package usercontracts

type UpdateUserRequest struct {
	Name string `json:"name"`
}

type UpdateUserResponse struct {
	UpdatedUser UserObject `json:"updatedUser"`
}
