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
	Favorites   []User `gorm:"many2many:favorites;"`
	Tags        []Tag  `gorm:"many2many:product_tags;association_autocreate:false"`
}

type Tag struct {
	gorm.Model
	Tag      string    `gorm:"unique_index"`
	Products []Product `gorm:"many2many:product_tags;"`
}
