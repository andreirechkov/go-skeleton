// Package domain contains the core business entities for the organisations module.
package domain

import "fmt"

// Organisation represents the organisation aggregate root in the domain layer.
type Organisation struct {
	ID          string
	Name        string
	Description *string // nulable
}

// NewOrganisation creates a new Organisation entity after validating input data.
func NewOrganisation(id string, name string, desc *string) (*Organisation, error) {
	if id == "" || name == "" {
		return nil, fmt.Errorf("cannot create organisation")
	}
	return &Organisation{ID: id, Name: name, Description: desc}, nil
}
