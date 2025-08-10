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
	"io"
	"net/http"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/go-querystring/query"
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

	RequestContainer interface {
		MethodParams() validation.Validatable
		RequestParams() Params
		ParamsEncoding() ParamsEncoding
		ContentType() string
		Method() string
		Path() string
		Body() io.Reader
	}

	RequestProxy[T validation.Validatable] struct {
		requestParams Params
		methodParams  T
		method        string
		encoding      ParamsEncoding
		body          io.Reader
		path          string
		contentType   string
	}

	ResponseProxy[T any] struct {
		Resource[T]
	}

	ResponseSliceProxy[T any] struct {
		Resource[[]T]
	}

	ClientParamsKey string

	ParamsEncoding string
)

const (
	clientParamsKey ClientParamsKey = "client_params"

	ParamsEncodingJSON      ParamsEncoding = "json"
	ParamsEncodingMultipart ParamsEncoding = "multipart"
	ParamsEncodingQuery     ParamsEncoding = "query"
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

func NewRequest[T validation.Validatable](ctx context.Context, path string, methodParams T) *RequestProxy[T] {
	return &RequestProxy[T]{
		requestParams: ParamsFromContext(ctx),
		methodParams:  methodParams,
		method:        http.MethodGet,
		path:          path,
		contentType:   "application/json",
		encoding:      ParamsEncodingQuery,
		body:          nil,
	}
}

func (p *RequestProxy[T]) WithMethod(method string) *RequestProxy[T] {
	p.method = method
	return p
}

func (p *RequestProxy[T]) Get() *RequestProxy[T] {
	p.method = http.MethodGet
	p.encoding = ParamsEncodingQuery
	return p
}

func (p *RequestProxy[T]) Post() *RequestProxy[T] {
	p.method = http.MethodPost
	p.encoding = ParamsEncodingJSON
	return p
}

func (p *RequestProxy[T]) Patch() *RequestProxy[T] {
	p.method = http.MethodPatch
	p.encoding = ParamsEncodingJSON
	return p
}

func (p *RequestProxy[T]) Put() *RequestProxy[T] {
	p.method = http.MethodPut
	p.encoding = ParamsEncodingJSON
	return p
}

func (p *RequestProxy[T]) Delete() *RequestProxy[T] {
	p.method = http.MethodDelete
	p.encoding = ParamsEncodingQuery
	return p
}

func (p *RequestProxy[T]) WithContentType(contentType string) *RequestProxy[T] {
	p.contentType = contentType

	switch contentType {
	case "application/json":
		p.encoding = ParamsEncodingJSON
	case "multipart/form-data":
		p.encoding = ParamsEncodingMultipart
	}

	if strings.HasPrefix(contentType, "multipart/form-data") {
		p.encoding = ParamsEncodingMultipart
	}

	return p
}

func (p *RequestProxy[T]) WithEncoding(encoding ParamsEncoding) *RequestProxy[T] {
	p.encoding = encoding
	return p
}

func (p *RequestProxy[T]) WithBody(body io.Reader) *RequestProxy[T] {
	p.body = body
	return p
}

func (p *RequestProxy[T]) RequestParams() Params {
	return p.requestParams
}

func (p *RequestProxy[T]) MethodParams() validation.Validatable {
	return p.methodParams
}

func (p *RequestProxy[T]) Method() string {
	return p.method
}

func (p *RequestProxy[T]) Path() string {
	if p.encoding == ParamsEncodingQuery {
		values, err := query.Values(p.methodParams)
		if err != nil {
			return p.path
		}

		if len(values) > 0 {
			return p.path + "?" + values.Encode()
		}
	}

	return p.path
}

func (p *RequestProxy[T]) Body() io.Reader {
	if p.body != nil {
		return p.body
	}

	if p.encoding == ParamsEncodingJSON {
		data, err := json.Marshal(p.methodParams)
		if err != nil {
			return nil
		}
		return bytes.NewReader(data)
	}

	return nil
}

func (p *RequestProxy[T]) ContentType() string {
	if p.requestParams.Headers.Get("Content-Type") != "" {
		return p.requestParams.Headers.Get("Content-Type")
	}

	return p.contentType
}

func (p *RequestProxy[T]) ParamsEncoding() ParamsEncoding {
	return p.encoding
}

func (p *RequestProxy[T]) MarshalJSON() ([]byte, error) {
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
