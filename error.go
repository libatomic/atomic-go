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

import "fmt"

type (
	Error struct {
		Code    string `json:"code,omitempty"`
		Message string `json:"message,omitempty"`
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
	switch {
	case e.Message != "" && e.Code != "":
		return fmt.Sprintf("%s (%s)", e.Message, e.Code)
	case e.Message != "":
		return e.Message
	case e.Code != "" && e.Raw != "":
		return fmt.Sprintf("%s: %s", e.Code, e.Raw)
	case e.Code != "":
		return e.Code
	case e.Raw != "":
		return e.Raw
	case e.Status != "":
		return e.Status
	default:
		return "unknown error"
	}
}
