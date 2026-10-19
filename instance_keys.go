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

	"github.com/libatomic/atomic/v2/pkg/atomic"
)

type (
	InstanceKey            = atomic.InstanceKey
	InstanceKeyKind        = atomic.InstanceKeyKind
	InstanceKeyListInput   = atomic.InstanceKeyListInput
	InstanceKeyRotateInput = atomic.InstanceKeyRotateInput
	InstanceKeyRevertInput = atomic.InstanceKeyRevertInput
	InstanceKeyRetireInput = atomic.InstanceKeyRetireInput
)

const (
	InstanceKeyListPath   = "/api/1.0.0/instances/%s/keys"
	InstanceKeyRotatePath = "/api/1.0.0/instances/%s/keys"
	InstanceKeyRevertPath = "/api/1.0.0/instances/%s/keys/revert"
	InstanceKeyRetirePath = "/api/1.0.0/instances/%s/keys/%s"
)

// InstanceKeyList lists the instance's token and session keys, current
// first, without their material.
func (c *Client) InstanceKeyList(ctx context.Context, params *InstanceKeyListInput) ([]*InstanceKey, error) {
	var resp ResponseProxy[[]*InstanceKey]

	path := fmt.Sprintf(InstanceKeyListPath, idString(params.InstanceID))

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Get(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Value(), nil
}

// InstanceKeyRotate makes a new current key of one kind; the key it
// replaces keeps verifying for the grace period.
func (c *Client) InstanceKeyRotate(ctx context.Context, params *InstanceKeyRotateInput) (*InstanceKey, error) {
	var resp ResponseProxy[InstanceKey]

	path := fmt.Sprintf(InstanceKeyRotatePath, idString(params.InstanceID))

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Post(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

// InstanceKeyRevert undoes the last rotation of one kind.
func (c *Client) InstanceKeyRevert(ctx context.Context, params *InstanceKeyRevertInput) (*InstanceKey, error) {
	var resp ResponseProxy[InstanceKey]

	path := fmt.Sprintf(InstanceKeyRevertPath, idString(params.InstanceID))

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Post(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

// InstanceKeyRetire ends a previous key's grace period now.
func (c *Client) InstanceKeyRetire(ctx context.Context, params *InstanceKeyRetireInput) (*InstanceKey, error) {
	var resp ResponseProxy[InstanceKey]

	path := fmt.Sprintf(InstanceKeyRetirePath, idString(params.InstanceID), params.KeyID)

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Delete(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}
