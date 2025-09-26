package application

import (
	"context"

	"github.com/andreirechkov/go-skeleton/internal/modules/organisations/domain"
)

// Repository defines the application port for accessing organisation data.
// This is realisation only for demo (sceleton), hexo arhitecture (ports & adapters)
type Repository interface {
	List(ctx context.Context) ([]domain.Organisation, error)
	// FindByID(ctx context.Context, id string) (*domain.Organisation, error)
	// Create(ctx context.Context, org any) error
	// Update(ctx context.Context, id string, org any) error
	// Delete(ctx context.Context, id string) error
}
