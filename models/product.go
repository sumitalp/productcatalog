package models

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Slug        string `gorm:"unique_index;not null"`
	Title       string `gorm:"not null"`
	Description string
	Image       string
	Owner       User
	OwnerID     uint
	Categories  []Category `gorm:"many2many:product_categories;association_autocreate:false"`
}

type Category struct {
	gorm.Model
	Category    string `gorm:"unique_index"`
	Description string
	Products    []Product `gorm:"many2many:product_tags;"`
}
