package connection

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var sotDb *gorm.DB

func InitializeMariaDb(dsn string) {
	if sotDb == nil {
		freshDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			log.Fatal("can not create MariaDb Connection")
		}
		sotDb = freshDb
	}
}

func GetMariDbConnection() *gorm.DB {
	if sotDb == nil {
		log.Fatal("there are no connection opened.")
	}
	return sotDb
}

func CloseMariaDb() {
	if sotDb != nil {
		sqlDb, err := sotDb.DB()
		if err != nil {
			log.Fatal("can not get generic sql object")
		}
		if err := sqlDb.Close(); err != nil {
			log.Fatal("can not close MariaDb connection")
		}
	}
}
