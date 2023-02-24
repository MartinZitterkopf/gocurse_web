package enrollment

import (
	"log"

	"github.com/MartinZitterkopf/gocurse_web/internal/domain"
	"gorm.io/gorm"
)

type (
	Repository interface {
		Create(enroll *domain.Enrollment) error
	}

	repo struct {
		log *log.Logger
		db  *gorm.DB
	}
)

func NewRepo(l *log.Logger, db *gorm.DB) Repository {
	return &repo{
		log: l,
		db:  db,
	}
}

func (r *repo) Create(enroll *domain.Enrollment) error {

	if err := r.db.Create(enroll).Error; err != nil {
		r.log.Printf("error: %v", err)
		return err
	}

	r.log.Println("enrollment created with id: ", enroll.ID)
	return nil
}
