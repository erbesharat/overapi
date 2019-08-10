package db

import (
	"database/sql"
	"fmt"
)

// Hero contains the database field for Heros table
type Hero struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	RealName string `json:"real_name"`
	Health   int    `json:"health"`
	Armour   int    `json:"armour"`
	Shield   int    `json:"shield"`
}

func (d *DB) InsertHero(hero *Hero) error {
	_, err := d.db.Exec("INSERT INTO heroes (name, realName, health, armour, shield) VALUES (?, ?, ?, ?, ?)",
		hero.Name,
		hero.RealName,
		hero.Health,
		hero.Armour,
		hero.Shield,
	)
	if err != nil {
		return fmt.Errorf("Couldn't insert hero to the table: %s", err.Error())
	}

	return nil
}

func (d *DB) GetHeroes() ([]Hero, error) {
	heroes := []Hero{}
	rows, err := d.db.Query("SELECT id, name, realName, health, armour, shield FROM heroes")
	if err != nil {
		return nil, fmt.Errorf("Couldn't select the heroes from the database: %s", err.Error())
	}
	for rows.Next() {
		hero := Hero{}
		rows.Scan(&hero.ID, &hero.Name, &hero.RealName, &hero.Health, &hero.Armour, &hero.Shield)
		heroes = append(heroes, hero)
	}

	return heroes, nil
}

func (d *DB) GetHero(id string) (*Hero, error) {
	hero := Hero{}
	row := d.db.QueryRow("SELECT id, name, realName, health, armour, shield FROM heroes WHERE id = ?", id)
	err := row.Scan(&hero.ID, &hero.Name, &hero.RealName, &hero.Health, &hero.Armour, &hero.Shield)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("Couldn't select the record: %s", err.Error())
	}
	return &hero, nil
}
