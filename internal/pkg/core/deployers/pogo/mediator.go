/*
Copyright 2025 The Synapse Authors
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

package pogo

import (
	"encoding/xml"

	"github.com/apache/synapse-go/internal/pkg/core/artifacts"
)

type Mediator interface {
	Unmarshal(d *xml.Decoder, start xml.StartElement, position artifacts.Position) (artifacts.Mediator, error)
}
