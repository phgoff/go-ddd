package json

import (
	"encoding/json"
	"path"
	"runtime"
	"time"

	scribble "github.com/nanobox-io/golang-scribble"
	"github.com/phgoff/go-ddd/pkg/adding"
	"github.com/phgoff/go-ddd/pkg/listing"
	"github.com/phgoff/go-ddd/pkg/reviewing"
	"github.com/phgoff/go-ddd/pkg/storage"
)

const (
	dir              = "/data/"
	CollectionFood   = "foods"
	CollectionReview = "reviews"
)

type Storage struct {
	db *scribble.Driver
}

func NewStorage() (*Storage, error) {
	var err error

	s := new(Storage)
	_, filename, _, _ := runtime.Caller(0)
	p := path.Dir(filename)

	s.db, err = scribble.New(p+dir, nil)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Storage) AddFood(f adding.Food) error {
	id, _ := storage.GetID("food")

	newFood := Food{
		ID:        id,
		Name:      f.Name,
		ShortDesc: f.ShortDesc,
		Created:   time.Now(),
	}

	if err := s.db.Write(CollectionFood, newFood.ID, newFood); err != nil {
		return err
	}

	return nil
}

func (s *Storage) AddReview(r reviewing.Review) error {
	var food Food

	// check if food exists
	if err := s.db.Read(CollectionFood, r.FoodID, &food); err != nil {
		return listing.ErrNotFound
	}

	id, _ := storage.GetID("review")
	newReview := Review{
		ID:      id,
		FoodID:  r.FoodID,
		Name:    r.Name,
		Score:   r.Score,
		Created: time.Now(),
	}

	if err := s.db.Write(CollectionReview, newReview.ID, newReview); err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetFoods() []listing.Food {
	list := []listing.Food{}

	records, err := s.db.ReadAll(CollectionFood)
	if err != nil {
		return list
	}

	for _, r := range records {
		var f Food
		var food listing.Food

		if err := json.Unmarshal([]byte(r), &f); err != nil {
			return list
		}

		food.ID = f.ID
		food.Name = f.Name
		food.ShortDesc = f.ShortDesc
		food.Created = f.Created

		list = append(list, food)
	}

	return list
}

func (s *Storage) GetReviews() []listing.Review {
	list := []listing.Review{}

	records, err := s.db.ReadAll(CollectionReview)
	if err != nil {
		return list
	}

	for _, record := range records {
		var r Review
		var review listing.Review

		if err := json.Unmarshal([]byte(record), &r); err != nil {
			return list
		}

		review.ID = r.ID
		review.FoodID = r.FoodID
		review.Name = r.Name
		review.Score = r.Score
		review.Created = r.Created

		list = append(list, review)
	}

	return list
}
