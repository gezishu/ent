// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// FriendshipsColumns holds the columns for the "friendships" table.
	FriendshipsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "weight", Type: field.TypeInt, Default: 1},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeInt},
		{Name: "friend_id", Type: field.TypeInt},
	}
	// FriendshipsTable holds the schema information for the "friendships" table.
	FriendshipsTable = &schema.Table{
		Name:       "friendships",
		Columns:    FriendshipsColumns,
		PrimaryKey: []*schema.Column{FriendshipsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "friendships_users_user",
				Columns:    []*schema.Column{FriendshipsColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "friendships_users_friend",
				Columns:    []*schema.Column{FriendshipsColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "friendship_created_at",
				Unique:  false,
				Columns: []*schema.Column{FriendshipsColumns[2]},
			},
			{
				Name:    "friendships_edge",
				Unique:  true,
				Columns: []*schema.Column{FriendshipsColumns[3], FriendshipsColumns[4]},
			},
		},
	}
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Default: "Unknown"},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:       "groups",
		Columns:    GroupsColumns,
		PrimaryKey: []*schema.Column{GroupsColumns[0]},
	}
	// GroupTagsColumns holds the columns for the "group_tags" table.
	GroupTagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "tag_id", Type: field.TypeInt},
		{Name: "group_id", Type: field.TypeInt},
	}
	// GroupTagsTable holds the schema information for the "group_tags" table.
	GroupTagsTable = &schema.Table{
		Name:       "group_tags",
		Columns:    GroupTagsColumns,
		PrimaryKey: []*schema.Column{GroupTagsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "group_tags_tags_tag",
				Columns:    []*schema.Column{GroupTagsColumns[1]},
				RefColumns: []*schema.Column{TagsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "group_tags_groups_group",
				Columns:    []*schema.Column{GroupTagsColumns[2]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "grouptag_tag_id_group_id",
				Unique:  true,
				Columns: []*schema.Column{GroupTagsColumns[1], GroupTagsColumns[2]},
			},
		},
	}
	// RelationshipsColumns holds the columns for the "relationships" table.
	RelationshipsColumns = []*schema.Column{
		{Name: "weight", Type: field.TypeInt, Default: 1},
		{Name: "user_id", Type: field.TypeInt},
		{Name: "relative_id", Type: field.TypeInt},
		{Name: "info_id", Type: field.TypeInt, Nullable: true},
	}
	// RelationshipsTable holds the schema information for the "relationships" table.
	RelationshipsTable = &schema.Table{
		Name:       "relationships",
		Columns:    RelationshipsColumns,
		PrimaryKey: []*schema.Column{RelationshipsColumns[1], RelationshipsColumns[2]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "relationships_users_user",
				Columns:    []*schema.Column{RelationshipsColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "relationships_users_relative",
				Columns:    []*schema.Column{RelationshipsColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "relationships_relationship_infos_info",
				Columns:    []*schema.Column{RelationshipsColumns[3]},
				RefColumns: []*schema.Column{RelationshipInfosColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "relationship_weight",
				Unique:  false,
				Columns: []*schema.Column{RelationshipsColumns[0]},
			},
			{
				Name:    "relationship_info_id",
				Unique:  true,
				Columns: []*schema.Column{RelationshipsColumns[3]},
			},
		},
	}
	// RelationshipInfosColumns holds the columns for the "relationship_infos" table.
	RelationshipInfosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "text", Type: field.TypeString},
	}
	// RelationshipInfosTable holds the schema information for the "relationship_infos" table.
	RelationshipInfosTable = &schema.Table{
		Name:       "relationship_infos",
		Columns:    RelationshipInfosColumns,
		PrimaryKey: []*schema.Column{RelationshipInfosColumns[0]},
	}
	// RolesColumns holds the columns for the "roles" table.
	RolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
	}
	// RolesTable holds the schema information for the "roles" table.
	RolesTable = &schema.Table{
		Name:       "roles",
		Columns:    RolesColumns,
		PrimaryKey: []*schema.Column{RolesColumns[0]},
	}
	// RoleUsersColumns holds the columns for the "role_users" table.
	RoleUsersColumns = []*schema.Column{
		{Name: "created_at", Type: field.TypeTime},
		{Name: "role_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt},
	}
	// RoleUsersTable holds the schema information for the "role_users" table.
	RoleUsersTable = &schema.Table{
		Name:       "role_users",
		Columns:    RoleUsersColumns,
		PrimaryKey: []*schema.Column{RoleUsersColumns[2], RoleUsersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "role_users_roles_role",
				Columns:    []*schema.Column{RoleUsersColumns[1]},
				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "role_users_users_user",
				Columns:    []*schema.Column{RoleUsersColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// TagsColumns holds the columns for the "tags" table.
	TagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "value", Type: field.TypeString},
	}
	// TagsTable holds the schema information for the "tags" table.
	TagsTable = &schema.Table{
		Name:       "tags",
		Columns:    TagsColumns,
		PrimaryKey: []*schema.Column{TagsColumns[0]},
	}
	// TweetsColumns holds the columns for the "tweets" table.
	TweetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "text", Type: field.TypeString, Size: 2147483647},
	}
	// TweetsTable holds the schema information for the "tweets" table.
	TweetsTable = &schema.Table{
		Name:       "tweets",
		Columns:    TweetsColumns,
		PrimaryKey: []*schema.Column{TweetsColumns[0]},
	}
	// TweetLikesColumns holds the columns for the "tweet_likes" table.
	TweetLikesColumns = []*schema.Column{
		{Name: "liked_at", Type: field.TypeTime},
		{Name: "tweet_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt},
	}
	// TweetLikesTable holds the schema information for the "tweet_likes" table.
	TweetLikesTable = &schema.Table{
		Name:       "tweet_likes",
		Columns:    TweetLikesColumns,
		PrimaryKey: []*schema.Column{TweetLikesColumns[2], TweetLikesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tweet_likes_tweets_tweet",
				Columns:    []*schema.Column{TweetLikesColumns[1]},
				RefColumns: []*schema.Column{TweetsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "tweet_likes_users_user",
				Columns:    []*schema.Column{TweetLikesColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// TweetTagsColumns holds the columns for the "tweet_tags" table.
	TweetTagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "added_at", Type: field.TypeTime},
		{Name: "tag_id", Type: field.TypeInt},
		{Name: "tweet_id", Type: field.TypeInt},
	}
	// TweetTagsTable holds the schema information for the "tweet_tags" table.
	TweetTagsTable = &schema.Table{
		Name:       "tweet_tags",
		Columns:    TweetTagsColumns,
		PrimaryKey: []*schema.Column{TweetTagsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tweet_tags_tags_tag",
				Columns:    []*schema.Column{TweetTagsColumns[2]},
				RefColumns: []*schema.Column{TagsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "tweet_tags_tweets_tweet",
				Columns:    []*schema.Column{TweetTagsColumns[3]},
				RefColumns: []*schema.Column{TweetsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "tweettag_tag_id_tweet_id",
				Unique:  true,
				Columns: []*schema.Column{TweetTagsColumns[2], TweetTagsColumns[3]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Default: "Unknown"},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// UserGroupsColumns holds the columns for the "user_groups" table.
	UserGroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "joined_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeInt},
		{Name: "group_id", Type: field.TypeInt},
	}
	// UserGroupsTable holds the schema information for the "user_groups" table.
	UserGroupsTable = &schema.Table{
		Name:       "user_groups",
		Columns:    UserGroupsColumns,
		PrimaryKey: []*schema.Column{UserGroupsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_groups_users_user",
				Columns:    []*schema.Column{UserGroupsColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "user_groups_groups_group",
				Columns:    []*schema.Column{UserGroupsColumns[3]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "usergroup_user_id_group_id",
				Unique:  true,
				Columns: []*schema.Column{UserGroupsColumns[2], UserGroupsColumns[3]},
			},
		},
	}
	// UserTweetsColumns holds the columns for the "user_tweets" table.
	UserTweetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeInt},
		{Name: "tweet_id", Type: field.TypeInt},
	}
	// UserTweetsTable holds the schema information for the "user_tweets" table.
	UserTweetsTable = &schema.Table{
		Name:       "user_tweets",
		Columns:    UserTweetsColumns,
		PrimaryKey: []*schema.Column{UserTweetsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_tweets_users_user",
				Columns:    []*schema.Column{UserTweetsColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "user_tweets_tweets_tweet",
				Columns:    []*schema.Column{UserTweetsColumns[3]},
				RefColumns: []*schema.Column{TweetsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "usertweet_tweet_id",
				Unique:  true,
				Columns: []*schema.Column{UserTweetsColumns[3]},
			},
			{
				Name:    "usertweet_user_id_tweet_id",
				Unique:  true,
				Columns: []*schema.Column{UserTweetsColumns[2], UserTweetsColumns[3]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		FriendshipsTable,
		GroupsTable,
		GroupTagsTable,
		RelationshipsTable,
		RelationshipInfosTable,
		RolesTable,
		RoleUsersTable,
		TagsTable,
		TweetsTable,
		TweetLikesTable,
		TweetTagsTable,
		UsersTable,
		UserGroupsTable,
		UserTweetsTable,
	}
)

func init() {
	FriendshipsTable.ForeignKeys[0].RefTable = UsersTable
	FriendshipsTable.ForeignKeys[1].RefTable = UsersTable
	GroupTagsTable.ForeignKeys[0].RefTable = TagsTable
	GroupTagsTable.ForeignKeys[1].RefTable = GroupsTable
	RelationshipsTable.ForeignKeys[0].RefTable = UsersTable
	RelationshipsTable.ForeignKeys[1].RefTable = UsersTable
	RelationshipsTable.ForeignKeys[2].RefTable = RelationshipInfosTable
	RoleUsersTable.ForeignKeys[0].RefTable = RolesTable
	RoleUsersTable.ForeignKeys[1].RefTable = UsersTable
	TweetLikesTable.ForeignKeys[0].RefTable = TweetsTable
	TweetLikesTable.ForeignKeys[1].RefTable = UsersTable
	TweetTagsTable.ForeignKeys[0].RefTable = TagsTable
	TweetTagsTable.ForeignKeys[1].RefTable = TweetsTable
	UserGroupsTable.ForeignKeys[0].RefTable = UsersTable
	UserGroupsTable.ForeignKeys[1].RefTable = GroupsTable
	UserTweetsTable.ForeignKeys[0].RefTable = UsersTable
	UserTweetsTable.ForeignKeys[1].RefTable = TweetsTable
}
