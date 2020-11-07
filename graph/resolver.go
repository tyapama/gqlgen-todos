package graph

import (
	"context"
	"fmt"
	"github.com/tyapama/gqlgen-todos/graph/model"
	"math/rand"
)

//go:generate go run github.com/99designs/gqlgen
type Resolver struct{
	todos []*model.Todo
}

func (r *mutationResolver) createTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Text:   input.Text,
		ID:     fmt.Sprintf("T%d", rand.Int()),
		UserID: input.UserID, // fix this line
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *todoResolver) getUser(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return &model.User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
}

func(r *queryResolver) getTodos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos,nil
}
