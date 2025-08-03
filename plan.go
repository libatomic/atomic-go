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
	PlanGetPath       = "/api/1.0.0/plans/%s"
	PlanCreatePath    = "/api/1.0.0/plans"
	PlanUpdatePath    = "/api/1.0.0/plans/%s"
	PlanDeletePath    = "/api/1.0.0/plans/%s"
	PlanListPath      = "/api/1.0.0/plans"
	PlanSubscribePath = "/api/1.0.0/plans/%s/subscribe"
)

func (c *Client) PlanGet(ctx context.Context, params *atomic.PlanGetInput) (*atomic.Plan, error) {
	var resp ResponseProxy[atomic.Plan]

	if err := params.Validate(); err != nil {
		return nil, err
	}

	path := fmt.Sprintf(PlanGetPath, params.PlanID.String())

	if err := c.Backend.ExecContext(
		ctx,
		http.MethodGet,
		path,
		&ParamsProxy[atomic.PlanGetInput]{
			methodParams:  *params,
			requestParams: ParamsFromContext(ctx),
		}, &resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) PlanCreate(ctx context.Context, params *atomic.PlanCreateInput) (*atomic.Plan, error) {
	var resp ResponseProxy[atomic.Plan]

	if err := params.Validate(); err != nil {
		return nil, err
	}

	if err := c.Backend.ExecContext(
		ctx,
		http.MethodPost,
		PlanCreatePath,
		&ParamsProxy[atomic.PlanCreateInput]{
			methodParams:  *params,
			requestParams: ParamsFromContext(ctx),
		}, &resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) PlanUpdate(ctx context.Context, params *atomic.PlanUpdateInput) (*atomic.Plan, error) {
	var resp ResponseProxy[atomic.Plan]

	if err := params.Validate(); err != nil {
		return nil, err
	}

	path := fmt.Sprintf(PlanUpdatePath, params.PlanID.String())

	if err := c.Backend.ExecContext(
		ctx,
		http.MethodPut,
		path,
		&ParamsProxy[atomic.PlanUpdateInput]{
			methodParams:  *params,
			requestParams: ParamsFromContext(ctx),
		}, &resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) PlanDelete(ctx context.Context, params *atomic.PlanDeleteInput) error {
	if err := params.Validate(); err != nil {
		return err
	}

	path := fmt.Sprintf(PlanDeletePath, params.PlanID.String())

	if err := c.Backend.ExecContext(
		ctx,
		http.MethodDelete,
		path,
		&ParamsProxy[atomic.PlanDeleteInput]{
			methodParams:  *params,
			requestParams: ParamsFromContext(ctx),
		}, nil); err != nil {
		return err
	}

	return nil
}

func (c *Client) PlanList(ctx context.Context, params *atomic.PlanListInput) ([]*atomic.Plan, error) {
	var resp ResponseProxy[[]*atomic.Plan]

	if err := params.Validate(); err != nil {
		return nil, err
	}

	if err := c.Backend.ExecContext(
		ctx,
		http.MethodGet,
		PlanListPath,
		&ParamsProxy[atomic.PlanListInput]{
			methodParams:  *params,
			requestParams: ParamsFromContext(ctx),
		}, &resp); err != nil {
		return nil, err
	}

	return resp.Value(), nil
}

func (c *Client) PlanSubscribe(ctx context.Context, params *atomic.PlanSubscribeInput) (*atomic.Subscription, error) {
	var resp ResponseProxy[atomic.Subscription]

	if err := params.Validate(); err != nil {
		return nil, err
	}

	path := fmt.Sprintf(PlanSubscribePath, params.PlanID.String())

	if err := c.Backend.ExecContext(
		ctx,
		http.MethodPost,
		path,
		&ParamsProxy[atomic.PlanSubscribeInput]{
			methodParams:  *params,
			requestParams: ParamsFromContext(ctx),
		}, &resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}
