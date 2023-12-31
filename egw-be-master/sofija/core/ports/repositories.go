package ports

import (
	"context"

	domain "github.com/Bloxico/exchange-gateway/sofija/core/domain"
)

type EgwUserRepo interface {
	Insert(ctx context.Context, user *domain.EgwUser) error
	Update(ctx context.Context, user *domain.EgwUser) error
	FindByID(ctx context.Context, id string) (*domain.EgwUser, error)
	FindByEmail(ctx context.Context, email string) (*domain.EgwUser, error)
}

type EgwProductRepo interface {
	Insert(ctx context.Context, product *domain.EgwProduct) error
	Update(ctx context.Context, product *domain.EgwProduct) error
	FindByID(ctx context.Context, id string) (*domain.EgwProduct, error)
	Delete(ctx context.Context, id string) error
}
