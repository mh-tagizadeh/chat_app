package repositories

import (
	"gorm.io/gorm"
	"github.com/mh-tagizadeh/chat_app/models"
	"github.com/mh-tagizadeh/chat_app/collections"
	"github.com/mh-tagizadeh/chat_app/repositories/scopes"
)


type IUserRepository interface {
	Migratable

	// single fetch options

	GetUserByID(ID uint) (user models.User, err error)


	// multipile fetch options

	GeUsers(ID uint) (user models.User, err error)

	// ID fetch options

	GetUserIDs(pagination scopes.GormPager) (userIDs []uint, totalCount int64, err error)

	// FirstOrCreate & Updates & Delete

	FirstOrCreate(user *models.User) (err error)
	Updates(user *models.User, updates map[string]interface{}) (err error)
	Delete(user *models.User) (err error)

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


// SINGLE FETCH OPTIONS 

// GetUserByID get user by id 
// @param uint 
// @return models.User, error
func (repository *UserRepository) GetUserByID(ID uint) (user models.User, err error) {
	err = repository.Database.First(&user, "users.id = ?", ID).Error
	return 
}


// MULTIPLE FETCH OPTIONS

// GetRoles get roles by ids.
// @param []uint
// @return collections.Role, error
func (repository *UserRepository) GetUsers(IDs []uint) (users collections.User, err error) {
	err = repository.Database.Where("users.id IN (?)", IDs).Find(&users).Error
	return
}

