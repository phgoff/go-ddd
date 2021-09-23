# Go structure using Domain Drive Design

# Design:
    - Context: foodreview
    - Language: food, review, storage, json
    - Entities: food, review
    - Value objects: reviewer
    - Aggregate: -
    - Service: adding food, listing food, adding review, listing review
    - Event: food added, review added, food not found
    - Repository: food repository, review Repository

# Run
```
    go build cmd/server/main.go
    ./main
```
