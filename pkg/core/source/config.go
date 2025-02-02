/*
Copyright 2021 Loggie Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package source

import (
	"github.com/loggie-io/loggie/pkg/core/cfg"
	"github.com/pkg/errors"
)

var (
	ErrSourceNameRequired = errors.New("source name is required")
)

type Config struct {
	cfg.ComponentBaseConfig `yaml:",inline"`
	FieldsUnderRoot         bool                   `yaml:"fieldsUnderRoot,omitempty" default:"false"`
	FieldsUnderKey          string                 `yaml:"fieldsUnderKey,omitempty" default:"fields"`
	Fields                  map[string]interface{} `yaml:"fields,omitempty"`
	FieldsFromEnv           map[string]string      `yaml:"fieldsFromEnv,omitempty"`
}

func (c *Config) Validate() error {
	if c.Name == "" {
		return ErrSourceNameRequired
	}
	return nil
}
