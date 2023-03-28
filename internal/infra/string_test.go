package infra

import (
	"testing"
)

func Test_SnakeCase(t *testing.T) {
	if got, want := SnakeCase("PabB2BicNotify"), "pab_b2bic_notify"; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Test_CamelCase(t *testing.T) {
	if got, want := CamelCase("pab_b2bic_notify"), "PabB2BicNotify"; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
