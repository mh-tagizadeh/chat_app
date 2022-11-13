package collections

import (
	"github.com/mh-tagizadeh/chat_app/models"
)

// Role provides methods for you to manage array data more easily.
type Role []models.Role

// Origin convert the collection to role array.
// @return []models.Role
func (u Role) Origin() []models.Role {
	return []models.Role(u)
}

// Len returns the number of elements of the array.
// @return int64
func (u Role) Len() (length int64) {
	return int64(len(u))
}

// IDs returns an array of the role array's ids.
// @return []uint
func (u Role) IDs() (IDs []uint) {
	for _, role := range u {
		IDs = append(IDs, role.ID)
	}
	return IDs
}


