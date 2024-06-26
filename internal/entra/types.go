// Package entra provides functions and data structures to interact with Azure Entra ID.
package entra

import (
	"time"
)

// ServicePrincipalDetails contains basic information about a service principal
type ServicePrincipalDetails struct {
	ID                        string `json:"id"`
	DisplayName               string `json:"display_name"`
	AppID                     string `json:"app_id"`
	ServicePrincipalType      string `json:"service_principal_type"`
	AppRoleAssignmentRequired bool   `json:"app_role_assignment_required"`
	AccountEnabled            bool   `json:"account_enabled"`
}

// UserDetails contains basic information about a user
type UserDetails struct {
	ID                string `json:"id"`
	DisplayName       string `json:"display_name"`
	UserPrincipalName string `json:"user_principal_name"`
	Mail              string `json:"mail"`
}

// GroupDetails represents an Entra ID group
type GroupDetails struct {
	ID              string        `json:"id"`
	DisplayName     string        `json:"displayName"`
	Description     string        `json:"description"`
	CreatedDateTime time.Time     `json:"createdDateTime"`
	Members         []GroupMember `json:"members"`
}

// GroupMember represents a member of an Entra group
type GroupMember struct {
	ID string `json:"id"`
}

type AzureResources struct {
	TenantID          string                    `json:"tenant_id"`
	Users             []UserDetails             `json:"users"`
	Groups            []GroupDetails            `json:"groups"`
	ServicePrincipals []ServicePrincipalDetails `json:"service_principals"`
}

type AzureResourceReport struct {
	Resources AzureResources `json:"resources"`
	Errors    []string       `json:"errors"`
}
