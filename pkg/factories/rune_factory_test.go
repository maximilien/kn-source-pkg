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

	"github.com/maximilien/kn-source-pkg/pkg/types"
	"github.com/spf13/cobra"
)

func TestNewDefaultRunEFactory(t *testing.T) {
	runEFactory := createDefaultRunEFactory()

	assert.Assert(t, runEFactory != nil)
}

func TestRunEFactory_KnSourceParams(t *testing.T) {
	runEFactory := createDefaultRunEFactory()

	assert.Assert(t, runEFactory.KnSourceFactory().KnSourceParams() != nil)
}

func KnSourceClientFactory(t *testing.T) {
	runEFactory := createDefaultRunEFactory()

	assert.Assert(t, runEFactory.KnSourceFactory() != nil)
}

func KnSourceClient(t *testing.T) {
	runEFactory := createDefaultRunEFactory()

	knSourceClient, err := runEFactory.KnSourceClient(&cobra.Command{})
	assert.NilError(t, err)
	assert.Assert(t, knSourceClient, nil)
}

func TestCreateRunE(t *testing.T) {
	runEFactory := createDefaultRunEFactory()

	assert.Assert(t, runEFactory.CreateRunE() != nil)
}

func TestDeleteRunE(t *testing.T) {
	runEFactory := createDefaultRunEFactory()

	assert.Assert(t, runEFactory.DeleteRunE() != nil)
}

func TestUpdateRunE(t *testing.T) {
	runEFactory := createDefaultRunEFactory()

	assert.Assert(t, runEFactory.UpdateRunE() != nil)
}

func TestDescribeRunE(t *testing.T) {
	runEFactory := createDefaultRunEFactory()

	assert.Assert(t, runEFactory.DescribeRunE() != nil)
}

// Private

func createDefaultRunEFactory() types.RunEFactory {
	knSourceFactory := NewDefaultKnSourceFactory()
	return NewDefaultRunEFactory(knSourceFactory)
}
