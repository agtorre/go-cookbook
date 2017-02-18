package database

import _ "github.com/go-sql-driver/mysql" //we import supported libraries for database/sql
// Exec grabs a new connection
// creates tables, and later drops them
// and issues some queries
func Exec() error {
	db, err := Setup()
	if err != nil {
		return err
	}
	// uncaught error on cleanup, but we always
	// want to cleanup
	defer db.Exec("DROP TABLE example")

	if err := Create(db); err != nil {
		return err
	}

	if err := Query(db); err != nil {
		return err
	}
	return nil

}
