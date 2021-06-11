// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type DeploymentStatus string

// Enum values for DeploymentStatus
const (
	DeploymentStatusInProgress       DeploymentStatus = "IN_PROGRESS"
	DeploymentStatusFailed           DeploymentStatus = "FAILED"
	DeploymentStatusSucceeded        DeploymentStatus = "SUCCEEDED"
	DeploymentStatusDeleteInProgress DeploymentStatus = "DELETE_IN_PROGRESS"
	DeploymentStatusDeleteFailed     DeploymentStatus = "DELETE_FAILED"
	DeploymentStatusDeleteComplete   DeploymentStatus = "DELETE_COMPLETE"
	DeploymentStatusCancelling       DeploymentStatus = "CANCELLING"
	DeploymentStatusCancelled        DeploymentStatus = "CANCELLED"
)

// Values returns all known values for DeploymentStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DeploymentStatus) Values() []DeploymentStatus {
	return []DeploymentStatus{
		"IN_PROGRESS",
		"FAILED",
		"SUCCEEDED",
		"DELETE_IN_PROGRESS",
		"DELETE_FAILED",
		"DELETE_COMPLETE",
		"CANCELLING",
		"CANCELLED",
	}
}

type DeploymentUpdateType string

// Enum values for DeploymentUpdateType
const (
	DeploymentUpdateTypeNone           DeploymentUpdateType = "NONE"
	DeploymentUpdateTypeCurrentVersion DeploymentUpdateType = "CURRENT_VERSION"
	DeploymentUpdateTypeMinorVersion   DeploymentUpdateType = "MINOR_VERSION"
	DeploymentUpdateTypeMajorVersion   DeploymentUpdateType = "MAJOR_VERSION"
)

// Values returns all known values for DeploymentUpdateType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DeploymentUpdateType) Values() []DeploymentUpdateType {
	return []DeploymentUpdateType{
		"NONE",
		"CURRENT_VERSION",
		"MINOR_VERSION",
		"MAJOR_VERSION",
	}
}

type EnvironmentAccountConnectionRequesterAccountType string

// Enum values for EnvironmentAccountConnectionRequesterAccountType
const (
	EnvironmentAccountConnectionRequesterAccountTypeManagementAccount  EnvironmentAccountConnectionRequesterAccountType = "MANAGEMENT_ACCOUNT"
	EnvironmentAccountConnectionRequesterAccountTypeEnvironmentAccount EnvironmentAccountConnectionRequesterAccountType = "ENVIRONMENT_ACCOUNT"
)

// Values returns all known values for
// EnvironmentAccountConnectionRequesterAccountType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (EnvironmentAccountConnectionRequesterAccountType) Values() []EnvironmentAccountConnectionRequesterAccountType {
	return []EnvironmentAccountConnectionRequesterAccountType{
		"MANAGEMENT_ACCOUNT",
		"ENVIRONMENT_ACCOUNT",
	}
}

type EnvironmentAccountConnectionStatus string

// Enum values for EnvironmentAccountConnectionStatus
const (
	EnvironmentAccountConnectionStatusPending   EnvironmentAccountConnectionStatus = "PENDING"
	EnvironmentAccountConnectionStatusConnected EnvironmentAccountConnectionStatus = "CONNECTED"
	EnvironmentAccountConnectionStatusRejected  EnvironmentAccountConnectionStatus = "REJECTED"
)

// Values returns all known values for EnvironmentAccountConnectionStatus. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (EnvironmentAccountConnectionStatus) Values() []EnvironmentAccountConnectionStatus {
	return []EnvironmentAccountConnectionStatus{
		"PENDING",
		"CONNECTED",
		"REJECTED",
	}
}

type Provisioning string

// Enum values for Provisioning
const (
	ProvisioningCustomerManaged Provisioning = "CUSTOMER_MANAGED"
)

// Values returns all known values for Provisioning. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (Provisioning) Values() []Provisioning {
	return []Provisioning{
		"CUSTOMER_MANAGED",
	}
}

type ServiceStatus string

// Enum values for ServiceStatus
const (
	ServiceStatusCreateInProgress              ServiceStatus = "CREATE_IN_PROGRESS"
	ServiceStatusCreateFailedCleanupInProgress ServiceStatus = "CREATE_FAILED_CLEANUP_IN_PROGRESS"
	ServiceStatusCreateFailedCleanupComplete   ServiceStatus = "CREATE_FAILED_CLEANUP_COMPLETE"
	ServiceStatusCreateFailedCleanupFailed     ServiceStatus = "CREATE_FAILED_CLEANUP_FAILED"
	ServiceStatusCreateFailed                  ServiceStatus = "CREATE_FAILED"
	ServiceStatusActive                        ServiceStatus = "ACTIVE"
	ServiceStatusDeleteInProgress              ServiceStatus = "DELETE_IN_PROGRESS"
	ServiceStatusDeleteFailed                  ServiceStatus = "DELETE_FAILED"
	ServiceStatusUpdateInProgress              ServiceStatus = "UPDATE_IN_PROGRESS"
	ServiceStatusUpdateFailedCleanupInProgress ServiceStatus = "UPDATE_FAILED_CLEANUP_IN_PROGRESS"
	ServiceStatusUpdateFailedCleanupComplete   ServiceStatus = "UPDATE_FAILED_CLEANUP_COMPLETE"
	ServiceStatusUpdateFailedCleanupFailed     ServiceStatus = "UPDATE_FAILED_CLEANUP_FAILED"
	ServiceStatusUpdateFailed                  ServiceStatus = "UPDATE_FAILED"
	ServiceStatusUpdateCompleteCleanupFailed   ServiceStatus = "UPDATE_COMPLETE_CLEANUP_FAILED"
)

// Values returns all known values for ServiceStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ServiceStatus) Values() []ServiceStatus {
	return []ServiceStatus{
		"CREATE_IN_PROGRESS",
		"CREATE_FAILED_CLEANUP_IN_PROGRESS",
		"CREATE_FAILED_CLEANUP_COMPLETE",
		"CREATE_FAILED_CLEANUP_FAILED",
		"CREATE_FAILED",
		"ACTIVE",
		"DELETE_IN_PROGRESS",
		"DELETE_FAILED",
		"UPDATE_IN_PROGRESS",
		"UPDATE_FAILED_CLEANUP_IN_PROGRESS",
		"UPDATE_FAILED_CLEANUP_COMPLETE",
		"UPDATE_FAILED_CLEANUP_FAILED",
		"UPDATE_FAILED",
		"UPDATE_COMPLETE_CLEANUP_FAILED",
	}
}

type TemplateVersionStatus string

// Enum values for TemplateVersionStatus
const (
	TemplateVersionStatusRegistrationInProgress TemplateVersionStatus = "REGISTRATION_IN_PROGRESS"
	TemplateVersionStatusRegistrationFailed     TemplateVersionStatus = "REGISTRATION_FAILED"
	TemplateVersionStatusDraft                  TemplateVersionStatus = "DRAFT"
	TemplateVersionStatusPublished              TemplateVersionStatus = "PUBLISHED"
)

// Values returns all known values for TemplateVersionStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TemplateVersionStatus) Values() []TemplateVersionStatus {
	return []TemplateVersionStatus{
		"REGISTRATION_IN_PROGRESS",
		"REGISTRATION_FAILED",
		"DRAFT",
		"PUBLISHED",
	}
}