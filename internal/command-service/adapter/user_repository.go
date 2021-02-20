package adapter

import (
	"github.com/lantosgyuri/auction-portal/internal/command-service/domain"
	"gorm.io/gorm"
)

type MariaDbUserRepository struct {
	Db *gorm.DB
}

func (m MariaDbUserRepository) SaveUserEvent(event domain.UserEventRaw) error {
	return m.Db.Create(&event).Error
}

func (m MariaDbUserRepository) CreateUser(user domain.User) error {
	return m.Db.Create(&user).Error
}

func (m MariaDbUserRepository) DeleteUser(user domain.User) error {
	return m.Db.Delete(&user).Error
}
