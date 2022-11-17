package collections

import "chat_app/models"

type Room []models.Room

func (r Room) Origin() []models.Room {
	return []models.Room(r) //? What does it do ?! PURPOSE ?!
}

func (r Room) Len() (length int64) {
	return int64(len(r)) //? Why int64 ?!
}

func (r Room) IDs() (IDs []uint) {
	for _, room := range r {
		IDs = append(IDs, room.ID)
	}
	return IDs
}
