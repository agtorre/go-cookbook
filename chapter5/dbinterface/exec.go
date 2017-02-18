package dbinterface

import "github.com/agtorre/go-cookbook/chapter5/database"

// Exec replaces the Exec from the previous
// recipe
func Exec() error {
	db, err := database.Setup()
	if err != nil {
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	// this wont do anything if commit is successful
	defer tx.Rollback()
	// uncaught error on cleanup, but we always
	// want to cleanup
	defer db.Exec("DROP TABLE example")

	if err := Create(tx); err != nil {
		return err
	}

	if err := Query(tx); err != nil {
		return err
	}

	err = tx.Commit()
	return err
}
