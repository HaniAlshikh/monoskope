// Copyright 2021 Monoskope Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package commands

import (
	"context"

	"github.com/finleap-connect/monoskope/pkg/domain/constants/aggregates"
	"github.com/finleap-connect/monoskope/pkg/domain/constants/commands"
	"github.com/finleap-connect/monoskope/pkg/domain/constants/roles"
	"github.com/finleap-connect/monoskope/pkg/domain/constants/scopes"
	es "github.com/finleap-connect/monoskope/pkg/eventsourcing"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/anypb"
)

func init() {
	es.DefaultCommandRegistry.RegisterCommand(NewDeleteTenantClusterBindingCommand)
}

// DeleteTenantClusterBindingCommand is a command for deleting a TenantClusterBinding.
type DeleteTenantClusterBindingCommand struct {
	*es.BaseCommand
}

// NewDeleteTenantClusterBindingCommand creates a DeleteTenantClusterBindingCommand.
func NewDeleteTenantClusterBindingCommand(id uuid.UUID) es.Command {
	return &DeleteTenantClusterBindingCommand{
		BaseCommand: es.NewBaseCommand(id, aggregates.TenantClusterBinding, commands.DeleteTenantClusterBinding),
	}
}

func (c *DeleteTenantClusterBindingCommand) SetData(a *anypb.Any) error {
	return nil
}

// Policies returns the Role/Scope/Resource combination allowed to execute.
func (c *DeleteTenantClusterBindingCommand) Policies(ctx context.Context) []es.Policy {
	return []es.Policy{
		es.NewPolicy().WithRole(roles.Admin).WithScope(scopes.System), // Allows system admins
	}
}