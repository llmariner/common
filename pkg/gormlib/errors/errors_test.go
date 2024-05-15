package errors

import (
	"fmt"
	"testing"
)

func TestIsUniqueConstraintViolation(t *testing.T) {
	type testCase struct {
		err      error
		expected bool
	}
	tcs := []testCase{
		{
			err:      fmt.Errorf("UNIQUE constraint failed: orgs.uuid"),
			expected: true,
		},
		{
			err:      fmt.Errorf("unique constraint failed: orgs.uuid"),
			expected: true,
		},
		{
			err:      fmt.Errorf("some other error"),
			expected: false,
		},
		{
			err:      nil,
			expected: false,
		},
	}
	for i, tc := range tcs {
		t.Run(fmt.Sprintf("Test case %d", i), func(t *testing.T) {
			actual := IsUniqueConstraintViolation(tc.err)
			if actual != tc.expected {
				t.Errorf("expected %t, but got %t", tc.expected, actual)
			}
		})
	}
}
