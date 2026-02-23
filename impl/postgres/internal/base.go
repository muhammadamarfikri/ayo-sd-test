package internal

import "gorm.io/gorm"

type PostgresDB struct {
	Db *gorm.DB
}

func NewPostgresDB(db *gorm.DB) *PostgresDB {
	return &PostgresDB{db}
}
