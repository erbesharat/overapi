package db

import "fmt"

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
