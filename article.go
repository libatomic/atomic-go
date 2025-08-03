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
	ArticleCreatePath = "/api/1.0.0/articles"
	ArticleUpdatePath = "/api/1.0.0/articles/%s"
	ArticleDeletePath = "/api/1.0.0/articles/%s"
	ArticleListPath   = "/api/1.0.0/articles"
	ArticleGetPath    = "/api/1.0.0/articles/%s"
)

func (c *Client) ArticleCreate(ctx context.Context, params *atomic.ArticleCreateInput) (*atomic.Article, error) {
	var resp ResponseProxy[atomic.Article]

	if err := params.Validate(); err != nil {
		return nil, err
	}

	if err := c.Backend.ExecContext(
		ctx,
		http.MethodPost,
		ArticleCreatePath,
		&ParamsProxy[atomic.ArticleCreateInput]{
			methodParams:  *params,
			requestParams: ParamsFromContext(ctx),
		}, &resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) ArticleGet(ctx context.Context, params *atomic.ArticleGetInput) (*atomic.Article, error) {
	var resp ResponseProxy[atomic.Article]

	if err := params.Validate(); err != nil {
		return nil, err
	}

	path := fmt.Sprintf(ArticleGetPath, params.ArticleID.String())

	if err := c.Backend.ExecContext(
		ctx,
		http.MethodGet,
		path,
		&ParamsProxy[atomic.ArticleGetInput]{
			methodParams:  *params,
			requestParams: ParamsFromContext(ctx),
		}, &resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) ArticleUpdate(ctx context.Context, params *atomic.ArticleUpdateInput) (*atomic.Article, error) {
	var resp ResponseProxy[atomic.Article]

	if err := params.Validate(); err != nil {
		return nil, err
	}

	path := fmt.Sprintf(ArticleUpdatePath, params.ArticleID.String())

	if err := c.Backend.ExecContext(
		ctx,
		http.MethodPut,
		path,
		&ParamsProxy[atomic.ArticleUpdateInput]{
			methodParams:  *params,
			requestParams: ParamsFromContext(ctx),
		}, &resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) ArticleDelete(ctx context.Context, params *atomic.ArticleDeleteInput) error {
	if err := params.Validate(); err != nil {
		return err
	}

	path := fmt.Sprintf(ArticleDeletePath, params.ArticleID.String())

	return c.Backend.ExecContext(ctx, http.MethodDelete, path, &ParamsProxy[atomic.ArticleDeleteInput]{
		methodParams:  *params,
		requestParams: ParamsFromContext(ctx),
	}, nil)
}

func (c *Client) ArticleList(ctx context.Context, params *atomic.ArticleListInput) ([]*atomic.Article, error) {
	var resp ResponseProxy[[]*atomic.Article]

	if err := c.Backend.ExecContext(
		ctx,
		http.MethodGet,
		ArticleListPath,
		&ParamsProxy[atomic.ArticleListInput]{
			methodParams:  *params,
			requestParams: ParamsFromContext(ctx),
		}, &resp); err != nil {
		return nil, err
	}

	return resp.Value(), nil
}
