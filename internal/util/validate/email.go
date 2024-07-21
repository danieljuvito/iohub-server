package validate

import (
    "regexp"
)

// Email checks if the provided email address is valid.
func Email(email string) bool {
    // A simple regex pattern for email validation (you can use a more comprehensive one).
    // This pattern allows for alphanumeric characters, dots, hyphens, and underscores.
    emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
    match, _ := regexp.MatchString(emailPattern, email)
    return match
}
