package scheme

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	Email     string         `gorm:"index:idx_email,unique;type:varchar(255);" json:"email"`
	Username  string         `gorm:"index:idx_username,unique;type:varchar(255);not null" json:"username"`
	Password  *string        `gorm:"type:text" json:"password"`
	Roles     datatypes.JSON `gorm:"type:longtext;not null" json:"roles"`
	CreatedAt time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;" json:"deleted_at"`
}

type Division struct {
	ID        uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	Name      string         `gorm:"column:name;type:varchar(255);not null" json:"name"`
	CreatedAt time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;" json:"deleted_at"`
}

type Location struct {
	ID        uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	Name      string         `gorm:"column:name;type:varchar(255);not null" json:"name"`
	CreatedAt time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;" json:"deleted_at"`
}

type Complaint struct {
	ID         uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;primaryKey;" json:"id"`
	DivisionID uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_division_id;not null;" json:"division_id"`
	LocationID uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_location_id;not null;" json:"location_id"`
	TicketID   string         `gorm:"type:varchar(255);unique;not null;" json:"ticket_id"`
	Date       datatypes.Date `gorm:"type:date;not null;" json:"date"`
	Complaint  string         `gorm:"type:text;not null;" json:"complaint"`
	Image      string         `gorm:"type:text;" json:"image"`
	SupportID  uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_support_id;not null;" json:"support_id"`
	EngineerID uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_engineer_id;" json:"engineer_id"`
	AccessorID uuid.UUID      `gorm:"type:char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;index:idx_accessor_id;" json:"accessor_id"`
	Estimate   uint           `gorm:"type:int(11);not null;default:0;" json:"estimate"`
	Status     uint           `gorm:"column:status;type:smallint(6);not null;default:0;" json:"status"`
	CreatedAt  time.Time      `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at;not null" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;" json:"deleted_at"`
	Division   Division       `gorm:"foreignKey:DivisionID"`
	Location   Location       `gorm:"foreignKey:LocationID"`
	Support    User           `gorm:"foreignKey:SupportID"`
	Engineer   *User          `gorm:"foreignKey:EngineerID"`
	Accessor   *User          `gorm:"foreignKey:AccessorID"`
}
