package models

import (
	"time"
	"github.com/jinzhu/gorm"
)

type ModelBase struct {
	ID        uint       `gorm:"primary_key"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type Product struct {
	ModelBase
	Slug        string `gorm:"unique_index;not null"`
	Title       string `gorm:"not null"`
	Description string
	Image       string
	Owner       User
	OwnerID     uint
	Categories  []Category `gorm:"many2many:product_categories;association_autocreate:false"`
}

type Category struct {
	ModelBase
	Category    string `gorm:"unique_index"`
	Description string
	Products    []Product `gorm:"many2many:products;"`
}

var GormDB *gorm.DB