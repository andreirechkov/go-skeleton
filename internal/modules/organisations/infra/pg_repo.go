package infra

import (
	"context"
	"database/sql"

	"github.com/andreirechkov/go-skeleton/internal/modules/organisations/domain"
)

// PgRepository implements the organisations repository backed by PostgreSQL.
type PgRepository struct {
	db *sql.DB
}

// NewPgRepository creates a new PgRepository with the given database connection.
func NewPgRepository(db *sql.DB) *PgRepository {
	return &PgRepository{db: db}
}

// List retrieves all organisations from the database.
func (r *PgRepository) List(ctx context.Context) ([]domain.Organisation, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, name, description
		FROM organisations
		ORDER BY name ASC
	`)

	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var out []domain.Organisation
	for rows.Next() {
		var (
			id, name string
			desc     sql.NullString
		)
		if err := rows.Scan(&id, &name, &desc); err != nil {
			return nil, err
		}
		out = append(out, toDomain(id, name, desc))
	}
	return out, rows.Err()
}
