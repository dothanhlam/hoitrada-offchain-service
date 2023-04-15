package wallet

import (
	"context"

	"github.com/google/uuid"
)

type Service interface {
	Store(ctx context.Context, w Wallet) (uuid.UUID, error)
}

type service struct{}

func NewService() *service {
	return &service{}
}
func (s *service) Store(ctx context.Context, w Wallet) (uuid.UUID, error) {
	//@TODO create store rules, using databases or something else
	return w.ID, nil
}
