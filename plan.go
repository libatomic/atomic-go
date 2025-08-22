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
	Plan               = atomic.Plan
	PlanCreateInput    = atomic.PlanCreateInput
	PlanGetInput       = atomic.PlanGetInput
	PlanUpdateInput    = atomic.PlanUpdateInput
	PlanDeleteInput    = atomic.PlanDeleteInput
	PlanListInput      = atomic.PlanListInput
	PlanSubscribeInput = atomic.PlanSubscribeInput
)

const (
	PlanGetPath       = "/api/1.0.0/plans/%s"
	PlanCreatePath    = "/api/1.0.0/plans"
	PlanUpdatePath    = "/api/1.0.0/plans/%s"
	PlanDeletePath    = "/api/1.0.0/plans/%s"
	PlanListPath      = "/api/1.0.0/plans"
	PlanSubscribePath = "/api/1.0.0/plans/%s/subscribe"
)

func (c *Client) PlanGet(ctx context.Context, params *PlanGetInput) (*Plan, error) {
	var resp ResponseProxy[Plan]

	path := fmt.Sprintf(PlanGetPath, params.PlanID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Get(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) PlanCreate(ctx context.Context, params *PlanCreateInput) (*Plan, error) {
	var resp ResponseProxy[Plan]

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, PlanCreatePath, params).Post(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) PlanUpdate(ctx context.Context, params *PlanUpdateInput) (*Plan, error) {
	var resp ResponseProxy[Plan]

	path := fmt.Sprintf(PlanUpdatePath, params.PlanID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Put(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) PlanDelete(ctx context.Context, params *PlanDeleteInput) error {
	path := fmt.Sprintf(PlanDeletePath, params.PlanID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Delete(),
		nil,
	); err != nil {
		return err
	}

	return nil
}

func (c *Client) PlanList(ctx context.Context, params *PlanListInput) ([]*Plan, error) {
	var resp ResponseProxy[[]*Plan]

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, PlanListPath, params).Get(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Value(), nil
}

func (c *Client) PlanSubscribe(ctx context.Context, params *PlanSubscribeInput) (*Subscription, error) {
	var resp ResponseProxy[Subscription]

	path := fmt.Sprintf(PlanSubscribePath, params.PlanID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Post(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}
