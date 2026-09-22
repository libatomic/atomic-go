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
	Secret            = atomic.Secret
	SecretListInput   = atomic.SecretListInput
	SecretSetInput    = atomic.SecretSetInput
	SecretDeleteInput = atomic.SecretDeleteInput
)

const (
	SecretListPath   = "/api/1.0.0/secrets"
	SecretSetPath    = "/api/1.0.0/secrets"
	SecretDeletePath = "/api/1.0.0/secrets/%s"
)

func (c *Client) SecretList(ctx context.Context, params *SecretListInput) ([]*Secret, error) {
	var resp ResponseProxy[[]*Secret]

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, SecretListPath, params).Get(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Value(), nil
}

func (c *Client) SecretSet(ctx context.Context, params *SecretSetInput) (*Secret, error) {
	var resp ResponseProxy[Secret]

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, SecretSetPath, params).Put(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) SecretDelete(ctx context.Context, params *SecretDeleteInput) error {
	path := fmt.Sprintf(SecretDeletePath, params.Name)

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Delete(),
		nil,
	); err != nil {
		return err
	}

	return nil
}
