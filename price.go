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
	PriceGetPath    = "/api/1.0.0/prices/%s"
	PriceCreatePath = "/api/1.0.0/prices"
	PriceUpdatePath = "/api/1.0.0/prices/%s"
	PriceDeletePath = "/api/1.0.0/prices/%s"
	PriceListPath   = "/api/1.0.0/prices"
)

func (c *Client) PriceGet(ctx context.Context, params *atomic.PriceGetInput) (*atomic.Price, error) {
	var resp ResponseProxy[atomic.Price]

	if err := params.Validate(); err != nil {
		return nil, err
	}

	path := fmt.Sprintf(PriceGetPath, params.PriceID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Get(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) PriceCreate(ctx context.Context, params *atomic.PriceCreateInput) (*atomic.Price, error) {
	var resp ResponseProxy[atomic.Price]

	if err := params.Validate(); err != nil {
		return nil, err
	}

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, PriceCreatePath, params).Post(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) PriceUpdate(ctx context.Context, params *atomic.PriceUpdateInput) (*atomic.Price, error) {
	var resp ResponseProxy[atomic.Price]

	if err := params.Validate(); err != nil {
		return nil, err
	}

	path := fmt.Sprintf(PriceUpdatePath, params.PriceID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Put(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) PriceDelete(ctx context.Context, params *atomic.PriceDeleteInput) error {
	if err := params.Validate(); err != nil {
		return err
	}

	path := fmt.Sprintf(PriceDeletePath, params.PriceID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Delete(),
		nil,
	); err != nil {
		return err
	}

	return nil
}

func (c *Client) PriceList(ctx context.Context, params *atomic.PriceListInput) ([]*atomic.Price, error) {
	var resp ResponseProxy[[]*atomic.Price]

	if err := params.Validate(); err != nil {
		return nil, err
	}

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, PriceListPath, params).Get(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Value(), nil
}
