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

package core

import (
	"github.com/maximilien/kn-source-pkg/pkg/commands"
	"github.com/maximilien/kn-source-pkg/pkg/commands/source"

	"github.com/spf13/cobra"
)

type DefautSourceCommandFactory struct {
}

func NewDefaultSourceCommandFactory() commands.SourceCommandFactory {
	return &DefautSourceCommandFactory{}
}

func (f *DefautSourceCommandFactory) CreateCommand(params *commands.KnSourceParams) *cobra.Command {
	return source.NewSourceCreateCommand(params)
}

func (f *DefautSourceCommandFactory) DeleteCommand(params *commands.KnSourceParams) *cobra.Command {
	return source.NewSourceDeleteCommand(params)
}

func (f *DefautSourceCommandFactory) UpdateCommand(params *commands.KnSourceParams) *cobra.Command {
	return source.NewSourceUpdateCommand(params)
}

func (f *DefautSourceCommandFactory) DescribeCommand(params *commands.KnSourceParams) *cobra.Command {
	return source.NewSourceDescribeCommand(params)
}
