// Package http provides HTTP handlers for the organisations module.
package http

import (
	"net/http"

	"github.com/andreirechkov/go-skeleton/internal/modules/organisations/application"
)

// HandleListOrganisations returns an HTTP handler for listing organisations.
func HandleListOrganisations(service *application.OrganisationService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		items, err := service.List(r.Context())
		if err != nil {
			httpError(w, http.StatusInternalServerError, "failed to list organisations")
			return
		}
		writeJSON(w, ToResponseSlice(items))
	}
}
