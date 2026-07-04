/*
 * This file is part of the Passport Atomic Stack (https://github.com/libatomic/atomic).
 * Copyright (c) 2026 Passport, Inc.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, version 3.
 *
 * This program is distributed in the hope that it will be useful, but
 * WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
 * General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program. If not, see <http://www.gnu.org/licenses/>.
 */

package atomic

import (
	"context"
	"fmt"

	"github.com/libatomic/atomic/pkg/atomic"
)

type (
	CustomRole            = atomic.CustomRole
	CustomRoleCreateInput = atomic.CustomRoleCreateInput
	CustomRoleGetInput    = atomic.CustomRoleGetInput
	CustomRoleUpdateInput = atomic.CustomRoleUpdateInput
	CustomRoleDeleteInput = atomic.CustomRoleDeleteInput
	CustomRoleListInput   = atomic.CustomRoleListInput
)

const (
	CustomRoleGetPath    = "/api/1.0.0/roles/%s"
	CustomRoleCreatePath = "/api/1.0.0/roles"
	CustomRoleUpdatePath = "/api/1.0.0/roles/%s"
	CustomRoleDeletePath = "/api/1.0.0/roles/%s"
	CustomRoleListPath   = "/api/1.0.0/roles"
)

func (c *Client) CustomRoleGet(ctx context.Context, params *CustomRoleGetInput) (*CustomRole, error) {
	var resp ResponseProxy[CustomRole]

	path := fmt.Sprintf(CustomRoleGetPath, params.RoleID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Get(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) CustomRoleCreate(ctx context.Context, params *CustomRoleCreateInput) (*CustomRole, error) {
	var resp ResponseProxy[CustomRole]

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, CustomRoleCreatePath, params).Post(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) CustomRoleUpdate(ctx context.Context, params *CustomRoleUpdateInput) (*CustomRole, error) {
	var resp ResponseProxy[CustomRole]

	path := fmt.Sprintf(CustomRoleUpdatePath, params.RoleID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Put(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) CustomRoleDelete(ctx context.Context, params *CustomRoleDeleteInput) error {
	path := fmt.Sprintf(CustomRoleDeletePath, params.RoleID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Delete(),
		nil,
	); err != nil {
		return err
	}

	return nil
}

func (c *Client) CustomRoleList(ctx context.Context, params *CustomRoleListInput) ([]*CustomRole, error) {
	var resp ResponseProxy[[]*CustomRole]

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, CustomRoleListPath, params).Get(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Value(), nil
}
