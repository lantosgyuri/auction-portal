package adapter

import (
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
	"github.com/lantosgyuri/auction-portal/internal/pkg/connection"
	"gorm.io/gorm"
)

type MariaDbUserRepository struct {
	db *gorm.DB
}

func CreateMariaDbUserRepository() MariaDbUserRepository {
	return MariaDbUserRepository{
		db: connection.GetMariDbConnection(),
	}
}

func (m MariaDbUserRepository) SaveUserEvent(event domain.UserEventRaw) error {
	return m.db.Create(&event).Error
}

func (m MariaDbUserRepository) CreateUser(user domain.User) error {
	return m.db.Create(&user).Error
}

func (m MariaDbUserRepository) DeleteUser(user domain.User) error {
	return m.db.Delete(&user).Error
}
