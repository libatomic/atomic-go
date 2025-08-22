/*
 * This file is part of the Passport Atomic Stack (https://github.com/libatomic/atomic).
 * Copyright (c) 2024 Atomic Publishing.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, version 3.
 *
 * This program is distributed in the hope that it will be useful, but
 * WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
 * General Public License for more detail
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
	Instance            = atomic.Instance
	InstanceCreateInput = atomic.InstanceCreateInput
	InstanceGetInput    = atomic.InstanceGetInput
	InstanceListInput   = atomic.InstanceListInput
	InstanceUpdateInput = atomic.InstanceUpdateInput
	InstanceDeleteInput = atomic.InstanceDeleteInput
)

const (
	InstanceCreatePath = "/api/1.0.0/instances"
	InstanceGetPath    = "/api/1.0.0/instances/%s"
	InstanceListPath   = "/api/1.0.0/instances"
	InstanceUpdatePath = "/api/1.0.0/instances/%s"
	InstanceDeletePath = "/api/1.0.0/instances/%s"
)

func (c *Client) InstanceCreate(ctx context.Context, params *InstanceCreateInput) (*Instance, error) {
	var resp ResponseProxy[Instance]

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, InstanceCreatePath, params).Post(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) InstanceGet(ctx context.Context, params *InstanceGetInput) (*Instance, error) {
	var resp ResponseProxy[Instance]

	if params.InstanceID == nil {
		return nil, fmt.Errorf("instance_id is required")
	}

	path := fmt.Sprintf(InstanceGetPath, params.InstanceID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Get(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) InstanceList(ctx context.Context, params *InstanceListInput) ([]*Instance, error) {
	var resp ResponseProxy[[]*Instance]

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, InstanceListPath, params).Get(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Value(), nil
}

func (c *Client) InstanceUpdate(ctx context.Context, params *InstanceUpdateInput) (*Instance, error) {
	var resp ResponseProxy[Instance]

	path := fmt.Sprintf(InstanceUpdatePath, params.InstanceID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Put(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) InstanceDelete(ctx context.Context, params *InstanceDeleteInput) error {
	path := fmt.Sprintf(InstanceDeletePath, params.InstanceID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Delete(),
		nil,
	); err != nil {
		return err
	}

	return nil
}
