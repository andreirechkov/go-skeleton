package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/andreirechkov/go-skeleton/internal/modules/organisations/application"
	"github.com/andreirechkov/go-skeleton/internal/modules/organisations/infra"
)

func RegisterOrganisationRoutes(mux *http.ServeMux, db *sql.DB) {
	repo := infra.NewPgRepository(db)
	listUC := application.NewListUseCase(repo)

	mux.HandleFunc("/organisations", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed); return
		}
		result, err := listUC.Execute(r.Context())
		if err != nil { http.Error(w, err.Error(), http.StatusInternalServerError); return }
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	})
}