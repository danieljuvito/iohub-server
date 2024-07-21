package validate

import "testing"

func TestEmail(t *testing.T) {
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
        result := Email(tc.email)
        if result != tc.expected {
            t.Errorf("Email(%s) = %v; expected %v", tc.email, result, tc.expected)
        }
    }
}
