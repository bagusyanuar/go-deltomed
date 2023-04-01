package manager

import (
	"github.com/bagusyanuar/go-deltomed/model"
)

type ProfileRepository interface {
	GetMyProfile(authorizedID string) (data *model.User, err error)
}
