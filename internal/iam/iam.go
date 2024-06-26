// Package iam provides functions and data structures to interact with Azure IAM resources.
package iam

import (
	"context"

	"github.com/Method-Security/methodazure/internal/config"
)

// EnumerateIAMResources enumerates all IAM related resources in the subscription, returning a report of the resources and
// any non-fatal errors encountered.
func EnumerateIAMResources(ctx context.Context, cfg config.AzureConfig) (*AzureResourceReport, error) {
	resources := AzureResources{}
	errors := []string{}

	// Get all roles in the subscription
	roles, err := listRoles(ctx, cfg)
	if err != nil {
		errors = append(errors, err.Error())
	} else {
		resources.Roles = roles
	}

	// Get all role assignments in the subscription
	roleAssignments, err := listRoleAssignments(ctx, cfg)
	if err != nil {
		errors = append(errors, err.Error())
	} else {
		resources.RoleAssignments = roleAssignments
	}

	report := AzureResourceReport{
		Resources: resources,
		Errors:    errors,
	}

	return &report, nil
}
