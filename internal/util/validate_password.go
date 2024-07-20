package util

import (
    "strings"
    "unicode"
)

// ValidatePassword checks if the provided password meets certain criteria (e.g., length).
func ValidatePassword(password string) bool {
    // Minimum length requirement
    minPasswordLength := 8
    if len(password) < minPasswordLength {
        return false
    }

    // Complexity checks
    hasUppercase := false
    hasLowercase := false
    hasDigit := false
    hasSpecialChar := false

    for _, char := range password {
        switch {
        case unicode.IsUpper(char):
            hasUppercase = true
        case unicode.IsLower(char):
            hasLowercase = true
        case unicode.IsDigit(char):
            hasDigit = true
        case strings.ContainsAny(string(char), "!@#$%^&*()_+{}[]:;<>,.?"):
            hasSpecialChar = true
        }
    }

    // Avoid common patterns
    blacklist := []string{"password", "123456", "qwerty"}
    for _, common := range blacklist {
        if strings.Contains(password, common) {
            return false
        }
    }

    // No sequential or repeating characters
    for i := 0; i < len(password)-2; i++ {
        if password[i] == password[i+1] && password[i] == password[i+2] {
            return false
        }
    }

    // All checks passed
    return hasUppercase && hasLowercase && hasDigit && hasSpecialChar
}
