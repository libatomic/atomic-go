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

import "golang.org/x/oauth2"

type (
	ClientConfig struct {
		OauthConfig oauth2.Config
		APIHost     string
	}

	Client struct {
		c ClientConfig
	}

	Option func(*ClientConfig)
)

const (
	DefaultAPIHost = "https://api.passport.online"
)

func NewClient(opts ...Option) *Client {
	c := &Client{
		ClientConfig{
			OauthConfig: oauth2.Config{
				Endpoint: DefaultOauthEndpoint,
			},
			APIHost: DefaultAPIHost,
		},
	}

	for _, opt := range opts {
		opt(&c.c)
	}
	return c
}

func WithAPIHost(host string) Option {
	return func(c *ClientConfig) {
		c.APIHost = host
	}
}

func WithOauthEndPoint(ep oauth2.Endpoint) Option {
	return func(c *ClientConfig) {
		c.OauthConfig.Endpoint = ep
	}
}

func WithOauthHost(host string) Option {
	return func(c *ClientConfig) {
		c.OauthConfig.Endpoint.AuthURL = host + "/oauth/authorize"
		c.OauthConfig.Endpoint.TokenURL = host + "/oauth/token"
	}
}

func WithOauthClientCredentials(id, secret string) Option {
	return func(c *ClientConfig) {
		c.OauthConfig.ClientID = id
		c.OauthConfig.ClientSecret = secret
	}
}
