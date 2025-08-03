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
	"net/http"

	"github.com/libatomic/atomic/pkg/atomic"
)

const (
	ApplicationGetPath    = "/api/1.0.0/applications/%s"
	ApplicationCreatePath = "/api/1.0.0/applications"
	ApplicationUpdatePath = "/api/1.0.0/applications/%s"
	ApplicationDeletePath = "/api/1.0.0/applications/%s"
	ApplicationListPath   = "/api/1.0.0/applications"
)

func (c *Client) ApplicationCreate(ctx context.Context, params *atomic.ApplicationCreateInput) (*atomic.Application, error) {
	var resp ResponseProxy[atomic.Application]

	if err := params.Validate(); err != nil {
		return nil, err
	}

	if err := c.Backend.ExecContext(
		ctx,
		http.MethodPost,
		ApplicationCreatePath,
		&ParamsProxy[atomic.ApplicationCreateInput]{
			methodParams:  *params,
			requestParams: ParamsFromContext(ctx),
		}, &resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) ApplicationGet(ctx context.Context, params *atomic.ApplicationGetInput) (*atomic.Application, error) {
	var resp ResponseProxy[atomic.Application]

	if err := params.Validate(); err != nil {
		return nil, err
	}

	path := fmt.Sprintf(ApplicationGetPath, params.ApplicationID.String())

	if err := c.Backend.ExecContext(
		ctx,
		http.MethodGet,
		path,
		&ParamsProxy[atomic.ApplicationGetInput]{
			methodParams:  *params,
			requestParams: ParamsFromContext(ctx),
		}, &resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) ApplicationUpdate(ctx context.Context, params *atomic.ApplicationUpdateInput) (*atomic.Application, error) {
	var resp ResponseProxy[atomic.Application]

	if err := params.Validate(); err != nil {
		return nil, err
	}

	path := fmt.Sprintf(ApplicationUpdatePath, params.ApplicationID.String())

	if err := c.Backend.ExecContext(
		ctx,
		http.MethodPut,
		path,
		&ParamsProxy[atomic.ApplicationUpdateInput]{
			methodParams:  *params,
			requestParams: ParamsFromContext(ctx),
		}, &resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) ApplicationDelete(ctx context.Context, params *atomic.ApplicationDeleteInput) error {
	if err := params.Validate(); err != nil {
		return err
	}

	path := fmt.Sprintf(ApplicationDeletePath, params.ApplicationID.String())

	return c.Backend.ExecContext(ctx, http.MethodDelete, path, &ParamsProxy[atomic.ApplicationDeleteInput]{
		methodParams:  *params,
		requestParams: ParamsFromContext(ctx),
	}, nil)
}

func (c *Client) ApplicationList(ctx context.Context, params *atomic.ApplicationListInput) ([]*atomic.Application, error) {
	var resp ResponseProxy[[]*atomic.Application]

	if err := c.Backend.ExecContext(
		ctx,
		http.MethodGet,
		ApplicationListPath,
		&ParamsProxy[atomic.ApplicationListInput]{
			methodParams:  *params,
			requestParams: ParamsFromContext(ctx),
		}, &resp); err != nil {
		return nil, err
	}

	return resp.Value(), nil
}
