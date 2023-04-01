package request

type CreateUserRequest struct {
	Email      string `json:"email"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Roles      string `json:"role"`
	DivisionID string `json:"division_id"`
}
