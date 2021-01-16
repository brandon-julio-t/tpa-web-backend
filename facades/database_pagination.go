package facades

import (
	"gorm.io/gorm"
)

func UsePagination(currentPage, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (currentPage - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
