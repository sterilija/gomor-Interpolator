package pg

import "gorm.io/gorm"

type Repository struct {
	DB *DB
}
type DB struct {
	DB *gorm.DB
}

func New(db *DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) Close() error {
	return r.DB.Close()
}
