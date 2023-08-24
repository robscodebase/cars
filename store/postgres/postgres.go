package postgres

import (
	"cars/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Postgres struct {
	DB *gorm.DB
}

func NewPostgres() (*Postgres, error) {
	c := config.Get().Postgres
	db, err := gorm.Open("postgres", "host="+c.Host+" port="+c.Port+" user="+c.User+" dbname="+c.Database+" password="+c.Password+" sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}
	return &Postgres{
		DB: db,
	}, nil
}
