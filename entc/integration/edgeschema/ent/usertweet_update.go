// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/edgeschema/ent/predicate"
	"entgo.io/ent/entc/integration/edgeschema/ent/tweet"
	"entgo.io/ent/entc/integration/edgeschema/ent/user"
	"entgo.io/ent/entc/integration/edgeschema/ent/usertweet"
	"entgo.io/ent/schema/field"
)

// UserTweetUpdate is the builder for updating UserTweet entities.
type UserTweetUpdate struct {
	config
	hooks    []Hook
	mutation *UserTweetMutation
}

// Where appends a list predicates to the UserTweetUpdate builder.
func (utu *UserTweetUpdate) Where(ps ...predicate.UserTweet) *UserTweetUpdate {
	utu.mutation.Where(ps...)
	return utu
}

// SetCreatedAt sets the "created_at" field.
func (utu *UserTweetUpdate) SetCreatedAt(t time.Time) *UserTweetUpdate {
	utu.mutation.SetCreatedAt(t)
	return utu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (utu *UserTweetUpdate) SetNillableCreatedAt(t *time.Time) *UserTweetUpdate {
	if t != nil {
		utu.SetCreatedAt(*t)
	}
	return utu
}

// SetUserID sets the "user_id" field.
func (utu *UserTweetUpdate) SetUserID(i int) *UserTweetUpdate {
	utu.mutation.SetUserID(i)
	return utu
}

// SetTweetID sets the "tweet_id" field.
func (utu *UserTweetUpdate) SetTweetID(i int) *UserTweetUpdate {
	utu.mutation.SetTweetID(i)
	return utu
}

// SetUser sets the "user" edge to the User entity.
func (utu *UserTweetUpdate) SetUser(u *User) *UserTweetUpdate {
	return utu.SetUserID(u.ID)
}

// SetTweet sets the "tweet" edge to the Tweet entity.
func (utu *UserTweetUpdate) SetTweet(t *Tweet) *UserTweetUpdate {
	return utu.SetTweetID(t.ID)
}

// Mutation returns the UserTweetMutation object of the builder.
func (utu *UserTweetUpdate) Mutation() *UserTweetMutation {
	return utu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (utu *UserTweetUpdate) ClearUser() *UserTweetUpdate {
	utu.mutation.ClearUser()
	return utu
}

// ClearTweet clears the "tweet" edge to the Tweet entity.
func (utu *UserTweetUpdate) ClearTweet() *UserTweetUpdate {
	utu.mutation.ClearTweet()
	return utu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (utu *UserTweetUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, UserTweetMutation](ctx, utu.sqlSave, utu.mutation, utu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (utu *UserTweetUpdate) SaveX(ctx context.Context) int {
	affected, err := utu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (utu *UserTweetUpdate) Exec(ctx context.Context) error {
	_, err := utu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (utu *UserTweetUpdate) ExecX(ctx context.Context) {
	if err := utu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (utu *UserTweetUpdate) check() error {
	if _, ok := utu.mutation.UserID(); utu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UserTweet.user"`)
	}
	if _, ok := utu.mutation.TweetID(); utu.mutation.TweetCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UserTweet.tweet"`)
	}
	return nil
}

func (utu *UserTweetUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := utu.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   usertweet.Table,
			Columns: usertweet.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: usertweet.FieldID,
			},
		},
	}
	if ps := utu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := utu.mutation.CreatedAt(); ok {
		_spec.SetField(usertweet.FieldCreatedAt, field.TypeTime, value)
	}
	if utu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   usertweet.UserTable,
			Columns: []string{usertweet.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := utu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   usertweet.UserTable,
			Columns: []string{usertweet.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if utu.mutation.TweetCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   usertweet.TweetTable,
			Columns: []string{usertweet.TweetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tweet.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := utu.mutation.TweetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   usertweet.TweetTable,
			Columns: []string{usertweet.TweetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tweet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, utu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usertweet.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	utu.mutation.done = true
	return n, nil
}

// UserTweetUpdateOne is the builder for updating a single UserTweet entity.
type UserTweetUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserTweetMutation
}

// SetCreatedAt sets the "created_at" field.
func (utuo *UserTweetUpdateOne) SetCreatedAt(t time.Time) *UserTweetUpdateOne {
	utuo.mutation.SetCreatedAt(t)
	return utuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (utuo *UserTweetUpdateOne) SetNillableCreatedAt(t *time.Time) *UserTweetUpdateOne {
	if t != nil {
		utuo.SetCreatedAt(*t)
	}
	return utuo
}

// SetUserID sets the "user_id" field.
func (utuo *UserTweetUpdateOne) SetUserID(i int) *UserTweetUpdateOne {
	utuo.mutation.SetUserID(i)
	return utuo
}

// SetTweetID sets the "tweet_id" field.
func (utuo *UserTweetUpdateOne) SetTweetID(i int) *UserTweetUpdateOne {
	utuo.mutation.SetTweetID(i)
	return utuo
}

// SetUser sets the "user" edge to the User entity.
func (utuo *UserTweetUpdateOne) SetUser(u *User) *UserTweetUpdateOne {
	return utuo.SetUserID(u.ID)
}

// SetTweet sets the "tweet" edge to the Tweet entity.
func (utuo *UserTweetUpdateOne) SetTweet(t *Tweet) *UserTweetUpdateOne {
	return utuo.SetTweetID(t.ID)
}

// Mutation returns the UserTweetMutation object of the builder.
func (utuo *UserTweetUpdateOne) Mutation() *UserTweetMutation {
	return utuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (utuo *UserTweetUpdateOne) ClearUser() *UserTweetUpdateOne {
	utuo.mutation.ClearUser()
	return utuo
}

// ClearTweet clears the "tweet" edge to the Tweet entity.
func (utuo *UserTweetUpdateOne) ClearTweet() *UserTweetUpdateOne {
	utuo.mutation.ClearTweet()
	return utuo
}

// Where appends a list predicates to the UserTweetUpdate builder.
func (utuo *UserTweetUpdateOne) Where(ps ...predicate.UserTweet) *UserTweetUpdateOne {
	utuo.mutation.Where(ps...)
	return utuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (utuo *UserTweetUpdateOne) Select(field string, fields ...string) *UserTweetUpdateOne {
	utuo.fields = append([]string{field}, fields...)
	return utuo
}

// Save executes the query and returns the updated UserTweet entity.
func (utuo *UserTweetUpdateOne) Save(ctx context.Context) (*UserTweet, error) {
	return withHooks[*UserTweet, UserTweetMutation](ctx, utuo.sqlSave, utuo.mutation, utuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (utuo *UserTweetUpdateOne) SaveX(ctx context.Context) *UserTweet {
	node, err := utuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (utuo *UserTweetUpdateOne) Exec(ctx context.Context) error {
	_, err := utuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (utuo *UserTweetUpdateOne) ExecX(ctx context.Context) {
	if err := utuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (utuo *UserTweetUpdateOne) check() error {
	if _, ok := utuo.mutation.UserID(); utuo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UserTweet.user"`)
	}
	if _, ok := utuo.mutation.TweetID(); utuo.mutation.TweetCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UserTweet.tweet"`)
	}
	return nil
}

func (utuo *UserTweetUpdateOne) sqlSave(ctx context.Context) (_node *UserTweet, err error) {
	if err := utuo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   usertweet.Table,
			Columns: usertweet.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: usertweet.FieldID,
			},
		},
	}
	id, ok := utuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserTweet.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := utuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usertweet.FieldID)
		for _, f := range fields {
			if !usertweet.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != usertweet.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := utuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := utuo.mutation.CreatedAt(); ok {
		_spec.SetField(usertweet.FieldCreatedAt, field.TypeTime, value)
	}
	if utuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   usertweet.UserTable,
			Columns: []string{usertweet.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := utuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   usertweet.UserTable,
			Columns: []string{usertweet.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if utuo.mutation.TweetCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   usertweet.TweetTable,
			Columns: []string{usertweet.TweetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tweet.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := utuo.mutation.TweetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   usertweet.TweetTable,
			Columns: []string{usertweet.TweetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tweet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &UserTweet{config: utuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, utuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usertweet.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	utuo.mutation.done = true
	return _node, nil
}
