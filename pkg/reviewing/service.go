package reviewing

// Repository provide access to review storage
type Repository interface {
	AddReview(Review) error
}

// Service provide access to reviewing operation
type Service interface {
	AddReview(Review) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) AddReview(r Review) error {
	err := s.r.AddReview(r)

	if err != nil {
		return err
	}
	return nil
}
