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
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/textproto"
	"strconv"

	"github.com/libatomic/atomic/pkg/atomic"
)

type (
	User            = atomic.User
	UserGetInput    = atomic.UserGetInput
	UserCreateInput = atomic.UserCreateInput
	UserUpdateInput = atomic.UserUpdateInput
	UserDeleteInput = atomic.UserDeleteInput
	UserListInput   = atomic.UserListInput
	UserImportInput = atomic.UserImportInput
)

const (
	UserGetPath    = "/api/1.0.0/users/%s"
	UserCreatePath = "/api/1.0.0/users"
	UserUpdatePath = "/api/1.0.0/users/%s"
	UserDeletePath = "/api/1.0.0/users/%s"
	UserListPath   = "/api/1.0.0/users"
	UserImportPath = "/api/1.0.0/users/import"
)

func (c *Client) UserGet(ctx context.Context, params *UserGetInput) (*User, error) {
	var resp ResponseProxy[User]

	path := fmt.Sprintf(UserGetPath, params.UserID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Get(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) UserCreate(ctx context.Context, params *UserCreateInput) (*User, error) {
	var resp ResponseProxy[User]

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, UserCreatePath, params).Post(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) UserUpdate(ctx context.Context, params *UserUpdateInput) (*User, error) {
	var resp ResponseProxy[User]

	path := fmt.Sprintf(UserUpdatePath, params.UserID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Put(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) UserDelete(ctx context.Context, params *UserDeleteInput) error {
	path := fmt.Sprintf(UserDeletePath, params.UserID.String())

	return c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Delete(),
		nil,
	)
}

func (c *Client) UserList(ctx context.Context, params *UserListInput) ([]*User, error) {
	var resp ResponseProxy[[]*User]

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, UserListPath, params).Get(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Value(), nil
}

func (c *Client) UserImport(ctx context.Context, params *UserImportInput) (*Job, error) {
	var resp ResponseProxy[Job]

	if params.File == nil {
		return nil, errors.New("file is required")
	}

	if params.Filename == "" {
		return nil, errors.New("filename is required")
	}

	var body bytes.Buffer

	writer := multipart.NewWriter(&body)

	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"; filename="%s"`, "file", params.Filename))
	h.Set("Content-Type", params.MimeType)
	h.Set("Content-Length", strconv.FormatInt(params.Size, 10))

	part, err := writer.CreatePart(h)
	if err != nil {
		return nil, fmt.Errorf("failed to create multipart part: %w", err)
	}

	if _, err := io.Copy(part, params.File); err != nil {
		return nil, fmt.Errorf("failed to copy file to multipart writer: %w", err)
	}

	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("failed to close multipart writer: %w", err)
	}

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, UserImportPath, params).Post().
			WithContentType(writer.FormDataContentType()).
			WithEncoding(ParamsEncodingQuery).
			WithBody(&body),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}
