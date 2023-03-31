package response

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Complaint struct {
	ID         uuid.UUID      `json:"id"`
	DivisionID uuid.UUID      `json:"division_id"`
	LocationID uuid.UUID      `json:"location_id"`
	TicketID   string         `json:"ticket_id"`
	Date       datatypes.Date `json:"date"`
	Complaint  string         `json:"complaint"`
	Image      string         `json:"image"`
	SupportID  uuid.UUID      `json:"support_id"`
	EngineerID *uuid.UUID     `json:"engineer_id"`
	AccessorID *uuid.UUID     `json:"accessor_id"`
}

type APIComplaintManager struct {
	Complaint
	Division *ComplainWithDivisionScheme `json:"division"`
	Location *ComplainWithLocationScheme `json:"location"`
}

type APIComplaintEngineer struct {
	Complaint
	Division *ComplainWithDivisionScheme `gorm:"foreignKey:DivisionID" json:"division"`
	Location *ComplainWithLocationScheme `gorm:"foreignKey:LocationID" json:"location"`
}

type ComplainWithDivisionScheme struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type ComplainWithLocationScheme struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
