// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/entc/integration/gremlin/ent/group"
	"entgo.io/ent/entc/integration/gremlin/ent/groupinfo"
	"entgo.io/ent/entc/integration/gremlin/ent/predicate"
)

// GroupInfoQuery is the builder for querying GroupInfo entities.
type GroupInfoQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.GroupInfo
	withGroups *GroupQuery
	// intermediate query (i.e. traversal path).
	gremlin *dsl.Traversal
	path    func(context.Context) (*dsl.Traversal, error)
}

// Where adds a new predicate for the GroupInfoQuery builder.
func (giq *GroupInfoQuery) Where(ps ...predicate.GroupInfo) *GroupInfoQuery {
	giq.predicates = append(giq.predicates, ps...)
	return giq
}

// Limit the number of records to be returned by this query.
func (giq *GroupInfoQuery) Limit(limit int) *GroupInfoQuery {
	giq.ctx.Limit = &limit
	return giq
}

// Offset to start from.
func (giq *GroupInfoQuery) Offset(offset int) *GroupInfoQuery {
	giq.ctx.Offset = &offset
	return giq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (giq *GroupInfoQuery) Unique(unique bool) *GroupInfoQuery {
	giq.ctx.Unique = &unique
	return giq
}

// Order specifies how the records should be ordered.
func (giq *GroupInfoQuery) Order(o ...OrderFunc) *GroupInfoQuery {
	giq.order = append(giq.order, o...)
	return giq
}

// QueryGroups chains the current query on the "groups" edge.
func (giq *GroupInfoQuery) QueryGroups() *GroupQuery {
	query := (&GroupClient{config: giq.config}).Query()
	query.path = func(ctx context.Context) (fromU *dsl.Traversal, err error) {
		if err := giq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		gremlin := giq.gremlinQuery(ctx)
		fromU = gremlin.InE(group.InfoLabel).OutV()
		return fromU, nil
	}
	return query
}

// First returns the first GroupInfo entity from the query.
// Returns a *NotFoundError when no GroupInfo was found.
func (giq *GroupInfoQuery) First(ctx context.Context) (*GroupInfo, error) {
	nodes, err := giq.Limit(1).All(setContextOp(ctx, giq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{groupinfo.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (giq *GroupInfoQuery) FirstX(ctx context.Context) *GroupInfo {
	node, err := giq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GroupInfo ID from the query.
// Returns a *NotFoundError when no GroupInfo ID was found.
func (giq *GroupInfoQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = giq.Limit(1).IDs(setContextOp(ctx, giq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{groupinfo.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (giq *GroupInfoQuery) FirstIDX(ctx context.Context) string {
	id, err := giq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GroupInfo entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one GroupInfo entity is found.
// Returns a *NotFoundError when no GroupInfo entities are found.
func (giq *GroupInfoQuery) Only(ctx context.Context) (*GroupInfo, error) {
	nodes, err := giq.Limit(2).All(setContextOp(ctx, giq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{groupinfo.Label}
	default:
		return nil, &NotSingularError{groupinfo.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (giq *GroupInfoQuery) OnlyX(ctx context.Context) *GroupInfo {
	node, err := giq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GroupInfo ID in the query.
// Returns a *NotSingularError when more than one GroupInfo ID is found.
// Returns a *NotFoundError when no entities are found.
func (giq *GroupInfoQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = giq.Limit(2).IDs(setContextOp(ctx, giq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{groupinfo.Label}
	default:
		err = &NotSingularError{groupinfo.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (giq *GroupInfoQuery) OnlyIDX(ctx context.Context) string {
	id, err := giq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GroupInfos.
func (giq *GroupInfoQuery) All(ctx context.Context) ([]*GroupInfo, error) {
	ctx = setContextOp(ctx, giq.ctx, "All")
	if err := giq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*GroupInfo, *GroupInfoQuery]()
	return withInterceptors[[]*GroupInfo](ctx, giq, qr, giq.inters)
}

// AllX is like All, but panics if an error occurs.
func (giq *GroupInfoQuery) AllX(ctx context.Context) []*GroupInfo {
	nodes, err := giq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GroupInfo IDs.
func (giq *GroupInfoQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	ctx = setContextOp(ctx, giq.ctx, "IDs")
	if err := giq.Select(groupinfo.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (giq *GroupInfoQuery) IDsX(ctx context.Context) []string {
	ids, err := giq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (giq *GroupInfoQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, giq.ctx, "Count")
	if err := giq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, giq, querierCount[*GroupInfoQuery](), giq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (giq *GroupInfoQuery) CountX(ctx context.Context) int {
	count, err := giq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (giq *GroupInfoQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, giq.ctx, "Exist")
	switch _, err := giq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (giq *GroupInfoQuery) ExistX(ctx context.Context) bool {
	exist, err := giq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GroupInfoQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (giq *GroupInfoQuery) Clone() *GroupInfoQuery {
	if giq == nil {
		return nil
	}
	return &GroupInfoQuery{
		config:     giq.config,
		ctx:        giq.ctx.Clone(),
		order:      append([]OrderFunc{}, giq.order...),
		inters:     append([]Interceptor{}, giq.inters...),
		predicates: append([]predicate.GroupInfo{}, giq.predicates...),
		withGroups: giq.withGroups.Clone(),
		// clone intermediate query.
		gremlin: giq.gremlin.Clone(),
		path:    giq.path,
	}
}

// WithGroups tells the query-builder to eager-load the nodes that are connected to
// the "groups" edge. The optional arguments are used to configure the query builder of the edge.
func (giq *GroupInfoQuery) WithGroups(opts ...func(*GroupQuery)) *GroupInfoQuery {
	query := (&GroupClient{config: giq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	giq.withGroups = query
	return giq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Desc string `json:"desc,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.GroupInfo.Query().
//		GroupBy(groupinfo.FieldDesc).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (giq *GroupInfoQuery) GroupBy(field string, fields ...string) *GroupInfoGroupBy {
	giq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &GroupInfoGroupBy{build: giq}
	grbuild.flds = &giq.ctx.Fields
	grbuild.label = groupinfo.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Desc string `json:"desc,omitempty"`
//	}
//
//	client.GroupInfo.Query().
//		Select(groupinfo.FieldDesc).
//		Scan(ctx, &v)
func (giq *GroupInfoQuery) Select(fields ...string) *GroupInfoSelect {
	giq.ctx.Fields = append(giq.ctx.Fields, fields...)
	sbuild := &GroupInfoSelect{GroupInfoQuery: giq}
	sbuild.label = groupinfo.Label
	sbuild.flds, sbuild.scan = &giq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a GroupInfoSelect configured with the given aggregations.
func (giq *GroupInfoQuery) Aggregate(fns ...AggregateFunc) *GroupInfoSelect {
	return giq.Select().Aggregate(fns...)
}

func (giq *GroupInfoQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range giq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, giq); err != nil {
				return err
			}
		}
	}
	if giq.path != nil {
		prev, err := giq.path(ctx)
		if err != nil {
			return err
		}
		giq.gremlin = prev
	}
	return nil
}

func (giq *GroupInfoQuery) gremlinAll(ctx context.Context, hooks ...queryHook) ([]*GroupInfo, error) {
	res := &gremlin.Response{}
	traversal := giq.gremlinQuery(ctx)
	if len(giq.ctx.Fields) > 0 {
		fields := make([]any, len(giq.ctx.Fields))
		for i, f := range giq.ctx.Fields {
			fields[i] = f
		}
		traversal.ValueMap(fields...)
	} else {
		traversal.ValueMap(true)
	}
	query, bindings := traversal.Query()
	if err := giq.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	var gis GroupInfos
	if err := gis.FromResponse(res); err != nil {
		return nil, err
	}
	gis.config(giq.config)
	return gis, nil
}

func (giq *GroupInfoQuery) gremlinCount(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := giq.gremlinQuery(ctx).Count().Query()
	if err := giq.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (giq *GroupInfoQuery) gremlinQuery(context.Context) *dsl.Traversal {
	v := g.V().HasLabel(groupinfo.Label)
	if giq.gremlin != nil {
		v = giq.gremlin.Clone()
	}
	for _, p := range giq.predicates {
		p(v)
	}
	if len(giq.order) > 0 {
		v.Order()
		for _, p := range giq.order {
			p(v)
		}
	}
	switch limit, offset := giq.ctx.Limit, giq.ctx.Offset; {
	case limit != nil && offset != nil:
		v.Range(*offset, *offset+*limit)
	case offset != nil:
		v.Range(*offset, math.MaxInt32)
	case limit != nil:
		v.Limit(*limit)
	}
	if unique := giq.ctx.Unique; unique == nil || *unique {
		v.Dedup()
	}
	return v
}

// GroupInfoGroupBy is the group-by builder for GroupInfo entities.
type GroupInfoGroupBy struct {
	selector
	build *GroupInfoQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gigb *GroupInfoGroupBy) Aggregate(fns ...AggregateFunc) *GroupInfoGroupBy {
	gigb.fns = append(gigb.fns, fns...)
	return gigb
}

// Scan applies the selector query and scans the result into the given value.
func (gigb *GroupInfoGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gigb.build.ctx, "GroupBy")
	if err := gigb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GroupInfoQuery, *GroupInfoGroupBy](ctx, gigb.build, gigb, gigb.build.inters, v)
}

func (gigb *GroupInfoGroupBy) gremlinScan(ctx context.Context, root *GroupInfoQuery, v any) error {
	var (
		trs   []any
		names []any
	)
	for _, fn := range gigb.fns {
		name, tr := fn("p", "")
		trs = append(trs, tr)
		names = append(names, name)
	}
	for _, f := range *gigb.flds {
		names = append(names, f)
		trs = append(trs, __.As("p").Unfold().Values(f).As(f))
	}
	query, bindings := root.gremlinQuery(ctx).Group().
		By(__.Values(*gigb.flds...).Fold()).
		By(__.Fold().Match(trs...).Select(names...)).
		Select(dsl.Values).
		Next().
		Query()
	res := &gremlin.Response{}
	if err := gigb.build.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(*gigb.flds)+len(gigb.fns) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}

// GroupInfoSelect is the builder for selecting fields of GroupInfo entities.
type GroupInfoSelect struct {
	*GroupInfoQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (gis *GroupInfoSelect) Aggregate(fns ...AggregateFunc) *GroupInfoSelect {
	gis.fns = append(gis.fns, fns...)
	return gis
}

// Scan applies the selector query and scans the result into the given value.
func (gis *GroupInfoSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gis.ctx, "Select")
	if err := gis.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GroupInfoQuery, *GroupInfoSelect](ctx, gis.GroupInfoQuery, gis, gis.inters, v)
}

func (gis *GroupInfoSelect) gremlinScan(ctx context.Context, root *GroupInfoQuery, v any) error {
	var (
		res       = &gremlin.Response{}
		traversal = root.gremlinQuery(ctx)
	)
	if fields := gis.ctx.Fields; len(fields) == 1 {
		if fields[0] != groupinfo.FieldID {
			traversal = traversal.Values(fields...)
		} else {
			traversal = traversal.ID()
		}
	} else {
		fields := make([]any, len(gis.ctx.Fields))
		for i, f := range gis.ctx.Fields {
			fields[i] = f
		}
		traversal = traversal.ValueMap(fields...)
	}
	query, bindings := traversal.Query()
	if err := gis.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(root.ctx.Fields) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}
