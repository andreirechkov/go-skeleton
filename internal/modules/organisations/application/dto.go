package application

type OrganisationView struct {
	ID 		string `json:"id"`
	Name 	string `json:"name"`
	Desc 	string `json:"description,omitempty"`
}