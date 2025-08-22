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
	Price            = atomic.Price
	PriceGetInput    = atomic.PriceGetInput
	PriceCreateInput = atomic.PriceCreateInput
	PriceUpdateInput = atomic.PriceUpdateInput
	PriceDeleteInput = atomic.PriceDeleteInput
	PriceListInput   = atomic.PriceListInput
)

const (
	PriceGetPath    = "/api/1.0.0/prices/%s"
	PriceCreatePath = "/api/1.0.0/prices"
	PriceUpdatePath = "/api/1.0.0/prices/%s"
	PriceDeletePath = "/api/1.0.0/prices/%s"
	PriceListPath   = "/api/1.0.0/prices"
)

func (c *Client) PriceGet(ctx context.Context, params *PriceGetInput) (*Price, error) {
	var resp ResponseProxy[Price]

	path := fmt.Sprintf(PriceGetPath, params.PriceID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Get(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) PriceCreate(ctx context.Context, params *PriceCreateInput) (*Price, error) {
	var resp ResponseProxy[Price]

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, PriceCreatePath, params).Post(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) PriceUpdate(ctx context.Context, params *PriceUpdateInput) (*Price, error) {
	var resp ResponseProxy[Price]

	path := fmt.Sprintf(PriceUpdatePath, params.PriceID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Put(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) PriceDelete(ctx context.Context, params *PriceDeleteInput) error {
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

func (c *Client) PriceList(ctx context.Context, params *PriceListInput) ([]*Price, error) {
	var resp ResponseProxy[[]*Price]

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
