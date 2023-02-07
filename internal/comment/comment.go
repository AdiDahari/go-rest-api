package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingComment = errors.New("Failed ro fetch comment by id.")
	ErrNotImplemented = errors.New("Not implemented.")
)

// Comment - representation of the comment structure
type Comment struct {
	ID string
	Slug string
	Body string
	Author string
}

// Store - interface that defines all methods for Service operation
type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

// Service - struct which holds all logic
type Service struct{
	Store Store
}

// NewService - returns a pointer to a new Servce
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("Retrieving a Comment")
	cmt, err := s.Store.GetComment(ctx, id)

	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrFetchingComment
	}

	return cmt, nil
}

func (s *Service) UpdateComment(ctx context.Context, cmt Comment) error {
	return ErrNotImplemented
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return ErrNotImplemented
}

func (s *Service) CreateComment(ctx context.Context, cmt Comment) (Comment, error) {
	return Comment{}, ErrNotImplemented
}