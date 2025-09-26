package http

import (
	"database/sql"
	"net/http"

	"github.com/andreirechkov/go-skeleton/internal/modules/organisations/application"
	"github.com/andreirechkov/go-skeleton/internal/modules/organisations/infra"
)

// RegisterOrganisationRoutes registers HTTP routes for the organisations module.
func RegisterOrganisationRoutes(mux *http.ServeMux, db *sql.DB) {
	repo := infra.NewPgRepository(db)
	service := application.NewOrganisationService(repo)

	mux.HandleFunc("/api/organisations", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			HandleListOrganisations(service)(w, r)
		default:
			methodNotAllowed(w)
		}
	})
}
