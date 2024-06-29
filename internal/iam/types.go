// Package iam provides functions and data structures to interact with Azure IAM resources.
package iam

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v2"
)

// RoleDetails contains details about a single Azure role definition
type RoleDetails struct {
	ID             string                          `json:"id" yaml:"id"`
	Name           string                          `json:"name" yaml:"name"`
	RoleDefinition armauthorization.RoleDefinition `json:"role_definition" yaml:"role_definition"`
}

// RoleAssignmentDetails contains details about a single Azure role assignment
type RoleAssignmentDetails struct {
	ID             string                          `json:"id" yaml:"id"`
	RoleAssignment armauthorization.RoleAssignment `json:"role_assignment" yaml:"role_assignment"`
}

// AzureResources contains details about all IAM related resources in the subscription.
type AzureResources struct {
	SubscriptionID  string                  `json:"subscription_id" yaml:"subscription_id"`
	TenantID        string                  `json:"tenant_id" yaml:"tenant_id"`
	Roles           []RoleDetails           `json:"roles" yaml:"roles"`
	RoleAssignments []RoleAssignmentDetails `json:"role_assignments" yaml:"role_assignments"`
}

// AzureResourceReport contains the AzureResources and any non-fatal errors encountered during enumeration.
type AzureResourceReport struct {
	Resources AzureResources `json:"resources" yaml:"resources"`
	Errors    []string       `json:"errors" yaml:"errors"`
}
