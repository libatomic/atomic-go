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
	SubscriptionGetPath    = "/api/1.0.0/subscriptions/%s"
	SubscriptionListPath   = "/api/1.0.0/subscriptions"
	SubscriptionCreatePath = "/api/1.0.0/subscriptions"
	SubscriptionUpdatePath = "/api/1.0.0/subscriptions/%s"
	SubscriptionDeletePath = "/api/1.0.0/subscriptions/%s"
)

func (c *Client) SubscriptionGet(ctx context.Context, params *atomic.SubscriptionGetInput) (*atomic.Subscription, error) {
	var resp ResponseProxy[atomic.Subscription]

	if err := params.Validate(); err != nil {
		return nil, err
	}

	path := fmt.Sprintf(SubscriptionGetPath, params.SubscriptionID.String())

	if err := c.Backend.ExecContext(
		ctx,
		http.MethodGet,
		path,
		&ParamsProxy[atomic.SubscriptionGetInput]{
			methodParams:  *params,
			requestParams: ParamsFromContext(ctx),
		}, &resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) SubscriptionList(ctx context.Context, params *atomic.SubscriptionListInput) ([]*atomic.Subscription, error) {
	var resp ResponseProxy[[]*atomic.Subscription]

	if err := params.Validate(); err != nil {
		return nil, err
	}

	if err := c.Backend.ExecContext(
		ctx,
		http.MethodGet,
		SubscriptionListPath,
		&ParamsProxy[atomic.SubscriptionListInput]{
			methodParams:  *params,
			requestParams: ParamsFromContext(ctx),
		}, &resp); err != nil {
		return nil, err
	}

	return resp.Value(), nil
}

func (c *Client) SubscriptionCreate(ctx context.Context, params *atomic.SubscriptionCreateInput) (*atomic.Subscription, error) {
	var resp ResponseProxy[atomic.Subscription]

	if err := params.Validate(); err != nil {
		return nil, err
	}

	if err := c.Backend.ExecContext(
		ctx,
		http.MethodPost,
		SubscriptionCreatePath,
		&ParamsProxy[atomic.SubscriptionCreateInput]{
			methodParams:  *params,
			requestParams: ParamsFromContext(ctx),
		}, &resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) SubscriptionUpdate(ctx context.Context, params *atomic.SubscriptionUpdateInput) (*atomic.Subscription, error) {
	var resp ResponseProxy[atomic.Subscription]

	if err := params.Validate(); err != nil {
		return nil, err
	}

	path := fmt.Sprintf(SubscriptionUpdatePath, params.SubscriptionID.String())

	if err := c.Backend.ExecContext(
		ctx,
		http.MethodPut,
		path,
		&ParamsProxy[atomic.SubscriptionUpdateInput]{
			methodParams:  *params,
			requestParams: ParamsFromContext(ctx),
		}, &resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) SubscriptionDelete(ctx context.Context, params *atomic.SubscriptionDeleteInput) error {
	if err := params.Validate(); err != nil {
		return err
	}

	path := fmt.Sprintf(SubscriptionDeletePath, params.SubscriptionID.String())

	if err := c.Backend.ExecContext(
		ctx,
		http.MethodDelete,
		path,
		&ParamsProxy[atomic.SubscriptionDeleteInput]{
			methodParams:  *params,
			requestParams: ParamsFromContext(ctx),
		}, nil); err != nil {
		return err
	}

	return nil
}
