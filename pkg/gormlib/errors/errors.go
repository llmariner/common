package errors

import "strings"

// IsUniqueConstraintViolation returns true if a given error is caused by unique constraint violation.
func IsUniqueConstraintViolation(err error) bool {
	if err == nil {
		return false
	}
	s := err.Error()
	// Check the error message. One for sqlite and the other for Posgres.
	return strings.Contains(s, "UNIQUE constraint") || strings.Contains(s, "unique constraint")
}

// IsDeadlock returns true if a given error is caused by deadlock.
func IsDeadlock(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), "deadlock detected")
}
