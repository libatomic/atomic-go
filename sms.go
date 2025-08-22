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

	"github.com/libatomic/atomic/pkg/atomic"
)

const (
	SMSSendPath = "/api/1.0.0/sms"
)

func (c *Client) SendSMS(ctx context.Context, params *atomic.SendSMSInput) ([]*atomic.SMS, error) {
	var resp ResponseProxy[[]*atomic.SMS]

	if err := c.Backend.ExecContext(
		ctx,
		NewRequest(ctx, SMSSendPath, params).Post(),
		&resp); err != nil {
		return nil, err
	}

	return resp.Value(), nil
}
