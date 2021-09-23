package adding

// Repository provide access to food storage
type Repository interface {
	AddFood(Food) error
}

// Repository provide access to adding operation
type Service interface {
	AddFood(...Food) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}

}

func (s *service) AddFood(food ...Food) error {
	for _, f := range food {
		_ = s.r.AddFood(f)
	}
	return nil
}
