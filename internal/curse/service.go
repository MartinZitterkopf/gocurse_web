package curse

import (
	"log"
	"time"

	"github.com/MartinZitterkopf/gocurse_web/internal/domain"
)

type (
	Service interface {
		Create(name, startDate, endDate string) (*domain.Curse, error)
		GetAll(filters Fillters, offset, limit int) ([]domain.Curse, error)
		GetByID(id string) (*domain.Curse, error)
		Update(id string, name, startDate, endDate *string) error
		Delete(id string) error
		Count(filters Fillters) (int, error)
	}

	service struct {
		log  *log.Logger
		repo Repository
	}

	Fillters struct {
		Name string
	}
)

func NewService(l *log.Logger, r Repository) Service {
	return &service{
		log:  l,
		repo: r,
	}
}

func (s service) Create(name, startDate, endDate string) (*domain.Curse, error) {

	startDateParsed, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		s.log.Println(err)
		return nil, err
	}

	endDateParsed, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		s.log.Println(err)
		return nil, err
	}

	curse := &domain.Curse{
		Name:      name,
		StartDate: startDateParsed,
		EndDate:   endDateParsed,
	}

	if err := s.repo.Create(curse); err != nil {
		s.log.Println(err)
		return nil, err
	}

	return curse, nil
}

func (s service) GetAll(filters Fillters, offset, limit int) ([]domain.Curse, error) {
	curses, err := s.repo.GetAll(filters, offset, limit)
	if err != nil {
		return nil, err
	}

	return curses, nil
}

func (s service) GetByID(id string) (*domain.Curse, error) {
	curse, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return curse, nil
}

func (s service) Update(id string, name, startDate, endDate *string) error {
	var startDateParsed, endDateParsed *time.Time

	if startDate != nil {
		date, err := time.Parse("2006-01-02", *startDate)
		if err != nil {
			s.log.Println(err)
			return err
		}
		startDateParsed = &date
	}

	if endDate != nil {
		date, err := time.Parse("2006-01-02", *endDate)
		if err != nil {
			s.log.Println(err)
			return err
		}
		endDateParsed = &date
	}

	return s.repo.Update(id, name, startDateParsed, endDateParsed)
}

func (s service) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s service) Count(filters Fillters) (int, error) {
	return s.repo.Count(filters)
}
