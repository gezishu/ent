// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/entc/integration/gremlin/ent/api"
	"entgo.io/ent/entc/integration/gremlin/ent/predicate"
)

// APIUpdate is the builder for updating Api entities.
type APIUpdate struct {
	config
	hooks    []Hook
	mutation *APIMutation
}

// Where appends a list predicates to the APIUpdate builder.
func (au *APIUpdate) Where(ps ...predicate.Api) *APIUpdate {
	au.mutation.Where(ps...)
	return au
}

// Mutation returns the APIMutation object of the builder.
func (au *APIUpdate) Mutation() *APIMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *APIUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, APIMutation](ctx, au.gremlinSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *APIUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *APIUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *APIUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *APIUpdate) gremlinSave(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := au.gremlin().Query()
	if err := au.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	if err, ok := isConstantError(res); ok {
		return 0, err
	}
	au.mutation.done = true
	return res.ReadInt()
}

func (au *APIUpdate) gremlin() *dsl.Traversal {
	v := g.V().HasLabel(api.Label)
	for _, p := range au.mutation.predicates {
		p(v)
	}
	var (
		trs []*dsl.Traversal
	)
	v.Count()
	trs = append(trs, v)
	return dsl.Join(trs...)
}

// APIUpdateOne is the builder for updating a single Api entity.
type APIUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *APIMutation
}

// Mutation returns the APIMutation object of the builder.
func (auo *APIUpdateOne) Mutation() *APIMutation {
	return auo.mutation
}

// Where appends a list predicates to the APIUpdate builder.
func (auo *APIUpdateOne) Where(ps ...predicate.Api) *APIUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *APIUpdateOne) Select(field string, fields ...string) *APIUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Api entity.
func (auo *APIUpdateOne) Save(ctx context.Context) (*Api, error) {
	return withHooks[*Api, APIMutation](ctx, auo.gremlinSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *APIUpdateOne) SaveX(ctx context.Context) *Api {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *APIUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *APIUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *APIUpdateOne) gremlinSave(ctx context.Context) (*Api, error) {
	res := &gremlin.Response{}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Api.id" for update`)}
	}
	query, bindings := auo.gremlin(id).Query()
	if err := auo.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	auo.mutation.done = true
	a := &Api{config: auo.config}
	if err := a.FromResponse(res); err != nil {
		return nil, err
	}
	return a, nil
}

func (auo *APIUpdateOne) gremlin(id string) *dsl.Traversal {
	v := g.V(id)
	var (
		trs []*dsl.Traversal
	)
	if len(auo.fields) > 0 {
		fields := make([]any, 0, len(auo.fields)+1)
		fields = append(fields, true)
		for _, f := range auo.fields {
			fields = append(fields, f)
		}
		v.ValueMap(fields...)
	} else {
		v.ValueMap(true)
	}
	trs = append(trs, v)
	return dsl.Join(trs...)
}
