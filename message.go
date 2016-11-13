package gov

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type MessageConfig struct {
	Params map[string]string
	Rules  map[string]string
}

type MessageData struct {
	Name   string
	Values []string
}

func LoadMessages(filename string, config *MessageConfig) error {
	if buf, err := ioutil.ReadFile(filename); err != nil {
		return err
	} else {
		// orz
		m := make(map[interface{}]interface{})
		err = yaml.Unmarshal(buf, &m)

		// return &MessageConfig{Params: cast(m["Params"].(map[interface{}]interface{})), Rules: cast(m["Rules"].(map[interface{}]interface{}))}, nil
		config.Params = cast(m["Params"].(map[interface{}]interface{}))
		config.Rules = cast(m["Rules"].(map[interface{}]interface{}))
		return nil
	}
}

func cast(m map[interface{}]interface{}) map[string]string {
	m2 := make(map[string]string)
	for key, value := range m {
		m2[key.(string)] = value.(string)
	}
	return m2
}
