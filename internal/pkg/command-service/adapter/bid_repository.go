package adapter

import "gorm.io/gorm"

type MariaDbBidRepository struct {
	Db *gorm.DB
}

func (m MariaDbBidRepository) 