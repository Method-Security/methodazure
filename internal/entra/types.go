// Package entra provides functions and data structures to interact with Azure Entra ID.
package entra

import (
    "time"
)

// UserDetails contains basic information about a user
type UserDetails struct {
	ID                string `json:"id"`
	DisplayName       string `json:"display_name"`
	UserPrincipalName string `json:"user_principal_name"`
	Mail              string `json:"mail"`
}

// Group represents an Azure AD (Entra ID) group
type GroupDetails struct {
    ID          		string    		`json:"id"`
    DisplayName 		string    		`json:"displayName"`
    Description 		string    		`json:"description,omitempty"`
    Mail        		string    		`json:"mail,omitempty"`
    CreatedDateTime 	time.Time 		`json:"createdDateTime"`
    Members     		[]GroupMember 	`json:"members"`
}

// GroupMember represents a member of an Azure AD group
type GroupMember struct {
    ID     string `json:"id"`
}

type AzureResourceReport struct {
	Users 		[]UserDetails 		`json:"users"`
	Errors    	[]string       		`json:"errors"`
}