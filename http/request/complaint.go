package request

import "mime/multipart"

type SendComplaintRequest struct {
	DivisionID string                `form:"division_id"`
	LocationID string                `form:"location_id"`
	Complaint  string                `form:"complaint"`
	Image      *multipart.FileHeader `form:"image"`
}
