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
	PartnerGetPath              = "/api/1.0.0/partners/%s"
	PartnerCreatePath           = "/api/1.0.0/partners"
	PartnerUpdatePath           = "/api/1.0.0/partners/%s"
	PartnerDeletePath           = "/api/1.0.0/partners/%s"
	PartnerListPath             = "/api/1.0.0/partners"
	PartnerCredentialCreatePath = "/api/1.0.0/partners/%s/credentials"
	PartnerCredentialGetPath    = "/api/1.0.0/partners/%s/credentials/%s"
	PartnerCredentialDeletePath = "/api/1.0.0/partners/%s/credentials/%s"
	PartnerTokenCreatePath      = "/api/1.0.0/partners/%s/tokens"
	PartnerTokenGetPath         = "/api/1.0.0/partners/%s/tokens/%s"
	PartnerTokenRevokePath      = "/api/1.0.0/partners/%s/tokens/%s"
)

func (c *Client) PartnerGet(ctx context.Context, params *atomic.PartnerGetInput) (*atomic.Partner, error) {
	var resp ResponseProxy[atomic.Partner]

	if err := params.Validate(); err != nil {
		return nil, err
	}

	path := fmt.Sprintf(PartnerGetPath, params.PartnerID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Get(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) PartnerCreate(ctx context.Context, params *atomic.PartnerCreateInput) (*atomic.Partner, error) {
	var resp ResponseProxy[atomic.Partner]

	if err := params.Validate(); err != nil {
		return nil, err
	}

	path := fmt.Sprintf(PartnerCreatePath)

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Post(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) PartnerUpdate(ctx context.Context, params *atomic.PartnerUpdateInput) (*atomic.Partner, error) {
	var resp ResponseProxy[atomic.Partner]

	if err := params.Validate(); err != nil {
		return nil, err
	}

	path := fmt.Sprintf(PartnerUpdatePath, params.PartnerID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Put(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) PartnerDelete(ctx context.Context, params *atomic.PartnerDeleteInput) error {
	if err := params.Validate(); err != nil {
		return err
	}

	path := fmt.Sprintf(PartnerDeletePath, params.PartnerID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Delete(),
		nil); err != nil {
		return err
	}

	return nil
}

func (c *Client) PartnerList(ctx context.Context, params *atomic.PartnerListInput) ([]*atomic.Partner, error) {
	var resp ResponseProxy[[]*atomic.Partner]

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, PartnerListPath, params).Get(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Value(), nil
}

func (c *Client) PartnerCredentialCreate(ctx context.Context, params *atomic.PartnerCredentialCreateInput) (*atomic.PartnerCredential, error) {
	var resp ResponseProxy[atomic.PartnerCredential]

	if err := params.Validate(); err != nil {
		return nil, err
	}

	path := fmt.Sprintf(PartnerCredentialCreatePath, params.PartnerID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Post(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) PartnerCredentialGet(ctx context.Context, params *atomic.PartnerCredentialGetInput) (*atomic.PartnerCredential, error) {
	var resp ResponseProxy[atomic.PartnerCredential]

	if err := params.Validate(); err != nil {
		return nil, err
	}

	path := fmt.Sprintf(PartnerCredentialGetPath, params.PartnerID.String(), *params.ClientID)

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Get(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) PartnerCredentialDelete(ctx context.Context, params *atomic.PartnerCredentialDeleteInput) error {
	if err := params.Validate(); err != nil {
		return err
	}

	path := fmt.Sprintf(PartnerCredentialDeletePath, params.PartnerID.String(), *params.ClientID)

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Delete(),
		nil); err != nil {
		return err
	}

	return nil
}

func (c *Client) PartnerTokenCreate(ctx context.Context, params *atomic.PartnerTokenCreateInput) (*atomic.PartnerAccessToken, error) {
	var resp ResponseProxy[atomic.PartnerAccessToken]

	if err := params.Validate(); err != nil {
		return nil, err
	}

	path := fmt.Sprintf(PartnerTokenCreatePath, params.PartnerID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Post(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) PartnerTokenGet(ctx context.Context, params *atomic.PartnerTokenGetInput) (*atomic.PartnerAccessToken, error) {
	var resp ResponseProxy[atomic.PartnerAccessToken]

	if err := params.Validate(); err != nil {
		return nil, err
	}

	path := fmt.Sprintf(PartnerTokenGetPath, params.PartnerID.String(), params.TokenID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Get(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) PartnerTokenRevoke(ctx context.Context, params *atomic.PartnerTokenRevokeInput) error {
	if err := params.Validate(); err != nil {
		return err
	}

	path := fmt.Sprintf(PartnerTokenRevokePath, params.PartnerID.String(), params.TokenID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Delete(),
		nil); err != nil {
		return err
	}

	return nil
}
