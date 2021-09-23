package listing

import "time"

type Review struct {
	ID      string    `json:"id"`
	FoodID  string    `json:"food_id"`
	Name    string    `json:"name"`
	Score   int       `json:"score"`
	Created time.Time `json:"created"`
}
