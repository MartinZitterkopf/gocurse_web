package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Enrollment struct {
	ID string `json:"id" gorm:"type:char(36);not null;primary_key;unique_index"`
	// UserID va a hacer referencia a la tabla user
	UserID string `json:"user_id,omitempty" gorm:"type:char(36)"`
	User   *User  `json:"user,omitempty"`
	// CurseID va a hacer referencia a la tabla curse
	CurseID   string     `json:"curse_id,omitempty" gorm:"type:char(36)"`
	Curse     *Curse     `json:"curse,omitempty"`
	Status    string     `json:"status" gorm:"type:char(2)"`
	CreatedAt *time.Time `json:"-"`
	UpdateAt  *time.Time `json:"-"`
}

func (e *Enrollment) BeforeCreate(tx *gorm.DB) (err error) {

	if e.ID == "" {
		e.ID = uuid.New().String()
	}
	return
}
