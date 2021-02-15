package connection

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var SotDb *gorm.DB

func InitializeMariaDb(dsn string) {
	if SotDb == nil {
		freshDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			log.Fatal("can not create MariaDb Connection")
		}
		SotDb = freshDb
	}
}

func CloseMariDb() {
	if SotDb != nil {
		sqlDb, err := SotDb.DB()
		if err != nil {
			log.Fatal("can not get generic sql object")
		}
		if err := sqlDb.Close(); err != nil {
			log.Fatal("can not close MariaDb connection")
		}
	}
}
