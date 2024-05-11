// lib/model/animal_summary_model.go
package model

import "time"

type AnimalSummary struct {
	AnimalId       int       `json:"animal_id"`
	AnimalName     string    `json:"animal_name"`
	Species        string    `json:"species"`
	Birthday       time.Time `json:"birthday"`
	Age            int       `json:"age"`
	Gender         string    `json:"gender"`
	CurrentZooId   int       `json:"current_zoo_id"`
	CurrentZooName string    `json:"current_zoo_name"`
	Parents        []int     `json:"parents"`
	Children       []int     `json:"children"`
}
