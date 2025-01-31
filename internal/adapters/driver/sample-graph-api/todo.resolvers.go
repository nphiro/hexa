package samplegraphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.55

import (
	"context"
	"fmt"

	"github.com/nphiro/hexa/internal/adapters/driver/sample-graph-api/gen"
	"github.com/nphiro/hexa/internal/adapters/driver/sample-graph-api/gen/model"
)

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented: Todos - todos"))
}

// Todo is the resolver for the todo field.
func (r *queryResolver) Todo(ctx context.Context, id string) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented: Todo - todo"))
}

// Query returns gen.QueryResolver implementation.
func (r *Resolver) Query() gen.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
