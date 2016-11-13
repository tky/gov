package rules_test

import (
	"gov/rules"
	"testing"
)

func toPtr(s string) *string {
	return &s
}

func TestMinLength(t *testing.T) {
	if err := rules.MinLength("a", []string{"1"}); err != nil {
		t.Error("Should not reutrn error")
	}
	if err := rules.MinLength(toPtr("a"), []string{"1"}); err != nil {
		t.Error("Should not reutrn error")
	}

	if err := rules.MinLength("a", []string{"2"}); err == nil {
		t.Error("Should reutrn error")
	}

	if err := rules.MinLength(toPtr("a"), []string{"2"}); err == nil {
		t.Error("Should reutrn error")
	}

	if err := rules.MinLength(nil, []string{"1"}); err != nil {
		t.Error("Should not reutrn error if target is nil")
	}
}

func TestRequired(t *testing.T) {
	if err := rules.Required("a", []string{}); err != nil {
		t.Error("Should not reutrn error")
	}

	if err := rules.Required(nil, []string{}); err == nil {
		t.Error("Should reutrn error")
	} else {
		if err != rules.ErrValidation {
			t.Error("Should reutrn Validation Error")
		}
	}

	if err := rules.Required("a", []string{"1"}); err == nil {
		t.Error("Should reutrn error if requires has parmas")
	} else if err != rules.ErrIllegalParamsOnRequired {
		t.Error("Should reutrn error if requires has parmas")
	}

	var v *string
	if err := rules.Required(v, []string{}); err == nil {
		t.Error("Should return errir if v is nil string pointer")
	}

	v = toPtr("a")
	if err := rules.Required(v, []string{}); err != nil {
		t.Error("Should not return errir if v is string pointer")
	}

	v = toPtr("")
	if err := rules.Required(v, []string{}); err == nil {
		t.Error("Should return errir if v is empty string pointer")
	}
}
