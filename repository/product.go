package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/sumitalp/productcatalog/models"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (as *ProductRepository) GetBySlug(s string) (*models.Product, error) {
	var m models.Product
	err := as.db.Where(&models.Product{Slug: s}).Preload("Categories").Preload("Owner").Find(&m).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, err
}

func (as *ProductRepository) GetUserProductBySlug(userID uint, slug string) (*models.Product, error) {
	var m models.Product
	err := as.db.Where(&models.Product{Slug: slug, OwnerID: userID}).Find(&m).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, err
}

func (as *ProductRepository) CreateProduct(a *models.Product) error {
	categories := a.Categories
	tx := as.db.Begin()
	if err := tx.Create(&a).Error; err != nil {
		return err
	}
	for _, t := range a.Categories {
		err := tx.Where(&models.Category{Category: t.Category}).First(&t).Error
		if err != nil && !gorm.IsRecordNotFoundError(err) {
			tx.Rollback()
			return err
		}
		if err := tx.Model(&a).Association("Categories").Append(t).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	if err := tx.Where(a.ID).Preload("Categories").Preload("Owner").Find(&a).Error; err != nil {
		tx.Rollback()
		return err
	}
	a.Categories = categories
	return tx.Commit().Error
}

func (as *ProductRepository) UpdateProduct(a *models.Product, categoryList []string) error {
	tx := as.db.Begin()
	if err := tx.Model(a).Update(a).Error; err != nil {
		return err
	}
	categories := make([]models.Category, 0)
	for _, t := range categoryList {
		category := models.Category{Category: t}
		err := tx.Where(&category).First(&category).Error
		if err != nil && !gorm.IsRecordNotFoundError(err) {
			tx.Rollback()
			return err
		}
		categories = append(categories, category)
	}
	if err := tx.Model(a).Association("Categories").Replace(categories).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Where(a.ID).Preload("Categories").Preload("Owner").Find(a).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (as *ProductRepository) DeleteProduct(a *models.Product) error {
	return as.db.Delete(a).Error
}

func (as *ProductRepository) List(offset, limit int) ([]models.Product, int, error) {
	var (
		products []models.Product
		count    int
	)
	as.db.Model(&products).Count(&count)
	as.db.Preload("Categories").Preload("Owner").Offset(offset).Limit(limit).Order("created_at desc").Find(&products)
	return products, count, nil
}

func (as *ProductRepository) ListByCategory(category string, offset, limit int) ([]models.Product, int, error) {
	var (
		t        models.Category
		products []models.Product
		count    int
	)
	err := as.db.Where(&models.Category{Category: category}).First(&t).Error
	if err != nil {
		return nil, 0, err
	}
	as.db.Model(&t).Preload("Categories").Preload("Owner").Offset(offset).Limit(limit).Order("created_at desc").Association("Categories").Find(&products)
	count = as.db.Model(&t).Association("Categories").Count()
	return products, count, nil
}

func (as *ProductRepository) ListByOwner(username string, offset, limit int) ([]models.Product, int, error) {
	var (
		u        models.User
		products []models.Product
		count    int
	)
	err := as.db.Where(&models.User{Username: username}).First(&u).Error
	if err != nil {
		return nil, 0, err
	}
	as.db.Where(&models.Product{OwnerID: u.ID}).Preload("Categories").Preload("Owner").Offset(offset).Limit(limit).Order("created_at desc").Find(&products)
	as.db.Where(&models.Product{OwnerID: u.ID}).Model(&models.Product{}).Count(&count)

	return products, count, nil
}

func (as *ProductRepository) ListCategories(offset, limit int) ([]models.Category, int, error) {
	var (
		categories []models.Category
		count int
	)

	as.db.Model(&categories).Count(&count)
	as.db.Offset(offset).Limit(limit).Order("created_at desc").Find(&categories)
	
	return categories, count, nil
}

func (as *ProductRepository) CreateCategory(c *models.Category) error {
	tx := as.db.Begin()
	if err := tx.Create(&c).Error; err != nil {
		return err
	}

	if err := tx.Where(c.ID).Find(&c).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (as *ProductRepository) UpdateCategory(c *models.Category) error {
	tx := as.db.Begin()
	if err := tx.Model(c).Update(c).Error; err != nil {
		return err
	}

	if err := tx.Where(c.ID).Find(c).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (as *ProductRepository) DeleteCategory(c *models.Category) error {
	return as.db.Delete(c).Error
}

func (as *ProductRepository) GetCategoryByID(id uint) (*models.Category, error) {
	var c models.Category

	err := as.db.Where(id).First(&c).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &c, err
}
