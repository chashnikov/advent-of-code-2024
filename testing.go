package main

import "testing"

func AssertEquals[V comparable](t *testing.T, expected, got V) {
	t.Helper()

	if expected != got {
		t.Errorf(`got:
%v
,
expected:
%v
`, got, expected)
	}
}
