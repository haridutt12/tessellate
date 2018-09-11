package acl

import "context"

type ACL struct {
	// user: {object: operation: array of scopes.}
	scope map[string]map[string]map[string][]string
}

// gets the email as per authentication workflow.
func WhoAmI(ctx context.Context) string {
	return "talina@trustingsocial.com"
}

// gets the scope for the object wrt the operation and user.
func GetScope(email, operation, object string) []string {
	return []string{"w1"}
}

func Check(id string, ids []string) bool {
	return true
}
