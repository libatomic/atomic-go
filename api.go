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
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type (
	ApiConfig struct {
		Key  string
		Host string
	}

	ApiBackend struct {
		c ApiConfig
	}

	ApiOption func(*ApiConfig)
)

const (
	DefaultAPIHost = "https://api.passport.online"
)

func NewApiBackend(opts ...ApiOption) *ApiBackend {
	c := &ApiBackend{
		ApiConfig{
			Host: DefaultAPIHost,
		},
	}

	for _, opt := range opts {
		opt(&c.c)
	}
	return c
}

func WithHost(host string) ApiOption {
	return func(c *ApiConfig) {
		c.Host = host
	}
}

func WithKey(key string) ApiOption {
	return func(c *ApiConfig) {
		c.Key = key
	}
}

func (b *ApiBackend) ExecContext(ctx context.Context, method string, path string, params validation.Validatable, result Responder) error {

	return nil
}

func (b *ApiBackend) NewRequest(ctx context.Context, method, path, ct string, params *Params) (*http.Request, error) {
	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		return nil, err
	}

	authorization := "Bearer " + b.c.Key

	req.Header.Add("Authorization", authorization)
	req.Header.Add("Content-Type", ct)

	if params != nil {
		if params.Context != nil {
			req = req.WithContext(params.Context)
		}

		if params.Instance != nil {
			req.Header.Add("Atomic-Instance", strings.TrimSpace(*params.Instance))
		}

		for k, v := range params.Headers {
			for _, line := range v {
				// Use Set to override the default value possibly set before
				req.Header.Set(k, line)
			}
		}
	}

	return req, nil
}
