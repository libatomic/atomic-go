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
	PublisherCreatePath = "/api/1.0.0/publishers"
	PublisherUpdatePath = "/api/1.0.0/publishers/%s"
	PublisherGetPath    = "/api/1.0.0/publishers/%s"
	PublisherListPath   = "/api/1.0.0/publishers"
	PublisherDeletePath = "/api/1.0.0/publishers/%s"
)

func (c *Client) PublisherCreate(ctx context.Context, params *atomic.PublisherCreateInput) (*atomic.Publisher, error) {
	var resp ResponseProxy[atomic.Publisher]

	if err := params.Validate(); err != nil {
		return nil, err
	}

	if err := c.Backend.ExecContext(
		ctx,
		http.MethodPost,
		PublisherCreatePath,
		NewParamsProxy(ctx, params),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) PublisherUpdate(ctx context.Context, params *atomic.PublisherUpdateInput) (*atomic.Publisher, error) {
	var resp ResponseProxy[atomic.Publisher]

	if err := params.Validate(); err != nil {
		return nil, err
	}

	if err := c.Backend.ExecContext(
		ctx,
		http.MethodPut,
		fmt.Sprintf(PublisherUpdatePath, params.PublisherID.String()),
		NewParamsProxy(ctx, params),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) PublisherGet(ctx context.Context, params *atomic.PublisherGetInput) (*atomic.Publisher, error) {
	var resp ResponseProxy[atomic.Publisher]

	if err := params.Validate(); err != nil {
		return nil, err
	}

	if err := c.Backend.ExecContext(
		ctx,
		http.MethodGet,
		fmt.Sprintf(PublisherGetPath, params.PublisherID.String()),
		NewParamsProxy(ctx, params),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) PublisherList(ctx context.Context, params *atomic.PublisherListInput) (*atomic.PublisherListOutput, error) {
	var resp ResponseProxy[atomic.PublisherListOutput]

	if err := params.Validate(); err != nil {
		return nil, err
	}

	if err := c.Backend.ExecContext(
		ctx,
		http.MethodGet,
		PublisherListPath,
		NewParamsProxy(ctx, params),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) PublisherDelete(ctx context.Context, params *atomic.PublisherDeleteInput) error {
	var resp ResponseProxy[atomic.Publisher]

	if err := params.Validate(); err != nil {
		return err
	}

	if err := c.Backend.ExecContext(
		ctx,
		http.MethodDelete,
		fmt.Sprintf(PublisherDeletePath, params.PublisherID.String()),
		NewParamsProxy(ctx, params),
		&resp); err != nil {
		return err
	}

	return nil
}
