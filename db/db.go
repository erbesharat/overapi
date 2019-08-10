package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	db *sql.DB
}

func New(path string) (*DB, error) {
	database, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("Couldn't open the database: %s", err.Error())
	}

	db := &DB{db: database}

	if err := db.Clear(); err != nil {
		return nil, fmt.Errorf("Couldn't clear the database: %s", err.Error())
	}

	return db, nil
}

func (d *DB) createTables() error {
	_, err := d.db.Exec("CREATE TABLE IF NOT EXISTS heroes (id INTEGER PRIMARY KEY, name TEXT, realName TEXT, health INTEGER, armour INTEGER, shield INTEGER)")
	if err != nil {
		return fmt.Errorf("Couldn't create the heroes table: %s", err.Error())
	}

	_, err = d.db.Exec("CREATE TABLE IF NOT EXISTS abilities (id INTEGER PRIMARY KEY, name TEXT, description TEXT, isUltimate INTEGER, heroID INTEGER)")
	if err != nil {
		return fmt.Errorf("Couldn't create the heroes table: %s", err.Error())
	}

	return nil
}

func (d *DB) Close() error {
	return d.db.Close()
}

func (d *DB) Clear() error {
	_, err := d.db.Exec("DROP TABLE IF EXISTS heroes")
	if err != nil {
		return fmt.Errorf("Couldn't drop the heroes table: %s", err.Error())
	}

	_, err = d.db.Exec("DROP TABLE IF EXISTS abilities")
	if err != nil {
		return fmt.Errorf("Couldn't drop the abilities table: %s", err.Error())
	}

	if err := d.createTables(); err != nil {
		return fmt.Errorf("Couldn't recreate the tables: %s", err.Error())
	}

	return nil
}
