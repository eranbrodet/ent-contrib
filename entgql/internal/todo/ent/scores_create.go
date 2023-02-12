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
	"fmt"

	"entgo.io/contrib/entgql/internal/todo/ent/scores"
	"entgo.io/contrib/entgql/internal/todo/ent/scoresv1"
	"entgo.io/contrib/entgql/internal/todo/ent/scoresv2"
	"entgo.io/contrib/entgql/internal/todo/ent/todo"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ScoresCreate is the builder for creating a Scores entity.
type ScoresCreate struct {
	config
	mutation *ScoresMutation
	hooks    []Hook
}

// AddTodoIDs adds the "todo" edge to the Todo entity by IDs.
func (sc *ScoresCreate) AddTodoIDs(ids ...int) *ScoresCreate {
	sc.mutation.AddTodoIDs(ids...)
	return sc
}

// AddTodo adds the "todo" edges to the Todo entity.
func (sc *ScoresCreate) AddTodo(t ...*Todo) *ScoresCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return sc.AddTodoIDs(ids...)
}

// SetScoresV1ID sets the "ScoresV1" edge to the ScoresV1 entity by ID.
func (sc *ScoresCreate) SetScoresV1ID(id int) *ScoresCreate {
	sc.mutation.SetScoresV1ID(id)
	return sc
}

// SetNillableScoresV1ID sets the "ScoresV1" edge to the ScoresV1 entity by ID if the given value is not nil.
func (sc *ScoresCreate) SetNillableScoresV1ID(id *int) *ScoresCreate {
	if id != nil {
		sc = sc.SetScoresV1ID(*id)
	}
	return sc
}

// SetScoresV1 sets the "ScoresV1" edge to the ScoresV1 entity.
func (sc *ScoresCreate) SetScoresV1(s *ScoresV1) *ScoresCreate {
	return sc.SetScoresV1ID(s.ID)
}

// SetScoresV2ID sets the "ScoresV2" edge to the ScoresV2 entity by ID.
func (sc *ScoresCreate) SetScoresV2ID(id int) *ScoresCreate {
	sc.mutation.SetScoresV2ID(id)
	return sc
}

// SetNillableScoresV2ID sets the "ScoresV2" edge to the ScoresV2 entity by ID if the given value is not nil.
func (sc *ScoresCreate) SetNillableScoresV2ID(id *int) *ScoresCreate {
	if id != nil {
		sc = sc.SetScoresV2ID(*id)
	}
	return sc
}

// SetScoresV2 sets the "ScoresV2" edge to the ScoresV2 entity.
func (sc *ScoresCreate) SetScoresV2(s *ScoresV2) *ScoresCreate {
	return sc.SetScoresV2ID(s.ID)
}

// Mutation returns the ScoresMutation object of the builder.
func (sc *ScoresCreate) Mutation() *ScoresMutation {
	return sc.mutation
}

// Save creates the Scores in the database.
func (sc *ScoresCreate) Save(ctx context.Context) (*Scores, error) {
	return withHooks[*Scores, ScoresMutation](ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *ScoresCreate) SaveX(ctx context.Context) *Scores {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *ScoresCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *ScoresCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *ScoresCreate) check() error {
	return nil
}

func (sc *ScoresCreate) sqlSave(ctx context.Context) (*Scores, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *ScoresCreate) createSpec() (*Scores, *sqlgraph.CreateSpec) {
	var (
		_node = &Scores{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: scores.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: scores.FieldID,
			},
		}
	)
	if nodes := sc.mutation.TodoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   scores.TodoTable,
			Columns: []string{scores.TodoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: todo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.ScoresV1IDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   scores.ScoresV1Table,
			Columns: []string{scores.ScoresV1Column},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: scoresv1.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.scores_v1_scores = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.ScoresV2IDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   scores.ScoresV2Table,
			Columns: []string{scores.ScoresV2Column},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: scoresv2.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.scores_v2_scores = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ScoresCreateBulk is the builder for creating many Scores entities in bulk.
type ScoresCreateBulk struct {
	config
	builders []*ScoresCreate
}

// Save creates the Scores entities in the database.
func (scb *ScoresCreateBulk) Save(ctx context.Context) ([]*Scores, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Scores, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ScoresMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *ScoresCreateBulk) SaveX(ctx context.Context) []*Scores {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *ScoresCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *ScoresCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}