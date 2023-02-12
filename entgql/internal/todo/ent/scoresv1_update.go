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
	"errors"
	"fmt"

	"entgo.io/contrib/entgql/internal/todo/ent/predicate"
	"entgo.io/contrib/entgql/internal/todo/ent/scores"
	"entgo.io/contrib/entgql/internal/todo/ent/scoresv1"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ScoresV1Update is the builder for updating ScoresV1 entities.
type ScoresV1Update struct {
	config
	hooks    []Hook
	mutation *ScoresV1Mutation
}

// Where appends a list predicates to the ScoresV1Update builder.
func (sv *ScoresV1Update) Where(ps ...predicate.ScoresV1) *ScoresV1Update {
	sv.mutation.Where(ps...)
	return sv
}

// SetScore sets the "score" field.
func (sv *ScoresV1Update) SetScore(i int) *ScoresV1Update {
	sv.mutation.ResetScore()
	sv.mutation.SetScore(i)
	return sv
}

// AddScore adds i to the "score" field.
func (sv *ScoresV1Update) AddScore(i int) *ScoresV1Update {
	sv.mutation.AddScore(i)
	return sv
}

// AddScoreIDs adds the "Scores" edge to the Scores entity by IDs.
func (sv *ScoresV1Update) AddScoreIDs(ids ...int) *ScoresV1Update {
	sv.mutation.AddScoreIDs(ids...)
	return sv
}

// AddScores adds the "Scores" edges to the Scores entity.
func (sv *ScoresV1Update) AddScores(s ...*Scores) *ScoresV1Update {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sv.AddScoreIDs(ids...)
}

// Mutation returns the ScoresV1Mutation object of the builder.
func (sv *ScoresV1Update) Mutation() *ScoresV1Mutation {
	return sv.mutation
}

// ClearScores clears all "Scores" edges to the Scores entity.
func (sv *ScoresV1Update) ClearScores() *ScoresV1Update {
	sv.mutation.ClearScores()
	return sv
}

// RemoveScoreIDs removes the "Scores" edge to Scores entities by IDs.
func (sv *ScoresV1Update) RemoveScoreIDs(ids ...int) *ScoresV1Update {
	sv.mutation.RemoveScoreIDs(ids...)
	return sv
}

// RemoveScores removes "Scores" edges to Scores entities.
func (sv *ScoresV1Update) RemoveScores(s ...*Scores) *ScoresV1Update {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sv.RemoveScoreIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sv *ScoresV1Update) Save(ctx context.Context) (int, error) {
	return withHooks[int, ScoresV1Mutation](ctx, sv.sqlSave, sv.mutation, sv.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (sv *ScoresV1Update) SaveX(ctx context.Context) int {
	affected, err := sv.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sv *ScoresV1Update) Exec(ctx context.Context) error {
	_, err := sv.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sv *ScoresV1Update) ExecX(ctx context.Context) {
	if err := sv.Exec(ctx); err != nil {
		panic(err)
	}
}

func (sv *ScoresV1Update) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   scoresv1.Table,
			Columns: scoresv1.Columns,
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
	if value, ok := sv.mutation.Score(); ok {
		_spec.SetField(scoresv1.FieldScore, field.TypeInt, value)
	}
	if value, ok := sv.mutation.AddedScore(); ok {
		_spec.AddField(scoresv1.FieldScore, field.TypeInt, value)
	}
	if sv.mutation.ScoresCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   scoresv1.ScoresTable,
			Columns: []string{scoresv1.ScoresColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: scores.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sv.mutation.RemovedScoresIDs(); len(nodes) > 0 && !sv.mutation.ScoresCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   scoresv1.ScoresTable,
			Columns: []string{scoresv1.ScoresColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: scores.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sv.mutation.ScoresIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   scoresv1.ScoresTable,
			Columns: []string{scoresv1.ScoresColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: scores.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, sv.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{scoresv1.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	sv.mutation.done = true
	return n, nil
}

// ScoresV1UpdateOne is the builder for updating a single ScoresV1 entity.
type ScoresV1UpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ScoresV1Mutation
}

// SetScore sets the "score" field.
func (svo *ScoresV1UpdateOne) SetScore(i int) *ScoresV1UpdateOne {
	svo.mutation.ResetScore()
	svo.mutation.SetScore(i)
	return svo
}

// AddScore adds i to the "score" field.
func (svo *ScoresV1UpdateOne) AddScore(i int) *ScoresV1UpdateOne {
	svo.mutation.AddScore(i)
	return svo
}

// AddScoreIDs adds the "Scores" edge to the Scores entity by IDs.
func (svo *ScoresV1UpdateOne) AddScoreIDs(ids ...int) *ScoresV1UpdateOne {
	svo.mutation.AddScoreIDs(ids...)
	return svo
}

// AddScores adds the "Scores" edges to the Scores entity.
func (svo *ScoresV1UpdateOne) AddScores(s ...*Scores) *ScoresV1UpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return svo.AddScoreIDs(ids...)
}

// Mutation returns the ScoresV1Mutation object of the builder.
func (svo *ScoresV1UpdateOne) Mutation() *ScoresV1Mutation {
	return svo.mutation
}

// ClearScores clears all "Scores" edges to the Scores entity.
func (svo *ScoresV1UpdateOne) ClearScores() *ScoresV1UpdateOne {
	svo.mutation.ClearScores()
	return svo
}

// RemoveScoreIDs removes the "Scores" edge to Scores entities by IDs.
func (svo *ScoresV1UpdateOne) RemoveScoreIDs(ids ...int) *ScoresV1UpdateOne {
	svo.mutation.RemoveScoreIDs(ids...)
	return svo
}

// RemoveScores removes "Scores" edges to Scores entities.
func (svo *ScoresV1UpdateOne) RemoveScores(s ...*Scores) *ScoresV1UpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return svo.RemoveScoreIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (svo *ScoresV1UpdateOne) Select(field string, fields ...string) *ScoresV1UpdateOne {
	svo.fields = append([]string{field}, fields...)
	return svo
}

// Save executes the query and returns the updated ScoresV1 entity.
func (svo *ScoresV1UpdateOne) Save(ctx context.Context) (*ScoresV1, error) {
	return withHooks[*ScoresV1, ScoresV1Mutation](ctx, svo.sqlSave, svo.mutation, svo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (svo *ScoresV1UpdateOne) SaveX(ctx context.Context) *ScoresV1 {
	node, err := svo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (svo *ScoresV1UpdateOne) Exec(ctx context.Context) error {
	_, err := svo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (svo *ScoresV1UpdateOne) ExecX(ctx context.Context) {
	if err := svo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (svo *ScoresV1UpdateOne) sqlSave(ctx context.Context) (_node *ScoresV1, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   scoresv1.Table,
			Columns: scoresv1.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: scoresv1.FieldID,
			},
		},
	}
	id, ok := svo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ScoresV1.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := svo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, scoresv1.FieldID)
		for _, f := range fields {
			if !scoresv1.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != scoresv1.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := svo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := svo.mutation.Score(); ok {
		_spec.SetField(scoresv1.FieldScore, field.TypeInt, value)
	}
	if value, ok := svo.mutation.AddedScore(); ok {
		_spec.AddField(scoresv1.FieldScore, field.TypeInt, value)
	}
	if svo.mutation.ScoresCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   scoresv1.ScoresTable,
			Columns: []string{scoresv1.ScoresColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: scores.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := svo.mutation.RemovedScoresIDs(); len(nodes) > 0 && !svo.mutation.ScoresCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   scoresv1.ScoresTable,
			Columns: []string{scoresv1.ScoresColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: scores.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := svo.mutation.ScoresIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   scoresv1.ScoresTable,
			Columns: []string{scoresv1.ScoresColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: scores.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ScoresV1{config: svo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, svo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{scoresv1.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	svo.mutation.done = true
	return _node, nil
}
