// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ent

import (
	"context"
	"entgo.io/contrib/entgql/internal/todo/ent/scores"
	"entgo.io/contrib/entgql/internal/todo/ent/scoresv1"
	"entgo.io/contrib/entgql/internal/todo/ent/todo"
	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
)

// A copy of the generated code that we can modify, as an attempted workaround for the issue
// of not being able to sort by edge fields in paginated queries (related: https://github.com/ent/ent/issues/2807).
func (t *TodoQuery) PaginateExt(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...TodoPaginateOption,
) (*TodoConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newTodoPager(opts)
	if err != nil {
		return nil, err
	}
	if t, err = pager.applyFilter(t); err != nil {
		return nil, err
	}
	conn := &TodoConnection{Edges: []*TodoEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = t.Clone().Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}

	t = pager.applyCursors(t, after, before)
	t = applyScoresOrderExt(t, pager, last != nil)
	if limit := paginateLimit(first, last); limit != 0 {
		t.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := t.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}

	nodes, err := t.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

func applyScoresOrderExt(v *TodoQuery, pager *todoPager, reverse bool) *TodoQuery {
	if pager.order.Field.String() != "V1_SCORE" {
		// Not out customer order
		return pager.applyOrder(v, reverse)
	}

	// This is far from ideal since we are joining on these tables in order to do the sorting,
	//	but then ignore it and do another query anyway for fetching the info from telated tables.
	// Still this allows us to work around our original sorting issue (sorting by any edge field).
	// We still have the issue that we can't sort by more than one edge field.
	v1Table := sql.Table(scores.ScoresV1InverseTable)
	v.Where(func(s *sql.Selector) {
		todoTable := sql.Table(scores.TodoInverseTable)
		scoresTable := sql.Table(todo.ScoresInverseTable)
		s.Join(scoresTable).On(
			todoTable.C(todo.ScoresColumn),
			scoresTable.C(scores.FieldID),
		)
		s.Join(v1Table).On(
			scoresTable.C(scores.ScoresV1Column),
			v1Table.C(todo.FieldID),
		)

	})
	v.Unique(false).Order(
		func(s *sql.Selector) {
			expression := v1Table.C(scoresv1.FieldScore) + " " + pager.order.Direction.String()
			s.OrderExpr(sql.Expr(expression))
		},
	)

	return v
}
