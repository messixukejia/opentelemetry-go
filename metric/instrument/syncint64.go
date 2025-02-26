// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package instrument // import "go.opentelemetry.io/otel/metric/instrument"

import (
	"go.opentelemetry.io/otel/metric/unit"
)

// Int64Config contains options for Synchronous instruments that record int64
// values.
type Int64Config struct {
	description string
	unit        unit.Unit
}

// NewInt64Config returns a new Int64Config with all opts
// applied.
func NewInt64Config(opts ...Int64Option) Int64Config {
	var config Int64Config
	for _, o := range opts {
		config = o.applyInt64(config)
	}
	return config
}

// Description returns the Config description.
func (c Int64Config) Description() string {
	return c.description
}

// Unit returns the Config unit.
func (c Int64Config) Unit() unit.Unit {
	return c.unit
}

// Int64Option applies options to synchronous int64 instruments.
type Int64Option interface {
	applyInt64(Int64Config) Int64Config
}
