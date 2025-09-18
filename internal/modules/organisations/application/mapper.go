package application

import "github.com/andreirechkov/go-skeleton/internal/modules/organisations/domain"

func toView(d domain.Organisation) OrganisationView {
	v := OrganisationView{ID: d.ID, Name: d.Name}
	if d.Description != nil { v.Desc = *d.Description }
	return v
}

func toViewSlice(dd []domain.Organisation) []OrganisationView {
	out := make([]OrganisationView, 0, len(dd))
	for _, d := range dd { out = append(out, toView(d))}
	return out
}