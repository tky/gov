package rules_test

import (
	"gov/rules"
	"testing"
)

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
}
