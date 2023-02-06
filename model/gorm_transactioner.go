package model

import "gorm.io/gorm"

type GormTransactionerRepository interface {
	BeginTransaction() *gorm.DB
	Rollback(tx *gorm.DB) *gorm.DB
	Commit(tx *gorm.DB) *gorm.DB
}
