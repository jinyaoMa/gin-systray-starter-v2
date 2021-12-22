package models

import "gorm.io/gorm"

type Record struct {
	gorm.Model
}

func (r *Record) Create() *gorm.DB {
	return db.Create(r)
}
