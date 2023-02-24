package enrollment

import (
	"errors"
	"log"

	"github.com/MartinZitterkopf/gocurse_web/internal/curse"
	"github.com/MartinZitterkopf/gocurse_web/internal/domain"
	"github.com/MartinZitterkopf/gocurse_web/internal/user"
)

type (
	Service interface {
		Create(userID, curseID string) (*domain.Enrollment, error)
	}

	service struct {
		log          *log.Logger
		userService  user.Service
		curseService curse.Service
		repo         Repository
	}
)

func NewService(l *log.Logger, userSvc user.Service, curseSvc curse.Service, r Repository) Service {
	return &service{
		log:          l,
		userService:  userSvc,
		curseService: curseSvc,
		repo:         r,
	}
}

func (s service) Create(userID, curseID string) (*domain.Enrollment, error) {

	enroll := &domain.Enrollment{
		UserID:  userID,
		CurseID: curseID,
		Status:  "P",
	}

	if _, err := s.userService.Get(enroll.UserID); err != nil {
		return nil, errors.New("user id doesn't exists")
	}

	if _, err := s.curseService.GetByID(enroll.CurseID); err != nil {
		return nil, errors.New("curse id doesn't exists")
	}

	if err := s.repo.Create(enroll); err != nil {
		s.log.Printf("error: %v", err)
		return nil, err
	}

	return enroll, nil
}
