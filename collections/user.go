package collections

import (
	"chat_app/models"
)

// User provides methods for you to manage array data more easily.
type User []models.User

// Origin convert the collection to user array.
// @return []models.User
func (u User) Origin() []models.User {
	return []models.User(u)
}

// Len returns the number of elements of the array.
// @return int64
func (u User) Len() (length int64) {
	return int64(len(u))
}

// IDs returns an array of the user array's ids.
// @return []uint
func (u User) IDs() (IDs []uint) {
	for _, user := range u {
		IDs = append(IDs, user.ID)
	}
	return IDs
}


