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
	Asset            = atomic.Asset
	AssetCreateInput = atomic.AssetCreateInput
	AssetGetInput    = atomic.AssetGetInput
	AssetUpdateInput = atomic.AssetUpdateInput
	AssetDeleteInput = atomic.AssetDeleteInput
	AssetListInput   = atomic.AssetListInput
)

const (
	AssetCreatePath = "/api/1.0.0/assets"
	AssetGetPath    = "/api/1.0.0/assets/%s"
	AssetUpdatePath = "/api/1.0.0/assets/%s"
	AssetDeletePath = "/api/1.0.0/assets/%s"
	AssetListPath   = "/api/1.0.0/assets"
)

func (c *Client) AssetCreate(ctx context.Context, params *AssetCreateInput) (*Asset, error) {
	var resp ResponseProxy[Asset]

	if params.Payload == nil {
		return nil, errors.New("payload is required")
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

	if _, err := io.Copy(part, params.Payload); err != nil {
		return nil, fmt.Errorf("failed to copy file to multipart writer: %w", err)
	}

	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("failed to close multipart writer: %w", err)
	}

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, AssetCreatePath, params).Post().
			WithContentType(writer.FormDataContentType()).
			WithEncoding(ParamsEncodingQuery).
			WithBody(&body),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) AssetGet(ctx context.Context, params *AssetGetInput) (*Asset, error) {
	var resp ResponseProxy[Asset]

	path := fmt.Sprintf(AssetGetPath, params.AssetID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Get(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) AssetUpdate(ctx context.Context, params *AssetUpdateInput) (*Asset, error) {
	var resp ResponseProxy[Asset]

	path := fmt.Sprintf(AssetUpdatePath, params.AssetID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Put(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) AssetDelete(ctx context.Context, params *AssetDeleteInput) error {
	path := fmt.Sprintf(AssetDeletePath, params.AssetID.String())

	return c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Delete(),
		nil,
	)
}

func (c *Client) AssetList(ctx context.Context, params *AssetListInput) ([]*Asset, error) {
	var resp ResponseProxy[[]*Asset]

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, AssetListPath, params).Get(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Value(), nil
}
