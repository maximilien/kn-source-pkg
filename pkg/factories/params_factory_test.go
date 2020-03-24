// Copyright © 2018 The Knative Authors
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

package factories

import (
	"testing"

	"gotest.tools/assert"
)

func TestNewDefaultParamsFactory(t *testing.T) {
	paramsFactory := NewDefaultParamsFactory()

	assert.Assert(t, paramsFactory != nil)
}

func TestParamsFactory_KnSourceParams(t *testing.T) {
	paramsFactory := NewDefaultParamsFactory()

	assert.Assert(t, paramsFactory.KnSourceParams() == nil)
	knSourceParams := paramsFactory.CreateKnSourceParams()
	assert.Equal(t, paramsFactory.KnSourceParams(), knSourceParams)
}

func TestCreateKnSourceParams(t *testing.T) {
	paramsFactory := NewDefaultParamsFactory()

	knSourceParams := paramsFactory.CreateKnSourceParams()
	assert.Assert(t, knSourceParams != nil)
}
