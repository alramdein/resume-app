package model

import "time"

type Occupation struct {
	Name      string     `json:"name"`
	Position  string     `json:"position"`
	StartDate time.Time  `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`
}
