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

type (
	AccessToken            = atomic.AccessToken
	AccessTokenCreateInput = atomic.AccessTokenCreateInput
	AccessTokenGetInput    = atomic.AccessTokenGetInput
	AccessTokenRevokeInput = atomic.AccessTokenRevokeInput
)

const (
	UserTokenCreatePath   = "/api/1.0.0/users/%s/tokens"
	AppTokenCreatePath    = "/api/1.0.0/applications/%s/tokens"
	AccessTokenGetPath    = "/api/1.0.0/tokens/%s"
	AccessTokenRevokePath = "/api/1.0.0/tokens/%s"
	AccessTokenDeletePath = "/api/1.0.0/tokens/%s"
)

func (c *Client) AccessTokenCreate(ctx context.Context, params *AccessTokenCreateInput) (*AccessToken, error) {
	var resp ResponseProxy[AccessToken]
	var path string

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

func (c *Client) AccessTokenGet(ctx context.Context, params *AccessTokenGetInput) (*AccessToken, error) {
	var resp ResponseProxy[AccessToken]

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

func (c *Client) AccessTokenRevoke(ctx context.Context, params *AccessTokenRevokeInput) error {
	path := fmt.Sprintf(AccessTokenRevokePath, params.AccessTokenID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Delete(),
		nil); err != nil {
		return err
	}

	return nil
}
