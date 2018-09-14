package acl

const (
	WorkspaceCapabilityDeny   = "deny"
	WorkspaceCapabilityList   = "list"
	WorkspaceCapabilityCreate = "create"
	WorkspaceCapabilityGet    = "get"

	LayoutCapabilityCreate  = "create"
	LayoutCapabilityGet     = "get"
	LayoutCapabilityApply   = "apply"
	LayoutCapabilityDestroy = "destroy"
)

// per user.
/*
1. workspace: list of id
2. layout: list of id.

if i have an array of policies, it will be for multiple users.
*/
type Policy struct {
	workspace *WorkspacePolicy
	layout    *LayoutPolicy
}

// for a workspace.
// operation maps to the exact command which is fired.
type WorkspacePolicy struct {
	capabilities []*Capability
}

// for a layout.
// operation maps to the exact command which is fired.
type LayoutPolicy struct {
	capabilities []*Capability
}

type Capability struct {
	operation string
	scope     []string
}
