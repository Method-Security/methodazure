// Package entra provides functions and data structures to interact with Azure Entra ID.
package entra

import (
	"time"
)

// ServicePrincipalDetails contains basic information about a service principal
type ServicePrincipalDetails struct {
	ID                        string `json:"id" yaml:"id"`
	DisplayName               string `json:"display_name" yaml:"display_name"`
	AppID                     string `json:"app_id" yaml:"app_id"`
	ServicePrincipalType      string `json:"service_principal_type" yaml:"service_principal_type"`
	AppRoleAssignmentRequired bool   `json:"app_role_assignment_required" yaml:"app_role_assignment_required"`
	AccountEnabled            bool   `json:"account_enabled" yaml:"account_enabled"`
}

// UserDetails contains basic information about a user
type UserDetails struct {
	ID                string `json:"id" yaml:"id"`
	DisplayName       string `json:"display_name" yaml:"display_name"`
	UserPrincipalName string `json:"user_principal_name" yaml:"user_principal_name"`
	Mail              string `json:"mail" yaml:"mail"`
}

// GroupDetails represents an Entra ID group
type GroupDetails struct {
	ID              string        `json:"id" yaml:"id"`
	DisplayName     string        `json:"displayName" yaml:"display_name"`
	Description     string        `json:"description" yaml:"description"`
	CreatedDateTime time.Time     `json:"createdDateTime" yaml:"created_date_time"`
	Members         []GroupMember `json:"members" yaml:"members"`
}

// GroupMember represents a member of an Entra group
type GroupMember struct {
	ID string `json:"id" yaml:"id"`
}

type AzureResources struct {
	TenantID          string                    `json:"tenant_id" yaml:"tenant_id"`
	Users             []UserDetails             `json:"users" yaml:"users"`
	Groups            []GroupDetails            `json:"groups" yaml:"groups"`
	ServicePrincipals []ServicePrincipalDetails `json:"service_principals" yaml:"service_principals"`
}

type AzureResourceReport struct {
	Resources AzureResources `json:"resources" yaml:"resources"`
	Errors    []string       `json:"errors" yaml:"errors"`
}
