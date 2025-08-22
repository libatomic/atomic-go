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
	Distribution            = atomic.Distribution
	DistributionCreateInput = atomic.DistributionCreateInput
	DistributionGetInput    = atomic.DistributionGetInput
	DistributionUpdateInput = atomic.DistributionUpdateInput
	DistributionDeleteInput = atomic.DistributionDeleteInput
	DistributionListInput   = atomic.DistributionListInput
)

const (
	DistributionGetPath    = "/api/1.0.0/distributions/%s"
	DistributionCreatePath = "/api/1.0.0/distributions"
	DistributionUpdatePath = "/api/1.0.0/distributions/%s"
	DistributionDeletePath = "/api/1.0.0/distributions/%s"
	DistributionListPath   = "/api/1.0.0/distributions"
)

func (c *Client) DistributionGet(ctx context.Context, params *DistributionGetInput) (*Distribution, error) {
	var resp ResponseProxy[Distribution]

	path := fmt.Sprintf(DistributionGetPath, params.DistributionID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Get(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) DistributionCreate(ctx context.Context, params *DistributionCreateInput) (*Distribution, error) {
	var resp ResponseProxy[Distribution]

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, DistributionCreatePath, params).Post(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) DistributionUpdate(ctx context.Context, params *DistributionUpdateInput) (*Distribution, error) {
	var resp ResponseProxy[Distribution]

	path := fmt.Sprintf(DistributionUpdatePath, params.DistributionID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Put(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) DistributionDelete(ctx context.Context, params *DistributionDeleteInput) error {
	path := fmt.Sprintf(DistributionDeletePath, params.DistributionID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Delete(),
		nil,
	); err != nil {
		return err
	}

	return nil
}

func (c *Client) DistributionList(ctx context.Context, params *DistributionListInput) ([]*Distribution, error) {
	var resp ResponseProxy[[]*Distribution]

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, DistributionListPath, params).Get(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Value(), nil
}
