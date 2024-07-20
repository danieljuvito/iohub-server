package util

import "testing"

func TestValidateEmail(t *testing.T) {
    // Test cases for email validation
    testCases := []struct {
        email    string
        expected bool
    }{
        {"user@example.com", true},
        {"invalid-email", false},
        {"another@example", false},
    }

    for _, tc := range testCases {
        result := ValidateEmail(tc.email)
        if result != tc.expected {
            t.Errorf("ValidateEmail(%s) = %v; expected %v", tc.email, result, tc.expected)
        }
    }
}
