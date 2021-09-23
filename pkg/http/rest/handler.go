package rest

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/phgoff/go-ddd/pkg/adding"
	"github.com/phgoff/go-ddd/pkg/listing"
	"github.com/phgoff/go-ddd/pkg/reviewing"
)

func Handler(l listing.Service, a adding.Service, r reviewing.Service) http.Handler {
	router := httprouter.New()

	router.GET("/foods", getFoods(l))
	router.POST("/foods", addFood(a))

	router.GET("/reviews", getReviews(l))
	router.POST("/reviews", addReview(r))

	return router
}

func addFood(s adding.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		decoder := json.NewDecoder(r.Body)

		var newFood []adding.Food
		err := decoder.Decode(&newFood)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		for _, f := range newFood {
			s.AddFood(f)

		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("New food was added.")

	}
}

func getFoods(s listing.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "apllcation/json")
		list := s.GetFoods()
		json.NewEncoder(w).Encode(list)
	}
}

func getReviews(s listing.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		list := s.GetReviews()
		json.NewEncoder(w).Encode(list)
	}
}

func addReview(s reviewing.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

		decoder := json.NewDecoder(r.Body)

		var newReview []reviewing.Review
		err := decoder.Decode(&newReview)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}

		for _, rev := range newReview {
			err := s.AddReview(rev)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("New Review was added.")

	}
}
