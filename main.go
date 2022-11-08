package chat_app


import (
	"errors"

	"gorm.io/gorm"
	"github.com/mh-tagizadeh/chat_app/repositories"
)

// Options has the options for initating the chat_app
type Options struct {
	Migrate bool
	DB *gorm.DB
}

// New initializer for chat_app
// If migration is true, it generate all tables in the database if they don't exist.
func (opts Options) New() {
	userRepository := &repositories.UserRepository{Database: opts.DB}


	if opts.Migrate {
		err = repositories.Migrates(userRepository)
		if err != nil {
			return nil, err
		}
	}
}
