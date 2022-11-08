package repositories

import (
	"gorm.io/gorm"
	"github.com/mh-tagizadeh/chat_app/models"
)


type IUserRepository interface {
}


// UserRepository its data access layer of user.
type UserRepository struct {
	Database *gorm.DB
}

// Migrate generate tables from the databases 
// @return error 
func (repository *UserRepository) Migrate() (err error) {
	err = repository.Database.AutoMigrate(models.User{})
	return 
}

