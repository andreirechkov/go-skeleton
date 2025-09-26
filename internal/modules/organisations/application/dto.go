// Package application contains the use-case layer for the organisations module.
package application

// CreateOrganisationParams defines parameters for creating an organisation.
type CreateOrganisationParams struct {
	Name string
	Desc *string
}

// UpdateOrganisationParams defines parameters for updating an organisation.
type UpdateOrganisationParams struct {
	Name *string
	Desc *string
}
