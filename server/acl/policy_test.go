package acl

import (
	"context"
	"testing"
)

func TestACL(t *testing.T) {
	var email string
	var scopes []string

	t.Run("Test whoami", func(t *testing.T) {
		email = WhoAmI(context.Background())
	})

	t.Run("Test GetScope ", func(t *testing.T) {
		scopes = GetScope(email, "GET", "workspace")
	})
}
