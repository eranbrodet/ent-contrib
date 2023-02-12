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
	"fmt"
	"strings"

	"entgo.io/contrib/entgql/internal/todo/ent/scoresv1"
	"entgo.io/ent/dialect/sql"
)

// ScoresV1 is the model entity for the ScoresV1 schema.
type ScoresV1 struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Score holds the value of the "score" field.
	Score int `json:"score,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ScoresV1Query when eager-loading is set.
	Edges ScoresV1Edges `json:"edges"`
}

// ScoresV1Edges holds the relations/edges for other nodes in the graph.
type ScoresV1Edges struct {
	// Scores holds the value of the Scores edge.
	Scores []*Scores `json:"Scores,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool

	namedScores map[string][]*Scores
}

// ScoresOrErr returns the Scores value or an error if the edge
// was not loaded in eager-loading.
func (e ScoresV1Edges) ScoresOrErr() ([]*Scores, error) {
	if e.loadedTypes[0] {
		return e.Scores, nil
	}
	return nil, &NotLoadedError{edge: "Scores"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ScoresV1) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case scoresv1.FieldID, scoresv1.FieldScore:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ScoresV1", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ScoresV1 fields.
func (s *ScoresV1) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case scoresv1.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case scoresv1.FieldScore:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field score", values[i])
			} else if value.Valid {
				s.Score = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryScores queries the "Scores" edge of the ScoresV1 entity.
func (s *ScoresV1) QueryScores() *ScoresQuery {
	return NewScoresV1Client(s.config).QueryScores(s)
}

// Update returns a builder for updating this ScoresV1.
// Note that you need to call ScoresV1.Unwrap() before calling this method if this ScoresV1
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *ScoresV1) Update() *ScoresV1UpdateOne {
	return NewScoresV1Client(s.config).UpdateOne(s)
}

// Unwrap unwraps the ScoresV1 entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *ScoresV1) Unwrap() *ScoresV1 {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: ScoresV1 is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *ScoresV1) String() string {
	var builder strings.Builder
	builder.WriteString("ScoresV1(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("score=")
	builder.WriteString(fmt.Sprintf("%v", s.Score))
	builder.WriteByte(')')
	return builder.String()
}

// NamedScores returns the Scores named value or an error if the edge was not
// loaded in eager-loading with this name.
func (s *ScoresV1) NamedScores(name string) ([]*Scores, error) {
	if s.Edges.namedScores == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := s.Edges.namedScores[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (s *ScoresV1) appendNamedScores(name string, edges ...*Scores) {
	if s.Edges.namedScores == nil {
		s.Edges.namedScores = make(map[string][]*Scores)
	}
	if len(edges) == 0 {
		s.Edges.namedScores[name] = []*Scores{}
	} else {
		s.Edges.namedScores[name] = append(s.Edges.namedScores[name], edges...)
	}
}

// ScoresV1s is a parsable slice of ScoresV1.
type ScoresV1s []*ScoresV1

func (s ScoresV1s) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
