package acl

import "context"

type ACL struct {
	administrator bool
	policies      *Policy
}

func NewACL(administrator bool, policies *Policy) *ACL {
	return &ACL{administrator: administrator, policies: policies}
}

// gets the email as per authentication workflow.
func WhoAmI(ctx context.Context) string {
	return "talina@trustingsocial.com"
}

// define acl for this whoami user. for this i need to know the policies for this user, which will be from a json?

// gets the scope for the object wrt the operation and user.
func GetScope(email, operation, object string) []string {
	return []string{"w1"}
}

func Check(id string, ids []string) bool {
	return true
}
