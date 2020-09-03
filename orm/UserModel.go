package orm

import (
	"github.com/jinzhu/gorm"
)

// User 测试User Model
type User struct {
	gorm.Model
	Name string
}
