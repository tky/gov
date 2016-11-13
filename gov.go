package gov

import (
	"bytes"
	"fmt"
	"gov/rules"
	"html/template"
	"reflect"
	"strings"
)

type Param struct {
	Key    string
	Values []string
}

type Meta struct {
	FieldName string
	Value     interface{}
	Tag       string
}

func (m Meta) ParseParam() ([]Param, error) {
	tags := strings.Split(m.Tag, ",")
	var params []Param
	for _, t := range tags {
		if p, err := parseParam(t); err != nil {
			return nil, err
		} else {
			params = append(params, p)
		}
	}
	return params, nil
}

var ruleMap map[string]rules.Validate
var messageConfig MessageConfig

func init() {
	ruleMap = map[string]rules.Validate{}
	ruleMap["required"] = rules.Required
	ruleMap["min-length"] = rules.MinLength
	if err := LoadMessages("validation.yml", &messageConfig); err != nil {
		panic("erro loading validation.yml")
	}
}

// Validate validate target object and return messages if validation errors occoured.
func Validate(target interface{}) []string {
	var messages []string
	ms := Parser(target) // []Meta
	for _, m := range ms {
		if ps, err := m.ParseParam(); err != nil {
			panic("Illegal tag")
		} else {
			for _, p := range ps {
				if rule, ok := ruleMap[p.Key]; !ok {
					panic("Missing key:" + p.Key)
				} else {
					if err := rule(m.Value, p.Values); err != nil {
						var name string
						if v, ok := messageConfig.Params[m.FieldName]; ok {
							name = v
						} else {
							name = m.FieldName
						}

						if tl, ok := messageConfig.Rules[p.Key]; !ok {
							panic("Missiong rules")
						} else {
							tmpl, err := template.New(p.Key).Parse(tl)
							// TODO: error handling
							fmt.Println(err)
							var doc bytes.Buffer
							tmpl.Execute(&doc, MessageData{Name: name, Values: p.Values})
							s := doc.String()
							messages = append(messages, s)
						}
					}
				}
			}
		}

	}
	return messages
}

func parseParam(v string) (Param, error) {
	vs := strings.Split(v, ":")
	return Param{
		Key:    vs[0],
		Values: vs[1:],
	}, nil
}

func Parser(st interface{}) []Meta {
	rt, rv := reflect.TypeOf(st), reflect.ValueOf(st)

	metas := make([]Meta, rt.NumField())

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		gov := field.Tag.Get("gov")
		value := rv.Field(i).Interface()

		metas[i] = Meta{FieldName: field.Name, Value: value, Tag: gov}
	}
	return metas
}
