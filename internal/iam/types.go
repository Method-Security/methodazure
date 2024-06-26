// Package iam provides functions and data structures to interact with Azure IAM resources.
package iam

import (
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v2"
)

// RoleDetails contains details about a single Azure role definition
type RoleDetails struct {
	ID             string
	Name           string
	RoleDefinition armauthorization.RoleDefinition
}

// RoleAssignmentDetails contains details about a single Azure role assignment
type RoleAssignmentDetails struct {
	ID             string
	RoleAssignment armauthorization.RoleAssignment
}

// AzureResources contains details about all IAM related resources in the subscription.
type AzureResources struct {
	SubscriptionID  string                  `json:"subscription_id"`
	Roles           []RoleDetails           `json:"roles"`
	RoleAssignments []RoleAssignmentDetails `json:"role_assignments"`
}

// AzureResourceReport contains the AzureResources and any non-fatal errors encountered during enumeration.
type AzureResourceReport struct {
	Resources AzureResources `json:"resources"`
	Errors    []string       `json:"errors"`
}
