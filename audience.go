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
	Audience            = atomic.Audience
	AudienceCreateInput = atomic.AudienceCreateInput
	AudienceGetInput    = atomic.AudienceGetInput
	AudienceUpdateInput = atomic.AudienceUpdateInput
	AudienceDeleteInput = atomic.AudienceDeleteInput
	AudienceListInput   = atomic.AudienceListInput
)

const (
	AudienceGetPath    = "/api/1.0.0/audiences/%s"
	AudienceCreatePath = "/api/1.0.0/audiences"
	AudienceUpdatePath = "/api/1.0.0/audiences/%s"
	AudienceDeletePath = "/api/1.0.0/audiences/%s"
	AudienceListPath   = "/api/1.0.0/audiences"
)

func (c *Client) AudienceGet(ctx context.Context, params *AudienceGetInput) (*Audience, error) {
	var resp ResponseProxy[Audience]

	path := fmt.Sprintf(AudienceGetPath, params.AudienceID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Get(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) AudienceCreate(ctx context.Context, params *AudienceCreateInput) (*Audience, error) {
	var resp ResponseProxy[Audience]

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, AudienceCreatePath, params).Post(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) AudienceUpdate(ctx context.Context, params *AudienceUpdateInput) (*Audience, error) {
	var resp ResponseProxy[Audience]

	path := fmt.Sprintf(AudienceUpdatePath, params.AudienceID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Put(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) AudienceDelete(ctx context.Context, params *AudienceDeleteInput) error {
	path := fmt.Sprintf(AudienceDeletePath, params.AudienceID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Delete(),
		nil,
	); err != nil {
		return err
	}

	return nil
}

func (c *Client) AudienceList(ctx context.Context, params *AudienceListInput) ([]*Audience, error) {
	var resp ResponseProxy[[]*Audience]

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, AudienceListPath, params).Get(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Value(), nil
}
