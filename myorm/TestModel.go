package myorm

import (
	"github.com/jinzhu/gorm"
)

// TestModel 测试Model
type TestModel struct {
	gorm.Model
	Name string
}
