package repositories

// Migratable gives models the ability to migrate.
type Migratable interface {
	Migrate() error
}


// Migrates migrate to migratable models.
func Migrates(repos ...Migratable) (err error) {
	for _, r := range repos {
		err = r.Migrate()
	}
	return 
}


