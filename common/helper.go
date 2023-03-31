package common

const (
	DefaultLimit int    = 5
	ImagePath    string = "assets/complaints"
)

type APIResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
