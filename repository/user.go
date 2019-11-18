package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/sumitalp/productcatalog/models"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (us *UserRepository) GetByID(id uint) (*models.User, error) {
	var m models.User
	if err := us.db.First(&m, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (us *UserRepository) GetByEmail(e string) (*models.User, error) {
	var m models.User
	if err := us.db.Where(&models.User{Email: e}).First(&m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (us *UserRepository) GetByUsername(username string) (*models.User, error) {
	var m models.User
	if err := us.db.Where(&models.User{Username: username}).First(&m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (us *UserRepository) Create(u *models.User) (err error) {
	return us.db.Create(u).Error
}

func (us *UserRepository) Update(u *models.User) error {
	return us.db.Model(u).Update(u).Error
}
