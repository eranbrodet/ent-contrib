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
//
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/contrib/entgql/internal/todo/ent/predicate"
	"entgo.io/contrib/entgql/internal/todo/ent/scoresv1"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ScoresV1Delete is the builder for deleting a ScoresV1 entity.
type ScoresV1Delete struct {
	config
	hooks    []Hook
	mutation *ScoresV1Mutation
}

// Where appends a list predicates to the ScoresV1Delete builder.
func (sv *ScoresV1Delete) Where(ps ...predicate.ScoresV1) *ScoresV1Delete {
	sv.mutation.Where(ps...)
	return sv
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sv *ScoresV1Delete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, ScoresV1Mutation](ctx, sv.sqlExec, sv.mutation, sv.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (sv *ScoresV1Delete) ExecX(ctx context.Context) int {
	n, err := sv.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sv *ScoresV1Delete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: scoresv1.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: scoresv1.FieldID,
			},
		},
	}
	if ps := sv.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, sv.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	sv.mutation.done = true
	return affected, err
}

// ScoresV1DeleteOne is the builder for deleting a single ScoresV1 entity.
type ScoresV1DeleteOne struct {
	sv *ScoresV1Delete
}

// Exec executes the deletion query.
func (svo *ScoresV1DeleteOne) Exec(ctx context.Context) error {
	n, err := svo.sv.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{scoresv1.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (svo *ScoresV1DeleteOne) ExecX(ctx context.Context) {
	svo.sv.ExecX(ctx)
}
