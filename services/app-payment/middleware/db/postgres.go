package db

import "gorm.io/gorm"

type Postgres struct {
	db gorm.DB
}
