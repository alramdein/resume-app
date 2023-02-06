package repository

import (
	"github.com/alramdein/karirlab-test/model"
	"gorm.io/gorm"
)

type gormTransactioner struct {
	db *gorm.DB
}

func NewGormTransactioner(db *gorm.DB) model.GormTransactionerRepository {
	return &gormTransactioner{
		db: db,
	}
}

func (g *gormTransactioner) BeginTransaction() *gorm.DB {
	return g.db.Begin()
}

func (g *gormTransactioner) Rollback(tx *gorm.DB) *gorm.DB {
	return tx.Rollback()
}

func (g *gormTransactioner) Commit(tx *gorm.DB) *gorm.DB {
	return tx.Commit()
}
