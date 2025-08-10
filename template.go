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
	TemplateGetPath    = "/api/1.0.0/templates/%s"
	TemplateListPath   = "/api/1.0.0/templates"
	TemplateCreatePath = "/api/1.0.0/templates"
	TemplateUpdatePath = "/api/1.0.0/templates/%s"
	TemplateDeletePath = "/api/1.0.0/templates/%s"
)

func (c *Client) TemplateGet(ctx context.Context, params *atomic.TemplateGetInput) (*atomic.Template, error) {
	var resp ResponseProxy[atomic.Template]

	if err := params.Validate(); err != nil {
		return nil, err
	}

	path := fmt.Sprintf(TemplateGetPath, params.TemplateID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Get(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) TemplateList(ctx context.Context, params *atomic.TemplateListInput) ([]*atomic.Template, error) {
	var resp ResponseProxy[[]*atomic.Template]

	if err := params.Validate(); err != nil {
		return nil, err
	}

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, TemplateListPath, params).Get(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Value(), nil
}

func (c *Client) TemplateCreate(ctx context.Context, params *atomic.TemplateCreateInput) (*atomic.Template, error) {
	var resp ResponseProxy[atomic.Template]

	if err := params.Validate(); err != nil {
		return nil, err
	}

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, TemplateCreatePath, params).Post(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) TemplateUpdate(ctx context.Context, params *atomic.TemplateUpdateInput) (*atomic.Template, error) {
	var resp ResponseProxy[atomic.Template]

	if err := params.Validate(); err != nil {
		return nil, err
	}

	path := fmt.Sprintf(TemplateUpdatePath, params.TemplateID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Put(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) TemplateDelete(ctx context.Context, params *atomic.TemplateDeleteInput) error {
	if err := params.Validate(); err != nil {
		return err
	}

	path := fmt.Sprintf(TemplateDeletePath, params.TemplateID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Delete(),
		nil,
	); err != nil {
		return err
	}

	return nil
}
