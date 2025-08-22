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
	Partner                      = atomic.Partner
	PartnerCredential            = atomic.PartnerCredential
	PartnerAccessToken           = atomic.PartnerAccessToken
	PartnerCreateInput           = atomic.PartnerCreateInput
	PartnerGetInput              = atomic.PartnerGetInput
	PartnerUpdateInput           = atomic.PartnerUpdateInput
	PartnerDeleteInput           = atomic.PartnerDeleteInput
	PartnerListInput             = atomic.PartnerListInput
	PartnerCredentialCreateInput = atomic.PartnerCredentialCreateInput
	PartnerCredentialGetInput    = atomic.PartnerCredentialGetInput
	PartnerCredentialDeleteInput = atomic.PartnerCredentialDeleteInput
	PartnerTokenCreateInput      = atomic.PartnerTokenCreateInput
	PartnerTokenGetInput         = atomic.PartnerTokenGetInput
	PartnerTokenRevokeInput      = atomic.PartnerTokenRevokeInput
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

func (c *Client) PartnerGet(ctx context.Context, params *PartnerGetInput) (*Partner, error) {
	var resp ResponseProxy[Partner]

	path := fmt.Sprintf(PartnerGetPath, params.PartnerID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Get(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) PartnerCreate(ctx context.Context, params *PartnerCreateInput) (*Partner, error) {
	var resp ResponseProxy[Partner]

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, PartnerCreatePath, params).Post(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) PartnerUpdate(ctx context.Context, params *PartnerUpdateInput) (*Partner, error) {
	var resp ResponseProxy[Partner]

	path := fmt.Sprintf(PartnerUpdatePath, params.PartnerID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Put(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) PartnerDelete(ctx context.Context, params *PartnerDeleteInput) error {
	path := fmt.Sprintf(PartnerDeletePath, params.PartnerID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Delete(),
		nil); err != nil {
		return err
	}

	return nil
}

func (c *Client) PartnerList(ctx context.Context, params *PartnerListInput) ([]*Partner, error) {
	var resp ResponseProxy[[]*Partner]

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, PartnerListPath, params).Get(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Value(), nil
}

func (c *Client) PartnerCredentialCreate(ctx context.Context, params *PartnerCredentialCreateInput) (*PartnerCredential, error) {
	var resp ResponseProxy[PartnerCredential]

	path := fmt.Sprintf(PartnerCredentialCreatePath, params.PartnerID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Post(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) PartnerCredentialGet(ctx context.Context, params *PartnerCredentialGetInput) (*PartnerCredential, error) {
	var resp ResponseProxy[PartnerCredential]

	path := fmt.Sprintf(PartnerCredentialGetPath, params.PartnerID.String(), *params.ClientID)

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Get(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) PartnerCredentialDelete(ctx context.Context, params *PartnerCredentialDeleteInput) error {
	path := fmt.Sprintf(PartnerCredentialDeletePath, params.PartnerID.String(), *params.ClientID)

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Delete(),
		nil); err != nil {
		return err
	}

	return nil
}

func (c *Client) PartnerTokenCreate(ctx context.Context, params *PartnerTokenCreateInput) (*PartnerAccessToken, error) {
	var resp ResponseProxy[PartnerAccessToken]

	path := fmt.Sprintf(PartnerTokenCreatePath, params.PartnerID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Post(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) PartnerTokenGet(ctx context.Context, params *PartnerTokenGetInput) (*PartnerAccessToken, error) {
	var resp ResponseProxy[PartnerAccessToken]

	path := fmt.Sprintf(PartnerTokenGetPath, params.PartnerID.String(), params.TokenID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Get(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) PartnerTokenRevoke(ctx context.Context, params *PartnerTokenRevokeInput) error {
	path := fmt.Sprintf(PartnerTokenRevokePath, params.PartnerID.String(), params.TokenID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Delete(),
		nil); err != nil {
		return err
	}

	return nil
}
