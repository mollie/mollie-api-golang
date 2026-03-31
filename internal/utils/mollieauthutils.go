package utils

import (
	"context"
	"reflect"
	"strings"

	"github.com/mollie/mollie-api-golang/internal/globals"
)

// CanHaveGlobalFields returns true if the security function resolves to an OAuth
// access token (starts with "access_"), meaning profileId and testmode can be injected.
// It accepts the raw security function to avoid import cycles with internal/config
// and models/components.
func CanHaveGlobalFields(securityFn func(context.Context) (interface{}, error)) bool {
	if securityFn == nil {
		return false
	}
	val, err := securityFn(context.Background())
	if err != nil || val == nil {
		return false
	}
	token := extractSecurityToken(val)
	return token != nil && strings.HasPrefix(*token, "access_")
}

// HasGlobalFields returns true when at least one global field (profileID or testmode) is set.
func HasGlobalFields(g globals.Globals) bool {
	return g.ProfileID != nil || g.Testmode != nil
}

// extractSecurityToken retrieves the first non-nil string pointer from an APIKey or OAuth
// field on the security value using reflection, avoiding a direct dependency on
// models/components.
func extractSecurityToken(val interface{}) *string {
	v := reflect.ValueOf(val)
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return nil
		}
		v = v.Elem()
	}
	for _, fieldName := range []string{"APIKey", "OrganizationAccessToken", "OAuth"} {
		field := v.FieldByName(fieldName)
		if field.IsValid() && field.Kind() == reflect.Ptr && !field.IsNil() {
			s := field.Elem().String()
			return &s
		}
	}
	return nil
}
