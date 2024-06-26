// Package entra provides functions and data structures to interact with Azure Entra ID.
package entra

// UserDetails contains basic information about a user
type UserDetails struct {
	ID                string `json:"id"`
	DisplayName       string `json:"display_name"`
	UserPrincipalName string `json:"user_principal_name"`
	Mail              string `json:"mail"`
}

type AzureResourceReport struct {
	Users 		[]UserDetails 		`json:"users"`
	Errors    	[]string       		`json:"errors"`
}