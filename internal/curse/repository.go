package curse

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/MartinZitterkopf/gocurse_web/internal/domain"
	"gorm.io/gorm"
)

type (
	Repository interface {
		Create(curse *domain.Curse) error
		GetAll(filters Fillters, limit, offset int) ([]domain.Curse, error)
		GetByID(id string) (*domain.Curse, error)
		Update(id string, name *string, startDate, endDate *time.Time) error
		Delete(id string) error
		Count(filters Fillters) (int, error)
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

func (r *repo) Create(curse *domain.Curse) error {

	if err := r.db.Create(curse).Error; err != nil {
		r.log.Printf("error; %v", err)
		return err
	}
	r.log.Println("Curse created with id: ", curse.ID)
	return nil
}

func (repo *repo) GetAll(filters Fillters, offset, limit int) ([]domain.Curse, error) {
	var c []domain.Curse

	// Model hace referencia al modelo de usuario y Find lo que hace es poblar la informacion que saca de la estructura
	tx := repo.db.Model(&c)
	tx = applyFilters(tx, filters)
	tx = tx.Limit(limit).Offset(offset)

	result := tx.Order("created_at desc").Find(&c)
	if result.Error != nil {
		return nil, result.Error
	}

	return c, nil
}

func (repo *repo) GetByID(id string) (*domain.Curse, error) {
	curse := domain.Curse{ID: id}

	if err := repo.db.First(&curse).Error; err != nil {
		return nil, err
	}

	return &curse, nil
}

func (repo *repo) Update(id string, name *string, startDate, endDate *time.Time) error {
	values := make(map[string]interface{})

	if name != nil {
		values["name"] = *name
	}

	if startDate != nil {
		values["start_date"] = *startDate
	}

	if endDate != nil {
		values["end_date"] = *endDate
	}

	if err := repo.db.Model(&domain.Curse{}).Where("id = ?", id).Updates(values).Error; err != nil {
		return err
	}

	return nil
}

func (repo *repo) Delete(id string) error {
	curse := domain.Curse{ID: id}

	if err := repo.db.Delete(&curse).Error; err != nil {
		return err
	}

	return nil
}

func (repo *repo) Count(filters Fillters) (int, error) {
	var count int64
	tx := repo.db.Model(domain.Curse{})
	tx = applyFilters(tx, filters)

	if err := tx.Count(&count).Error; err != nil {
		return 0, err
	}

	return int(count), nil
}

func applyFilters(tx *gorm.DB, filters Fillters) *gorm.DB {
	if filters.Name != "" {
		filters.Name = fmt.Sprintf("%%%s%%", strings.ToLower(filters.Name))
		tx = tx.Where("lower(name) like ?", filters.Name)
	}

	return tx
}
