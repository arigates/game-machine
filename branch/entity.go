package branch

import (
	"game-machine/helper"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Branch struct {
	ID      string `gorm:"type:uuid;primary_key;"`
	Code    string `gorm:"size:20"`
	Name    string `gorm:"size:100"`
	Address string
	helper.Timestamp
	helper.SoftDeletes
}

// BeforeCreate will set a UUID rather than numeric ID.
func (branch *Branch) BeforeCreate(tx *gorm.DB) error {
	branch.ID = uuid.NewV4().String()

	return nil
}
