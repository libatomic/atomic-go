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
	"errors"
	"fmt"

	"github.com/libatomic/atomic/pkg/atomic"
)

const (
	UserTokenCreatePath   = "/api/1.0.0/users/%s/tokens"
	AppTokenCreatePath    = "/api/1.0.0/applications/%s/tokens"
	AccessTokenGetPath    = "/api/1.0.0/tokens/%s"
	AccessTokenRevokePath = "/api/1.0.0/tokens/%s"
	AccessTokenDeletePath = "/api/1.0.0/tokens/%s"
)

func (c *Client) AccessTokenCreate(ctx context.Context, params *atomic.AccessTokenCreateInput) (*atomic.AccessToken, error) {
	var resp ResponseProxy[atomic.AccessToken]
	var path string

	if err := params.Validate(); err != nil {
		return nil, err
	}

	if params.UserID != nil {
		path = fmt.Sprintf(UserTokenCreatePath, params.UserID.String())
	} else if params.ApplicationID != nil {
		path = fmt.Sprintf(AppTokenCreatePath, params.ApplicationID.String())
	} else {
		return nil, errors.New("user_id or application_id is required")
	}

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Get(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) AccessTokenGet(ctx context.Context, params *atomic.AccessTokenGetInput) (*atomic.AccessToken, error) {
	var resp ResponseProxy[atomic.AccessToken]

	if err := params.Validate(); err != nil {
		return nil, err
	}

	path := fmt.Sprintf(AccessTokenGetPath, params.AccessTokenID.String())

	if params.AccessTokenID == nil {
		return nil, errors.New("token_id is required")
	}

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Get(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) AccessTokenRevoke(ctx context.Context, params *atomic.AccessTokenRevokeInput) error {
	if err := params.Validate(); err != nil {
		return err
	}

	path := fmt.Sprintf(AccessTokenRevokePath, params.AccessTokenID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Delete(),
		nil); err != nil {
		return err
	}

	return nil
}
