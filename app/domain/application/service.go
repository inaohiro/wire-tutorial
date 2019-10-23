package application

import (
	"context"
	"github.com/google/wire"
	"github.com/inaohiro-sandbox/wire-tutorial3/app/domain/domain"
	"log"
)

var Set = wire.NewSet(NewService)

type Service interface {
	Create(ctx context.Context, name string) (created *domain.User, err error)
	Show(ctx context.Context, id int64) (user *domain.User, err error)
}

func NewService(repo domain.UserRepository) Service {
	return &service{repo: repo}
}

type service struct {
	repo domain.UserRepository
}

func (service *service) Create(ctx context.Context, name string) (created *domain.User, err error) {
	user := &domain.User{Name: name}
	err = service.repo.Save(ctx, user)
	if err != nil {
		log.Printf("fail: %v", err)
	}

	return user, nil
}

func (service *service) Show(ctx context.Context, id int64) (user *domain.User, err error) {
	user, err = service.repo.Get(ctx, id)
	if err != nil {
		log.Printf("fail: %v", err)
	}
	return user, nil
}
