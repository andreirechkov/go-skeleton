// Package infra contains the PostgreSQL infrastructure layer for the organisations module.
package infra

import (
	"database/sql"

	"github.com/andreirechkov/go-skeleton/internal/modules/organisations/domain"
)

func toDomain(id string, name string, desc sql.NullString) domain.Organisation {
	var dptr *string
	if desc.Valid {
		dptr = &desc.String
	}
	return domain.Organisation{ID: id, Name: name, Description: dptr}
}
