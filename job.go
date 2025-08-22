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
	Job             = atomic.Job
	JobGetInput     = atomic.JobGetInput
	JobCreateInput  = atomic.JobCreateInput
	JobUpdateInput  = atomic.JobUpdateInput
	JobListInput    = atomic.JobListInput
	JobRestartInput = atomic.JobRestartInput
	JobCancelInput  = atomic.JobCancelInput
)

const (
	JobGetPath     = "/api/1.0.0/jobs/%s"
	JobCreatePath  = "/api/1.0.0/jobs"
	JobRestartPath = "/api/1.0.0/jobs/%s"
	JobUpdatePath  = "/api/1.0.0/jobs/%s"
	JobCancelPath  = "/api/1.0.0/jobs/%s"
	JobListPath    = "/api/1.0.0/jobs"
)

func (c *Client) JobCreate(ctx context.Context, params *JobCreateInput) (*Job, error) {
	var resp ResponseProxy[Job]

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, JobCreatePath, params).Post(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) JobGet(ctx context.Context, params *JobGetInput) (*Job, error) {
	var resp ResponseProxy[Job]

	path := fmt.Sprintf(JobGetPath, params.JobID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Get(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) JobUpdate(ctx context.Context, params *JobUpdateInput) (*Job, error) {
	var resp ResponseProxy[Job]

	path := fmt.Sprintf(JobUpdatePath, params.JobID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Put(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) JobList(ctx context.Context, params *JobListInput) ([]*Job, error) {
	var resp ResponseProxy[[]*Job]

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, JobListPath, params).Get(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Value(), nil
}

func (c *Client) JobRestart(ctx context.Context, params *JobRestartInput) (*Job, error) {
	var resp ResponseProxy[Job]

	path := fmt.Sprintf(JobRestartPath, params.JobID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Post(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Pointer(), nil
}

func (c *Client) JobCancel(ctx context.Context, params *JobCancelInput) error {
	var resp ResponseProxy[any]

	path := fmt.Sprintf(JobCancelPath, params.JobID.String())

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, path, params).Delete(),
		&resp); err != nil {
		return err
	}

	return nil
}
