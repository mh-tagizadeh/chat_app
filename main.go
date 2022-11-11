package main


import (
	"gorm.io/gorm"
    "gorm.io/driver/postgres"


	"github.com/mh-tagizadeh/chat_app/repositories"
)

// Options has the options for initating the chat_app
type Options struct {
	Migrate bool
	DB *gorm.DB
}

// New initializer for chat_app
// If migration is true, it generate all tables in the database if they don't exist.
func InitializeDatabase(opts Options) (err error) {
	userRepository := &repositories.UserRepository{Database: opts.DB}


	if opts.Migrate {
		err = repositories.Migrates(userRepository)
		if err != nil {
			return err
		}
	}

	return 
}


func main(){
	dsn := "host=localhost user=admin password=laravel dbname=chat_app_db port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})// hello

	InitializeDatabase(Options{
		Migrate: true,
		DB: db,
	})
}
