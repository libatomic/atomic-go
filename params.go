/*
 * This file is part of the Atomic Stack (https://github.com/libatomic/atomic).
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
	"net/http"
)

type (
	Params struct {
		Context  context.Context `schema:"-"`
		Headers  http.Header     `schema:"-"`
		Expand   []string        `schema:"expand,omitempty"`
		Fields   []string        `schema:"fields,omitempty"`
		Instance string          `schema:"instance,omitempty"`
	}

	ListParams struct {
		Params
		Limit  *int64 `schema:"limit,omitempty"`
		Offset *int64 `schema:"offset,omitempty"`
	}
)
