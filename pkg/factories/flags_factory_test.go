// Copyright © 2020 The Knative Authors
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

func TestNewDefaultFlagsFactory(t *testing.T) {
	flagsFactory := NewDefaultFlagsFactory(NewDefaultKnSourceFactory())

	assert.Assert(t, flagsFactory != nil)
}

func TestFlagsFactory_KnSourceFactory(t *testing.T) {
	knSourceFactory := NewDefaultKnSourceFactory()
	flagsFactory := NewDefaultFlagsFactory(knSourceFactory)

	assert.Equal(t, flagsFactory.KnSourceFactory(), knSourceFactory)
}

func TestCreateFlags(t *testing.T) {
	flagsFactory := NewDefaultFlagsFactory(NewDefaultKnSourceFactory())

	assert.Assert(t, flagsFactory.CreateFlags() != nil)
}

func TestDeleteFlags(t *testing.T) {
	flagsFactory := NewDefaultFlagsFactory(NewDefaultKnSourceFactory())

	assert.Assert(t, flagsFactory.DeleteFlags() != nil)
}

func TestUpdateFlags(t *testing.T) {
	flagsFactory := NewDefaultFlagsFactory(NewDefaultKnSourceFactory())

	assert.Assert(t, flagsFactory.UpdateFlags() != nil)
}

func TestDescribeFlags(t *testing.T) {
	flagsFactory := NewDefaultFlagsFactory(NewDefaultKnSourceFactory())

	assert.Assert(t, flagsFactory.DescribeFlags() != nil)
}
