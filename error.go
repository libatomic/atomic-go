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
	"fmt"
	"sort"
	"strings"

	"github.com/libatomic/atomic/v2/pkg/atomic"
)

type (
	Error struct {
		Code    string `json:"code,omitempty"`
		Message string `json:"message,omitempty"`
		// Detail carries per-field validation messages the server attaches to
		// request_validation errors (e.g. {"login": "must be a valid email"}).
		Detail map[string]any `json:"detail,omitempty"`
		// Cause is the server-side cause string sent with 5xx responses.
		Cause string `json:"cause,omitempty"`
		// Status carries the HTTP status line (e.g. "401 Unauthorized") so
		// callers still get useful output when the server returns a body that
		// doesn't include a Message.
		Status string `json:"-"`
		// Raw is the original response body when it existed but couldn't be
		// decoded into the standard shape. Helps surface non-standard error
		// envelopes (e.g. oauth2 responses).
		Raw string `json:"-"`
	}
)

// Error never returns an empty string so wrappers like fmt.Errorf("x: %w", e)
// always show something useful.
func (e Error) Error() string {
	var base string
	switch {
	case e.Message != "" && e.Code != "":
		base = fmt.Sprintf("%s (%s)", e.Message, e.Code)
	case e.Message != "":
		base = e.Message
	case e.Code != "" && e.Raw != "":
		base = fmt.Sprintf("%s: %s", e.Code, e.Raw)
	case e.Code != "":
		base = e.Code
	case e.Raw != "":
		base = e.Raw
	case e.Status != "":
		base = e.Status
	default:
		base = "unknown error"
	}

	if d := e.detailString(); d != "" {
		base += ": " + d
	}
	if e.Cause != "" {
		base += " (cause: " + e.Cause + ")"
	}
	return base
}

// detailString renders Detail as "field: message; field: message" in a
// stable order so callers (and agents) see why a request was rejected.
func (e Error) detailString() string {
	if len(e.Detail) == 0 {
		return ""
	}
	keys := make([]string, 0, len(e.Detail))
	for k := range e.Detail {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	parts := make([]string, 0, len(keys))
	for _, k := range keys {
		parts = append(parts, fmt.Sprintf("%s: %v", k, e.Detail[k]))
	}
	return strings.Join(parts, "; ")
}

// idString renders a path ID without dereferencing a nil pointer, so a
// missing required ID yields a clear 4xx from the server rather than a
// client-side panic.
func idString(v any) string {
	switch x := v.(type) {
	case nil:
		return ""
	case *atomic.ID:
		if x == nil {
			return ""
		}
		return x.String()
	case atomic.ID:
		return x.String()
	case fmt.Stringer:
		return x.String()
	default:
		return fmt.Sprint(x)
	}
}
