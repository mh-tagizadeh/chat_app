package models

type Room struct {
	ID uint `gorm:"primaryKey"`
}

func (Room) GetTableName() string {
	return "rooms"
}
