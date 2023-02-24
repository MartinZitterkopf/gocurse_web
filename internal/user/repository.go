package user

import (
	"fmt"
	"log"
	"strings"

	"github.com/MartinZitterkopf/gocurse_web/internal/domain"
	"gorm.io/gorm"
)

type (
	Repository interface {
		Create(user *domain.User) error
		GetAll(filters Fillters, limit, offset int) ([]domain.User, error)
		Get(id string) (*domain.User, error)
		Delete(id string) error
		Update(id string, firstName *string, lastName *string, email *string, phone *string) error
		Count(filters Fillters) (int, error)
	}

	repo struct {
		log *log.Logger
		db  *gorm.DB
	}
)

func NewRepo(log *log.Logger, db *gorm.DB) Repository {
	return &repo{
		log: log,
		db:  db,
	}
}

func (repo *repo) Create(user *domain.User) error {

	if err := repo.db.Create(user).Error; err != nil {
		repo.log.Printf("error; %v", err)
		return err
	}
	repo.log.Println("User created with id: ", user.ID)
	return nil
}

func (repo *repo) GetAll(filters Fillters, offset, limit int) ([]domain.User, error) {
	var u []domain.User

	// Model hace referencia al modelo de usuario y Find lo que hace es poblar la informacion que saca de la estructura
	tx := repo.db.Model(&u)
	tx = applyFilters(tx, filters)
	tx = tx.Limit(limit).Offset(offset)

	result := tx.Order("created_at desc").Find(&u)
	if result.Error != nil {
		return nil, result.Error
	}

	return u, nil
}

func (repo *repo) Get(id string) (*domain.User, error) {
	user := domain.User{ID: id}

	if err := repo.db.First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *repo) Delete(id string) error {
	user := domain.User{ID: id}

	if err := repo.db.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}

func (repo *repo) Update(id string, firstName *string, lastName *string, email *string, phone *string) error {
	values := make(map[string]interface{})

	if firstName != nil {
		values["first_name"] = *firstName
	}

	if lastName != nil {
		values["last_name"] = *lastName
	}

	if email != nil {
		values["email"] = *email
	}

	if phone != nil {
		values["phone"] = *phone
	}

	if err := repo.db.Model(&domain.User{}).Where("id = ?", id).Updates(values).Error; err != nil {
		return err
	}

	return nil
}

// recibe la base de datos desde gorm y el filtro a buscar
func applyFilters(tx *gorm.DB, filters Fillters) *gorm.DB {
	if filters.FirstName != "" {
		filters.FirstName = fmt.Sprintf("%%%s%%", strings.ToLower(filters.FirstName))
		tx = tx.Where("lower(first_name) like ?", filters.FirstName)
	}

	if filters.LastName != "" {
		filters.LastName = fmt.Sprintf("%%%s%%", strings.ToLower(filters.LastName))
		tx = tx.Where("lower(first_name) like ?", filters.LastName)
	}

	return tx
}

func (repo *repo) Count(filters Fillters) (int, error) {
	var count int64
	tx := repo.db.Model(domain.User{})
	tx = applyFilters(tx, filters)

	if err := tx.Count(&count).Error; err != nil {
		return 0, err
	}

	return int(count), nil
}
