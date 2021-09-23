package adding

import (
	"log"
	"testing"

	"github.com/phgoff/go-ddd/pkg/listing"
)

type mockStorage struct {
	foods []Food
}

func TestAddFoods(t *testing.T) {
	f1 := Food{
		Name:      "Food1",
		ShortDesc: "This is delicious",
	}

	f2 := Food{
		Name:      "Food2",
		ShortDesc: "Not bad!",
	}

	mS := new(mockStorage)

	s := NewService(mS)

	err := s.AddFood(f1, f2)

	if err != nil {
		log.Fatal(err)
	}

	foods := mS.GetFoods()

	if len(foods) != 2 {
		t.Errorf("Add Foods = %d; want 2", len(foods))
	}

}

func (m *mockStorage) AddFood(f Food) error {
	m.foods = append(m.foods, f)
	return nil
}

func (m *mockStorage) GetFoods() []listing.Food {
	foods := []listing.Food{}

	for _, food := range m.foods {

		f := listing.Food{
			Name:      food.Name,
			ShortDesc: food.ShortDesc,
		}

		foods = append(foods, f)
	}
	return foods
}
