package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"chat_app/repositories"
)

var (
	Connection *gorm.DB
)

func Connect(connectionString string) {
	var err error

	Connection, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		panic("Couldn't connect to the database")
	}
}


// Options has the options for initating the chat_app
type Options struct {
	Migrate bool
	DB *gorm.DB
}

// New initializer for chat_app
// If migration is true, it generate all tables in the database if they don't exist.
func Migrates(opts Options) (err error) {
	userRepository := &repositories.UserRepository{Database: opts.DB}


	if opts.Migrate {
		err = repositories.Migrates(userRepository)
		if err != nil {
			return err
		}
	}

	return 
}

