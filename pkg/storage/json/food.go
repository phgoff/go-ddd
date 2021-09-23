package json

import "time"

type Food struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	ShortDesc string    `json:"short_description"`
	Created   time.Time `json:"created"`
}
