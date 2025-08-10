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
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/oauth2/clientcredentials"
)

type (
	ApiConfig struct {
		AccessToken string
		Host        string
		http        *http.Client
	}

	ApiBackend struct {
		c ApiConfig
	}

	ApiOption func(*ApiConfig)
)

const (
	DefaultAPIHost = "http://localhost:9000"
)

func New(opts ...ApiOption) *Client {
	b := &ApiBackend{
		ApiConfig{
			Host: DefaultAPIHost,
			http: http.DefaultClient,
		},
	}

	for _, opt := range opts {
		opt(&b.c)
	}

	return NewClient(b)
}

func WithHost(host string) ApiOption {
	return func(c *ApiConfig) {
		c.Host = host
	}
}

func WithToken(token string) ApiOption {
	return func(c *ApiConfig) {
		c.AccessToken = token
	}
}

func WithHTTPClient(http *http.Client) ApiOption {
	return func(c *ApiConfig) {
		c.http = http
	}
}

func WithClientCredentials(clientID, clientSecret string, scopes ...string) ApiOption {
	return func(c *ApiConfig) {
		cc := clientcredentials.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Scopes:       scopes,
			TokenURL:     "https://" + c.Host + "/oauth/token",
		}

		c.http = cc.Client(context.Background())
	}
}

func (b *ApiBackend) ExecContext(ctx context.Context, params RequestContainer, result Responder) error {
	req, err := b.NewRequest(ctx, params)
	if err != nil {
		return err
	}

	resp, err := b.c.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 400 {
		if len(body) > 0 {
			var e Error

			if err := json.NewDecoder(bytes.NewReader(body)).Decode(&e); err != nil {
				return err
			}
			return e
		}

		return errors.New(resp.Status)
	}

	if len(body) > 0 {
		if err := json.NewDecoder(bytes.NewReader(body)).Decode(result.Response()); err != nil {
			return err
		}
	}

	return nil
}

func (b *ApiBackend) NewRequest(ctx context.Context, params RequestContainer) (*http.Request, error) {
	reqParams := params.RequestParams()

	path := fmt.Sprintf("https://%s/%s", b.c.Host, strings.TrimPrefix(params.Path(), "/"))

	req, err := http.NewRequest(params.Method(), path, params.Body())
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", params.ContentType())

	authorization := "Bearer " + b.c.AccessToken

	if params != nil {
		if reqParams.Context != nil {
			req = req.WithContext(reqParams.Context)
		}

		if reqParams.Instance != nil {
			req.Header.Add("Atomic-Instance", strings.TrimSpace(*reqParams.Instance))
		}

		for k, v := range reqParams.Headers {
			for _, line := range v {
				// Use Set to override the default value possibly set before
				req.Header.Set(k, line)
			}
		}

		if !reqParams.NoAuth && b.c.AccessToken != "" {
			req.Header.Add("Authorization", authorization)
		}
	} else if b.c.AccessToken != "" {
		req.Header.Add("Authorization", authorization)
	}

	return req, nil
}
