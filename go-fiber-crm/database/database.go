package database

import (
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var(
	DBConn *gorm.DB

)