package handler

import (
	"log"
	"os"
	"testing"

	"encoding/json"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo/v4"
	"github.com/sumitalp/productcatalog/db"
	"github.com/sumitalp/productcatalog/models"
	"github.com/sumitalp/productcatalog/product"
	"github.com/sumitalp/productcatalog/repository"
	"github.com/sumitalp/productcatalog/router"
	"github.com/sumitalp/productcatalog/user"
)

var (
	d  *gorm.DB
	us user.RepositoryInterface
	as product.RepositoryInterface
	h  *Handler
	e  *echo.Echo
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func authHeader(token string) string {
	return "Token " + token
}

func setup() {
	d = db.TestDB()
	db.AutoMigrate(d)
	us = repository.NewUserRepository(d)
	as = repository.NewProductRepository(d)
	h = NewHandler(us, as)
	e = router.New()
	loadFixtures()
}

func tearDown() {
	_ = d.Close()
	if err := db.DropTestDB(); err != nil {
		log.Fatal(err)
	}
}

func responseMap(b []byte, key string) map[string]interface{} {
	var m map[string]interface{}
	json.Unmarshal(b, &m)
	return m[key].(map[string]interface{})
}

func loadFixtures() error {
	u1bio := "user1 bio"
	u1image := "http://realworld.io/user1.jpg"
	u1 := models.User{
		Username: "user1",
		Email:    "user1@email.io",
		Bio:      &u1bio,
		Image:    &u1image,
	}
	u1.Password, _ = u1.HashPassword("secret")
	if err := us.Create(&u1); err != nil {
		return err
	}

	u2bio := "user2 bio"
	u2image := "http://realworld.io/user2.jpg"
	u2 := models.User{
		Username: "user2",
		Email:    "user2@email.io",
		Bio:      &u2bio,
		Image:    &u2image,
	}
	u2.Password, _ = u2.HashPassword("secret")
	if err := us.Create(&u2); err != nil {
		return err
	}

	c := models.Category{
		Category:    "category1",
		Description: "category1 description",
	}

	as.CreateCategory(&c)

	c2 := models.Category{
		Category:    "category2",
		Description: "category2 description",
	}

	as.CreateCategory(&c2)

	a := models.Product{
		Slug:        "product1-slug",
		Title:       "product1 title",
		Description: "product1 description",
		OwnerID:     1,
		Categories: []models.Category{
			{
				Category: "category1",
			},
			{
				Category: "category2",
			},
		},
	}
	as.CreateProduct(&a)

	a2 := models.Product{
		Slug:        "product2-slug",
		Title:       "product2 title",
		Description: "product2 description",
		OwnerID:     1,
		Categories: []models.Category{
			{
				Category: "category1",
			},
			{
				Category: "category2",
			},
		},
	}
	as.CreateProduct(&a2)

	return nil
}
