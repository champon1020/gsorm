package internal

import "testing"

// PrintTestDiff prints the difference between expected and actual.
func PrintTestDiff(t *testing.T, diff string) {
	t.Errorf("Differs: (-got +want)\n%s", diff)
}
