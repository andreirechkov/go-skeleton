package infra

import (
	"context"
	"database/sql"

	"github.com/andreirechkov/go-skeleton/internal/modules/organisations/domain"
)

type PgRepository struct {
	db *sql.DB
}

func NewPgRepository(db *sql.DB) *PgRepository {
	return &PgRepository{db: db}
}

func (r *PgRepository) List(ctx context.Context) ([]domain.Organisation, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, name, description
		FROM organisations
		ORDER BY name ASC
	`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []domain.Organisation
	for rows.Next() {
		var (
			id, name string
			desc sql.NullString
		)
		if err := rows.Scan(&id, &name, &desc); err != nil { return nil, err }
		var dptr *string
		if desc.Valid { dptr = &desc.String }
		out = append(out, domain.Organisation{ID: id, Name: name, Description: dptr })
	}
	return out, rows.Err()
}