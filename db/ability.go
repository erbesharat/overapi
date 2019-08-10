package db

import (
	"database/sql"
	"fmt"
)

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

func (d *DB) GetAbilities() ([]Ability, error) {
	abilities := []Ability{}
	rows, err := d.db.Query("SELECT id, name, description, isUltimate FROM abilities")
	if err != nil {
		return nil, fmt.Errorf("Couldn't select the heroes from the database: %s", err.Error())
	}
	for rows.Next() {
		ability := Ability{}
		rows.Scan(&ability.ID, &ability.Name, &ability.Description, &ability.IsUltimate)
		abilities = append(abilities, ability)
	}

	return abilities, nil
}

func (d *DB) GetAbility(id string) (*Ability, error) {
	ability := Ability{}
	row := d.db.QueryRow("SELECT id, name, description, isUltimate FROM abilities WHERE id = ?", id)
	err := row.Scan(&ability.ID, &ability.Name, &ability.Description, &ability.IsUltimate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("Couldn't select the record: %s", err.Error())
	}
	return &ability, nil
}

func (d *DB) GetHeroAbilities(id string) ([]Ability, error) {
	abilities := []Ability{}
	rows, err := d.db.Query("SELECT id, name, description, isUltimate FROM abilities WHERE heroID = ?", id)
	if err != nil {
		return nil, fmt.Errorf("Couldn't select the abilities: %s", err.Error())
	}
	for rows.Next() {
		ability := Ability{}
		rows.Scan(&ability.ID, &ability.Name, &ability.Description, &ability.IsUltimate)
		abilities = append(abilities, ability)
	}
	return abilities, nil
}
