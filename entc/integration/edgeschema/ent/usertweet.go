// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/edgeschema/ent/tweet"
	"entgo.io/ent/entc/integration/edgeschema/ent/user"
	"entgo.io/ent/entc/integration/edgeschema/ent/usertweet"
)

// UserTweet is the model entity for the UserTweet schema.
type UserTweet struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// TweetID holds the value of the "tweet_id" field.
	TweetID int `json:"tweet_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserTweetQuery when eager-loading is set.
	Edges UserTweetEdges `json:"edges"`
}

// UserTweetEdges holds the relations/edges for other nodes in the graph.
type UserTweetEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Tweet holds the value of the tweet edge.
	Tweet *Tweet `json:"tweet,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserTweetEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// TweetOrErr returns the Tweet value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserTweetEdges) TweetOrErr() (*Tweet, error) {
	if e.loadedTypes[1] {
		if e.Tweet == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: tweet.Label}
		}
		return e.Tweet, nil
	}
	return nil, &NotLoadedError{edge: "tweet"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserTweet) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case usertweet.FieldID, usertweet.FieldUserID, usertweet.FieldTweetID:
			values[i] = new(sql.NullInt64)
		case usertweet.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type UserTweet", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserTweet fields.
func (ut *UserTweet) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case usertweet.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ut.ID = int(value.Int64)
		case usertweet.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ut.CreatedAt = value.Time
			}
		case usertweet.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				ut.UserID = int(value.Int64)
			}
		case usertweet.FieldTweetID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tweet_id", values[i])
			} else if value.Valid {
				ut.TweetID = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the UserTweet entity.
func (ut *UserTweet) QueryUser() *UserQuery {
	return NewUserTweetClient(ut.config).QueryUser(ut)
}

// QueryTweet queries the "tweet" edge of the UserTweet entity.
func (ut *UserTweet) QueryTweet() *TweetQuery {
	return NewUserTweetClient(ut.config).QueryTweet(ut)
}

// Update returns a builder for updating this UserTweet.
// Note that you need to call UserTweet.Unwrap() before calling this method if this UserTweet
// was returned from a transaction, and the transaction was committed or rolled back.
func (ut *UserTweet) Update() *UserTweetUpdateOne {
	return NewUserTweetClient(ut.config).UpdateOne(ut)
}

// Unwrap unwraps the UserTweet entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ut *UserTweet) Unwrap() *UserTweet {
	_tx, ok := ut.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserTweet is not a transactional entity")
	}
	ut.config.driver = _tx.drv
	return ut
}

// String implements the fmt.Stringer.
func (ut *UserTweet) String() string {
	var builder strings.Builder
	builder.WriteString("UserTweet(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ut.ID))
	builder.WriteString("created_at=")
	builder.WriteString(ut.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", ut.UserID))
	builder.WriteString(", ")
	builder.WriteString("tweet_id=")
	builder.WriteString(fmt.Sprintf("%v", ut.TweetID))
	builder.WriteByte(')')
	return builder.String()
}

// UserTweets is a parsable slice of UserTweet.
type UserTweets []*UserTweet

func (ut UserTweets) config(cfg config) {
	for _i := range ut {
		ut[_i].config = cfg
	}
}
