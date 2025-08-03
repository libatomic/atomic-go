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
	AudienceGetPath    = "/api/1.0.0/audiences/%s"
	AudienceCreatePath = "/api/1.0.0/audiences"
	AudienceUpdatePath = "/api/1.0.0/audiences/%s"
	AudienceDeletePath = "/api/1.0.0/audiences/%s"
	AudienceListPath   = "/api/1.0.0/audiences"
)

func (c *Client) AudienceGet(ctx context.Context, params *atomic.AudienceGetInput) (*atomic.Audience, error) {
	var resp ResponseProxy[atomic.Audience]

	if err := params.Validate(); err != nil {
		return nil, err
	}

	path := fmt.Sprintf(AudienceGetPath, params.AudienceID.String())

	if err := c.Backend.ExecContext(
		ctx,
		http.MethodGet,
		path,
		&ParamsProxy[atomic.AudienceGetInput]{
			methodParams:  *params,
			requestParams: ParamsFromContext(ctx),
		}, &resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) AudienceCreate(ctx context.Context, params *atomic.AudienceCreateInput) (*atomic.Audience, error) {
	var resp ResponseProxy[atomic.Audience]

	if err := params.Validate(); err != nil {
		return nil, err
	}

	if err := c.Backend.ExecContext(
		ctx,
		http.MethodPost,
		AudienceCreatePath,
		&ParamsProxy[atomic.AudienceCreateInput]{
			methodParams:  *params,
			requestParams: ParamsFromContext(ctx),
		}, &resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) AudienceUpdate(ctx context.Context, params *atomic.AudienceUpdateInput) (*atomic.Audience, error) {
	var resp ResponseProxy[atomic.Audience]

	if err := params.Validate(); err != nil {
		return nil, err
	}

	path := fmt.Sprintf(AudienceUpdatePath, params.AudienceID.String())

	if err := c.Backend.ExecContext(
		ctx,
		http.MethodPut,
		path,
		&ParamsProxy[atomic.AudienceUpdateInput]{
			methodParams:  *params,
			requestParams: ParamsFromContext(ctx),
		}, &resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) AudienceDelete(ctx context.Context, params *atomic.AudienceDeleteInput) error {
	if err := params.Validate(); err != nil {
		return err
	}

	path := fmt.Sprintf(AudienceDeletePath, params.AudienceID.String())

	if err := c.Backend.ExecContext(
		ctx,
		http.MethodDelete,
		path,
		&ParamsProxy[atomic.AudienceDeleteInput]{
			methodParams:  *params,
			requestParams: ParamsFromContext(ctx),
		}, nil); err != nil {
		return err
	}

	return nil
}

func (c *Client) AudienceList(ctx context.Context, params *atomic.AudienceListInput) ([]*atomic.Audience, error) {
	var resp ResponseProxy[[]*atomic.Audience]

	if err := params.Validate(); err != nil {
		return nil, err
	}

	if err := c.Backend.ExecContext(
		ctx,
		http.MethodGet,
		AudienceListPath,
		&ParamsProxy[atomic.AudienceListInput]{
			methodParams:  *params,
			requestParams: ParamsFromContext(ctx),
		}, &resp); err != nil {
		return nil, err
	}

	return resp.Value(), nil
}
