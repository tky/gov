package gov_test

import (
	"fmt"
	"gov"
	"testing"
)

func TestLoadMessage(t *testing.T) {
	var config gov.MessageConfig
	if err := gov.LoadMessages("validation.yml", &config); err != nil {
		t.Error("Should not return error", err)
	} else {
		fmt.Println(config)
	}
}
