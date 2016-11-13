package gov_test

import (
	"gov"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Sample struct {
	Name    string `gov:"required,max:10"`
	Version string `gov:"numeric"`
	Index   int    `gov:"min:3"`
}

type Sample1 struct {
	Name *string `gov:"required"`
}

func TestValidate(t *testing.T) {
	messages := gov.Validate(Sample1{Name: nil})
	assert.Equal(t, 1, len(messages))
	assert.Equal(t, "名前は必須です", (messages[0]))
}

func TestParseParam(t *testing.T) {
	s := gov.Meta{Value: "abc", Tag: "required"}
	if params, err := s.ParseParam(); err == nil {
		assert.Equal(t, 1, len(params))
		assert.Equal(t, "required", params[0].Key)
		assert.Equal(t, 0, len(params[0].Values))
	} else {
		t.Error("Should not return error")
	}
}

func TestParseParamWithMultiValue(t *testing.T) {
	s := gov.Meta{Value: "abc", Tag: "required,maxlength:10"}
	if params, err := s.ParseParam(); err == nil {
		assert.Equal(t, 2, len(params))
		assert.Equal(t, "required", params[0].Key)
		assert.Equal(t, 0, len(params[0].Values))

		assert.Equal(t, "maxlength", params[1].Key)
		assert.Equal(t, 1, len(params[1].Values))
		assert.Equal(t, "10", params[1].Values[0])
	} else {
		t.Error("Should not return error")
	}
}

func TestParser(t *testing.T) {
	rs := gov.Parser(Sample{Name: "name", Version: "1.0"})

	assert.Equal(t, 3, len(rs))
	assert.Equal(t, "required,max:10", rs[0].Tag)
	assert.Equal(t, "name", rs[0].Value)
	assert.Equal(t, "Name", rs[0].FieldName)

	assert.Equal(t, "numeric", rs[1].Tag)
	assert.Equal(t, "1.0", rs[1].Value)
	assert.Equal(t, "Version", rs[1].FieldName)
}
