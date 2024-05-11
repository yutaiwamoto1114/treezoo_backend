// lib/model/animal_summary_model.go
package model

import "time"

type Animal struct {
	AnimalId           int       `json:"animal_id"`
	AnimalName         string    `json:"animal_name"`
	Species            string    `json:"species"`
	Birthday           time.Time `json:"birthday"`
	Age                int       `json:"age"`
	Gender             string    `json:"gender"`
	BirthZooId         int       `json:"birth_zoo_id"`
	BirthZooName       string    `json:"birth_zoo_name"`
	BirthZooLocation   string    `json:"birth_zoo_location"`
	CurrentZooId       int       `json:"current_zoo_id"`
	CurrentZooName     string    `json:"current_zoo_name"`
	CurrentZooLocation string    `json:"current_zoo_location"`
	Parents            []int     `json:"parents"`
	Children           []int     `json:"children"`
}
