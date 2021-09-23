package reviewing

type Review struct {
	FoodID string `json:"food_id"`
	Name   string `json:"name"`
	Score  int    `json:"score"`
}
