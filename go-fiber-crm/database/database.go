package database

import (
	"gorm.io/gorm"
	_ "gorm.io/dialects/sqlite"


var(
	DBConn *gorm.DB

)