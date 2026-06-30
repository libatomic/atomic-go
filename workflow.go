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
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/libatomic/atomic/pkg/atomic"
)

type (
	Workflow             = atomic.Workflow
	WorkflowDefinition   = atomic.WorkflowDefinition
	WorkflowCreateInput  = atomic.WorkflowCreateInput
	WorkflowUpdateInput  = atomic.WorkflowUpdateInput
	WorkflowGetInput     = atomic.WorkflowGetInput
	WorkflowListInput    = atomic.WorkflowListInput
	WorkflowDeleteInput  = atomic.WorkflowDeleteInput
	WorkflowRunInput     = atomic.WorkflowRunInput
	WorkflowRun          = atomic.WorkflowRun
	WorkflowRunListInput = atomic.WorkflowRunListInput
	WorkflowRunGetInput  = atomic.WorkflowRunGetInput
)

const (
	WorkflowListPath   = "/api/1.0.0/workflows"
	WorkflowCreatePath = "/api/1.0.0/workflows"
	WorkflowGetPath    = "/api/1.0.0/workflows/%s"
	WorkflowUpdatePath = "/api/1.0.0/workflows/%s"
	WorkflowDeletePath = "/api/1.0.0/workflows/%s"
	WorkflowRunPath    = "/api/1.0.0/workflows/%s/run"
	WorkflowRunListPath = "/api/1.0.0/workflows/%s/runs"
	WorkflowRunGetPath  = "/api/1.0.0/workflows/%s/runs/%s"
)

func (c *Client) WorkflowList(ctx context.Context, params *WorkflowListInput) ([]*Workflow, error) {
	var resp ResponseProxy[[]*Workflow]

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, WorkflowListPath, params).Get(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Value(), nil
}

func (c *Client) WorkflowCreate(ctx context.Context, params *WorkflowCreateInput, body io.Reader, contentType string) (*Workflow, error) {
	var resp ResponseProxy[Workflow]

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, WorkflowCreatePath, params).
			Post().
			WithContentType(contentType).
			WithBody(body),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) WorkflowGet(ctx context.Context, params *WorkflowGetInput) (*Workflow, error) {
	var resp ResponseProxy[Workflow]

	var pathID string
	if params.WorkflowID != nil {
		pathID = params.WorkflowID.String()
	} else if params.Slug != nil {
		pathID = *params.Slug
	}

	path := fmt.Sprintf(WorkflowGetPath, pathID)

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Get(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) WorkflowUpdate(ctx context.Context, params *WorkflowUpdateInput, body io.Reader, contentType string) (*Workflow, error) {
	var resp ResponseProxy[Workflow]

	path := fmt.Sprintf(WorkflowUpdatePath, params.WorkflowID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).
			Put().
			WithContentType(contentType).
			WithBody(body),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) WorkflowDelete(ctx context.Context, params *WorkflowDeleteInput) error {
	path := fmt.Sprintf(WorkflowDeletePath, params.WorkflowID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Delete(),
		nil,
	); err != nil {
		return err
	}

	return nil
}

func (c *Client) WorkflowRun(ctx context.Context, params *WorkflowRunInput) (*WorkflowRun, error) {
	var resp ResponseProxy[WorkflowRun]

	path := fmt.Sprintf(WorkflowRunPath, params.WorkflowID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Post(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) WorkflowRunList(ctx context.Context, params *WorkflowRunListInput) ([]*WorkflowRun, error) {
	var resp ResponseProxy[[]*WorkflowRun]

	path := fmt.Sprintf(WorkflowRunListPath, params.WorkflowID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Get(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Value(), nil
}

func (c *Client) WorkflowRunGet(ctx context.Context, params *WorkflowRunGetInput) (*WorkflowRun, error) {
	var resp ResponseProxy[WorkflowRun]

	wfID := ""
	if params.WorkflowID != nil {
		wfID = params.WorkflowID.String()
	}

	path := fmt.Sprintf(WorkflowRunGetPath, wfID, params.RunID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Get(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

// WorkflowCreateFromBytes is a convenience wrapper that creates a workflow from
// raw YAML or JSON bytes, auto-detecting the content type.
func (c *Client) WorkflowCreateFromBytes(ctx context.Context, params *WorkflowCreateInput, data []byte) (*Workflow, error) {
	ct := detectContentType(data)
	return c.WorkflowCreate(ctx, params, bytes.NewReader(data), ct)
}

// WorkflowUpdateFromBytes is a convenience wrapper that updates a workflow from
// raw YAML or JSON bytes, auto-detecting the content type.
func (c *Client) WorkflowUpdateFromBytes(ctx context.Context, params *WorkflowUpdateInput, data []byte) (*Workflow, error) {
	ct := detectContentType(data)
	return c.WorkflowUpdate(ctx, params, bytes.NewReader(data), ct)
}

func detectContentType(data []byte) string {
	trimmed := bytes.TrimSpace(data)
	if len(trimmed) > 0 && (trimmed[0] == '{' || trimmed[0] == '[') {
		return "application/json"
	}
	return "application/x-yaml"
}
