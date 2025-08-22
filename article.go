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
	Article            = atomic.Article
	ArticleCreateInput = atomic.ArticleCreateInput
	ArticleGetInput    = atomic.ArticleGetInput
	ArticleUpdateInput = atomic.ArticleUpdateInput
	ArticleDeleteInput = atomic.ArticleDeleteInput
	ArticleListInput   = atomic.ArticleListInput
)

const (
	ArticleCreatePath = "/api/1.0.0/articles"
	ArticleUpdatePath = "/api/1.0.0/articles/%s"
	ArticleDeletePath = "/api/1.0.0/articles/%s"
	ArticleListPath   = "/api/1.0.0/articles"
	ArticleGetPath    = "/api/1.0.0/articles/%s"
)

func (c *Client) ArticleCreate(ctx context.Context, params *ArticleCreateInput) (*Article, error) {
	var resp ResponseProxy[Article]

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, ArticleCreatePath, params).Post(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) ArticleGet(ctx context.Context, params *ArticleGetInput) (*Article, error) {
	var resp ResponseProxy[Article]

	path := fmt.Sprintf(ArticleGetPath, params.ArticleID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Get(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) ArticleUpdate(ctx context.Context, params *ArticleUpdateInput) (*Article, error) {
	var resp ResponseProxy[Article]

	path := fmt.Sprintf(ArticleUpdatePath, params.ArticleID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Put(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) ArticleDelete(ctx context.Context, params *ArticleDeleteInput) error {
	path := fmt.Sprintf(ArticleDeletePath, params.ArticleID.String())

	return c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Delete(),
		nil,
	)
}

func (c *Client) ArticleList(ctx context.Context, params *ArticleListInput) ([]*Article, error) {
	var resp ResponseProxy[[]*Article]

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, ArticleListPath, params).Get(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Value(), nil
}
