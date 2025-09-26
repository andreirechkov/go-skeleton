package http

import "github.com/andreirechkov/go-skeleton/internal/modules/organisations/domain"

// ToResponse maps a domain Organisation entity into an HTTP Response DTO.
func ToResponse(d domain.Organisation) Response {
	return Response{
		ID:   d.ID,
		Name: d.Name,
		Desc: d.Description,
	}
}

// ToResponseSlice maps a slice of domain Organisation entities into a slice of HTTP Response DTOs.
func ToResponseSlice(dd []domain.Organisation) []Response {
	out := make([]Response, 0, len(dd))
	for _, d := range dd {
		out = append(out, ToResponse(d))
	}
	return out
}
