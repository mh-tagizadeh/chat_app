package models


type User struct {
	ID uint `gorm:"primary_key" json:"id"`
	Username string `gorm:"size:255;not null" json:"name"`
	Password string `gorm:"size:255;not null" json:"password"`
	// todo: CREATED_AT AND UDPATED_AT
}


func (User) TableName() string {
	return "users"
}
