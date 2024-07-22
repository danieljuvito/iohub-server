package model

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestUser_Validate(t *testing.T) {
    // Test cases for User validation
    testCases := []struct {
        user    User
        isError bool
    }{
        {
            User{ID: 1, Email: "user@example.com", Password: "StrongP@ssw0rd"},
            false, // Valid user
        },
        {
            User{ID: 2, Email: "invalid-email", Password: "weak"},
            true, // Invalid email
        },
        {
            User{ID: 3, Email: "user@example.com", Password: "short"},
            true, // Invalid password
        },
    }

    for _, tc := range testCases {
        err := tc.user.Validate()
        if tc.isError {
            assert.Error(t, err)
        } else {
            assert.NoError(t, err)
        }
    }
}
