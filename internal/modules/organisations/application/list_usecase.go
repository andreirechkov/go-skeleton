package application

import (
	"context"
	"github.com/andreirechkov/go-skeleton/internal/modules/organisations/domain"
)

type ListUseCase struct {
	repo domain.Repository
}

func NewListUseCase(r domain.Repository) *ListUseCase {
	return &ListUseCase{repo: r}
}

func (uc *ListUseCase) Execute(ctx context.Context) ([]OrganisationView, error) {
	items, err := uc.repo.List(ctx)
	if err != nil { return nil, err }
	return toViewSlice(items), nil
}