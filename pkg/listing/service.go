package listing

import "errors"

var ErrNotFound = errors.New("food not found")

// Repository provide access to food storage
type Repository interface {
	GetFoods() []Food
	GetReviews() []Review
}

// Service provide access to listing operation
type Service interface {
	GetFoods() []Food
	GetReviews() []Review
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetFoods() []Food {
	return s.r.GetFoods()
}

func (s *service) GetReviews() []Review {
	return s.r.GetReviews()
}
