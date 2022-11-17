package repositories

import (
	"chat_app/collections"
	"chat_app/models"
	"chat_app/scopes"

	"gorm.io/gorm"
)

type RoomRepositoryI interface {
	Migratable

	GetRoomByID(ID uint) (room models.Room, err error)
	GetRooms(ID uint) (room models.Room, err error)
	GetRoomIDs(pagination scopes.GormPager) (userIDs []uint, totalCount int64, err error)
	FirstOrCreate(room models.Room) (err error)                          //? First ?! Why Pointer ?!
	Update(room *models.Room, update map[string]interface{}) (err error) //? Updates instead of Update ?! Update parameter type ?!
	Delete(room *models.Room) (err error)
}

type RoomRepository struct {
	Database *gorm.DB
}

func (repository *RoomRepository) Migrate() (err error) {
	err = repository.Database.AutoMigrate(models.Room{})
	return
}

func (repository *RoomRepository) GetRoomByID(ID uint) (room models.Room, err error) {
	err = repository.Database.First(&room, "rooms.id = ?", ID).Error
	return
}

func (repository *RoomRepository) GetRooms(IDs []uint) (rooms collections.Room, err error) {
	err = repository.Database.Where("rooms.ID IN (?)", IDs).Find(&rooms).Error
	return
}
