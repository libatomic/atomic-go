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

const (
	OptionGetPath    = "/api/1.0.0/options/%s"
	OptionListPath   = "/api/1.0.0/options"
	OptionUpdatePath = "/api/1.0.0/options/%s"
	OptionRemovePath = "/api/1.0.0/options/%s"
)

func (c *Client) OptionGet(ctx context.Context, params *atomic.OptionGetInput) (*atomic.Option, error) {
	var resp ResponseProxy[atomic.Option]

	if err := params.Validate(); err != nil {
		return nil, err
	}

	path := fmt.Sprintf(OptionGetPath, params.Name)

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Get(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) OptionList(ctx context.Context, params *atomic.OptionListInput) ([]*atomic.Option, error) {
	var resp ResponseProxy[[]*atomic.Option]

	if err := params.Validate(); err != nil {
		return nil, err
	}

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, OptionListPath, params).Get(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Value(), nil
}

func (c *Client) OptionUpdate(ctx context.Context, params *atomic.OptionUpdateInput) (*atomic.Option, error) {
	var resp ResponseProxy[atomic.Option]

	if err := params.Validate(); err != nil {
		return nil, err
	}

	path := fmt.Sprintf(OptionUpdatePath, params.Name)

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Put(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) OptionRemove(ctx context.Context, params *atomic.OptionRemoveInput) error {
	if err := params.Validate(); err != nil {
		return err
	}

	path := fmt.Sprintf(OptionRemovePath, params.Name)

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Delete(),
		nil,
	); err != nil {
		return err
	}

	return nil
}
