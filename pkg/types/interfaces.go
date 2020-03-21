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

package types

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type RunE = func(cmd *cobra.Command, args []string) error

type KnSourceClient interface {
	KnSourceParams() *KnSourceParams
}

type ParamsFactory interface {
	Create() *KnSourceParams
}

type ClientFactory interface {
	Create() KnSourceClient

	KnSourceParams() *KnSourceParams
}

type CommandFactory interface {
	SourceCommand() *cobra.Command

	CreateCommand() *cobra.Command
	DeleteCommand() *cobra.Command
	UpdateCommand() *cobra.Command
	DescribeCommand() *cobra.Command

	KnSourceParams() *KnSourceParams
}

type FlagsFactory interface {
	CreateFlags() *pflag.FlagSet
	DeleteFlags() *pflag.FlagSet
	UpdateFlags() *pflag.FlagSet
	DescribeFlags() *pflag.FlagSet

	KnSourceParams() *KnSourceParams
}

type RunEFactory interface {
	CreateRunE() RunE
	DeleteRunE() RunE
	UpdateRunE() RunE
	DescribeRunE() RunE

	KnSourceClient() KnSourceClient
}
