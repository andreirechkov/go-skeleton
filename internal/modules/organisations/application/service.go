package application

import (
	"context"

	"github.com/andreirechkov/go-skeleton/internal/modules/organisations/domain"
)

// OrganisationService provides use-cases for working with organisations.
type OrganisationService struct {
	repo Repository
}

// NewOrganisationService creates a new OrganisationService with the given repository.
func NewOrganisationService(r Repository) *OrganisationService {
	return &OrganisationService{repo: r}
}

// List returns all organisations from the repository.
func (uc *OrganisationService) List(ctx context.Context) ([]domain.Organisation, error) {
	items, err := uc.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	return items, nil
}

// func (s *OrganisationService) FindByID(ctx context.Context, id Org) (*OrganisationResponse, error) {
// 	org, err := s.repo.FindByID(ctx, id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if org == nil {
// 		return nil, fmt.Errorf("organisation %s not found", id)
// 	}
// 	view := toView(*org)
// 	return &view, nil
// }

// func (s *OrganisationService) Create(ctx context.Context, data CreateOrganisationParams) (*OrganisationView, error) {
// 	// org, err := s.repo.Save()
// }
