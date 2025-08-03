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
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type (
	Params struct {
		Context  context.Context `schema:"-" json:"-"`
		Headers  http.Header     `schema:"-" json:"-"`
		NoAuth   bool            `schema:"-" json:"-"`
		Expand   []string        `schema:"expand,omitempty" json:"expand,omitempty"`
		Fields   []string        `schema:"fields,omitempty" json:"fields,omitempty"`
		Instance *string         `schema:"instance,omitempty" json:"-"`
	}

	ListParams struct {
		Params
		Limit  *int64 `schema:"limit,omitempty"`
		Offset *int64 `schema:"offset,omitempty"`
	}

	ParamsContainer interface {
		MethodParams() validation.Validatable
		RequestParams() Params
	}

	ParamsProxy[T validation.Validatable] struct {
		requestParams Params
		methodParams  T
	}

	ResponseProxy[T any] struct {
		Resource[T]
	}

	ResponseSliceProxy[T any] struct {
		Resource[[]T]
	}

	ClientParamsKey string
)

const (
	clientParamsKey ClientParamsKey = "client_params"
)

func ContextWithParams(ctx context.Context, params Params) context.Context {
	return context.WithValue(ctx, clientParamsKey, params)
}

func ParamsFromContext(ctx context.Context) Params {
	var params Params

	if v, ok := ctx.Value(clientParamsKey).(Params); ok {
		params = v
	}

	return params
}

func NewParamsProxy[T validation.Validatable](ctx context.Context, methodParams T) ParamsContainer {
	return &ParamsProxy[T]{
		requestParams: ParamsFromContext(ctx),
		methodParams:  methodParams,
	}
}

func (p *ParamsProxy[T]) RequestParams() Params {
	return p.requestParams
}

func (p *ParamsProxy[T]) MethodParams() validation.Validatable {
	return p.methodParams
}

func (p *ParamsProxy[T]) MarshalJSON() ([]byte, error) {
	// Marshal methodParams to get its JSON representation
	methodData, err := json.Marshal(p.methodParams)
	if err != nil {
		return nil, err
	}

	// Unmarshal into a map to flatten the structure
	var methodMap map[string]interface{}
	if err := json.Unmarshal(methodData, &methodMap); err != nil {
		return nil, err
	}

	// Marshal requestParams to get its JSON representation
	requestData, err := json.Marshal(p.requestParams)
	if err != nil {
		return nil, err
	}

	// Unmarshal into a map to flatten the structure
	var requestMap map[string]interface{}
	if err := json.Unmarshal(requestData, &requestMap); err != nil {
		return nil, err
	}

	// Merge the maps, with requestParams taking precedence for overlapping keys
	for k, v := range requestMap {
		methodMap[k] = v
	}

	return json.Marshal(methodMap)
}
