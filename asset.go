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
	AssetCreatePath = "/api/1.0.0/assets"
	AssetGetPath    = "/api/1.0.0/assets/%s"
	AssetUpdatePath = "/api/1.0.0/assets/%s"
	AssetDeletePath = "/api/1.0.0/assets/%s"
	AssetListPath   = "/api/1.0.0/assets"
)

func (c *Client) AssetCreate(ctx context.Context, params *atomic.AssetCreateInput) (*atomic.Asset, error) {
	var resp ResponseProxy[atomic.Asset]

	if err := params.Validate(); err != nil {
		return nil, err
	}

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, AssetCreatePath, params).Post(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) AssetGet(ctx context.Context, params *atomic.AssetGetInput) (*atomic.Asset, error) {
	var resp ResponseProxy[atomic.Asset]

	if err := params.Validate(); err != nil {
		return nil, err
	}

	path := fmt.Sprintf(AssetGetPath, params.AssetID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Get(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) AssetUpdate(ctx context.Context, params *atomic.AssetUpdateInput) (*atomic.Asset, error) {
	var resp ResponseProxy[atomic.Asset]

	if err := params.Validate(); err != nil {
		return nil, err
	}

	path := fmt.Sprintf(AssetUpdatePath, params.AssetID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Put(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) AssetDelete(ctx context.Context, params *atomic.AssetDeleteInput) error {
	if err := params.Validate(); err != nil {
		return err
	}

	path := fmt.Sprintf(AssetDeletePath, params.AssetID.String())

	return c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Delete(),
		nil,
	)
}

func (c *Client) AssetList(ctx context.Context, params *atomic.AssetListInput) ([]*atomic.Asset, error) {
	var resp ResponseProxy[[]*atomic.Asset]

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, AssetListPath, params).Get(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Value(), nil
}
