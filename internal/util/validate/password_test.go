package validate

import "testing"

func TestPassword(t *testing.T) {
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
        result := Password(tc.password)
        if result != tc.expected {
            t.Errorf("Password(%s) = %v; expected %v", tc.password, result, tc.expected)
        }
    }
}
