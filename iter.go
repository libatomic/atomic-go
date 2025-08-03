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

type (
	Iter[P any, T any] struct {
		cur    T
		list   []T
		params P
		next   NextFunc[P, T]
	}

	NextFunc[P any, T any] func(P) (T, bool)
)

func NewIter[P any, T any](params P, n NextFunc[P, T]) *Iter[P, T] {
	return &Iter[P, T]{
		list:   make([]T, 0),
		params: params,
		next:   n,
	}
}
