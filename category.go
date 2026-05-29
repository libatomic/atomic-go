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
	Category            = atomic.Category
	CategoryCreateInput = atomic.CategoryCreateInput
	CategoryGetInput    = atomic.CategoryGetInput
	CategoryUpdateInput = atomic.CategoryUpdateInput
	CategoryDeleteInput = atomic.CategoryDeleteInput
	CategoryListInput   = atomic.CategoryListInput
)

const (
	CategoryGetPath    = "/api/1.0.0/categories/%s"
	CategoryCreatePath = "/api/1.0.0/categories"
	CategoryUpdatePath = "/api/1.0.0/categories/%s"
	CategoryDeletePath = "/api/1.0.0/categories/%s"
	CategoryListPath   = "/api/1.0.0/categories"
)

func (c *Client) CategoryGet(ctx context.Context, params *CategoryGetInput) (*Category, error) {
	var resp ResponseProxy[Category]

	path := fmt.Sprintf(CategoryGetPath, params.CategoryID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Get(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) CategoryCreate(ctx context.Context, params *CategoryCreateInput) (*Category, error) {
	var resp ResponseProxy[Category]

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, CategoryCreatePath, params).Post(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) CategoryUpdate(ctx context.Context, params *CategoryUpdateInput) (*Category, error) {
	var resp ResponseProxy[Category]

	path := fmt.Sprintf(CategoryUpdatePath, params.CategoryID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Put(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) CategoryDelete(ctx context.Context, params *CategoryDeleteInput) error {
	path := fmt.Sprintf(CategoryDeletePath, params.CategoryID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Delete(),
		nil,
	); err != nil {
		return err
	}

	return nil
}

func (c *Client) CategoryList(ctx context.Context, params *CategoryListInput) ([]*Category, error) {
	var resp ResponseProxy[[]*Category]

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, CategoryListPath, params).Get(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Value(), nil
}
