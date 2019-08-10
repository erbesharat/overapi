package db

import "fmt"

// Ability contains the database field for Abilities table
type Ability struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsUltimate  bool   `json:"is_ultimate"`
	Hero        *Hero  `json:"hero"`
}

func (d *DB) InsertAbility(ability *Ability) error {
	_, err := d.db.Exec("INSERT INTO abilities (name, description, isUltimate, heroID) VALUES (?, ?, ?, ?)",
		ability.Name,
		ability.Description,
		ability.IsUltimate,
		ability.Hero.ID,
	)
	if err != nil {
		return fmt.Errorf("Couldn't insert ability to the table: %s", err.Error())
	}

	return nil
}
