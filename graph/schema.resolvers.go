package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"


	"github.com/amar-jay/go-graphql/graph/generated"
	"github.com/amar-jay/go-graphql/graph/model"
	"github.com/amar-jay/go-graphql/db"
)

var database = db.Connnection()

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := database.Create(&input)
	if todo == nil {
		panic(fmt.Errorf("Cannot create Todo"))
	}
	return todo, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	todos := database.GetAll()
	if todos == nil {
	panic(fmt.Errorf("Cannot fetch todos"))
	}


	return todos, nil
}

// Todo is the resolver for the todo field.
func (r *queryResolver) Todo(ctx context.Context, id string) (*model.Todo, error) {
	todo := database.GetByID(id)
	if todo == nil {
	panic(fmt.Errorf("not implemented"))
	}

	return todo, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
