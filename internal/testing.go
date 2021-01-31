package internal

import "testing"

// PrintTestDiff prints the difference between expected and actual.
// This function depends on github.com/google/go-cmp.
// As using this function, you must implement comparison like cmp.Diff(want, go).
func PrintTestDiff(t *testing.T, diff string) {
	t.Errorf("Differs: (-want +got)\n%s", diff)
}
