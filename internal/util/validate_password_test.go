package util

import "testing"

func TestValidatePassword(t *testing.T) {
    // Test cases for password validation
    testCases := []struct {
        password string
        expected bool
    }{
        {"securepassword", false}, // Too short
        {"StrongP@ssw0rd", true},  // Meets requirements
        {"12345678", false},       // Numeric only
    }

    for _, tc := range testCases {
        result := ValidatePassword(tc.password)
        if result != tc.expected {
            t.Errorf("ValidatePassword(%s) = %v; expected %v", tc.password, result, tc.expected)
        }
    }
}
