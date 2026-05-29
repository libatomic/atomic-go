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
	Credit            = atomic.Credit
	CreditCreateInput = atomic.CreditCreateInput
	CreditGetInput    = atomic.CreditGetInput
	CreditUpdateInput = atomic.CreditUpdateInput
	CreditListInput   = atomic.CreditListInput
)

const (
	CreditGetPath    = "/api/1.0.0/credits/%s"
	CreditUpdatePath = "/api/1.0.0/credits/%s"
	CreditCreatePath = "/api/1.0.0/credits"
	CreditListPath   = "/api/1.0.0/credits"
)

func (c *Client) CreditGet(ctx context.Context, params *CreditGetInput) (*Credit, error) {
	var resp ResponseProxy[Credit]

	path := fmt.Sprintf(CreditGetPath, params.CreditID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Get(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) CreditUpdate(ctx context.Context, params *CreditUpdateInput) (*Credit, error) {
	var resp ResponseProxy[Credit]

	path := fmt.Sprintf(CreditUpdatePath, params.CreditID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Put(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) CreditCreate(ctx context.Context, params *CreditCreateInput) (*Credit, error) {
	var resp ResponseProxy[Credit]

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, CreditCreatePath, params).Post(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) CreditList(ctx context.Context, params *CreditListInput) ([]*Credit, error) {
	var resp ResponseProxy[[]*Credit]

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, CreditListPath, params).Get(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Value(), nil
}
