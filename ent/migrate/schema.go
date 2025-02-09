// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AccessTokensColumns holds the columns for the "access_tokens" table.
	AccessTokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "token", Type: field.TypeString, Nullable: true},
		{Name: "expiry", Type: field.TypeTime, Nullable: true},
		{Name: "refresh_token_id", Type: field.TypeUint},
		{Name: "user_id", Type: field.TypeUUID, SchemaType: map[string]string{"postgres": "uuid"}},
	}
	// AccessTokensTable holds the schema information for the "access_tokens" table.
	AccessTokensTable = &schema.Table{
		Name:       "access_tokens",
		Columns:    AccessTokensColumns,
		PrimaryKey: []*schema.Column{AccessTokensColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "access_tokens_refresh_tokens_access_tokens",
				Columns:    []*schema.Column{AccessTokensColumns[5]},
				RefColumns: []*schema.Column{RefreshTokensColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "access_tokens_users_access_tokens",
				Columns:    []*schema.Column{AccessTokensColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// LikeVideosColumns holds the columns for the "like_videos" table.
	LikeVideosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeUUID, SchemaType: map[string]string{"postgres": "uuid"}},
		{Name: "video_id", Type: field.TypeString},
	}
	// LikeVideosTable holds the schema information for the "like_videos" table.
	LikeVideosTable = &schema.Table{
		Name:       "like_videos",
		Columns:    LikeVideosColumns,
		PrimaryKey: []*schema.Column{LikeVideosColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "like_videos_users_like_videos",
				Columns:    []*schema.Column{LikeVideosColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "like_videos_videos_like_videos",
				Columns:    []*schema.Column{LikeVideosColumns[4]},
				RefColumns: []*schema.Column{VideosColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// RefreshTokensColumns holds the columns for the "refresh_tokens" table.
	RefreshTokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "token", Type: field.TypeString, Nullable: true},
		{Name: "expiry", Type: field.TypeTime, Nullable: true},
		{Name: "user_agent", Type: field.TypeString, Nullable: true},
		{Name: "ip_address", Type: field.TypeString, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, SchemaType: map[string]string{"postgres": "uuid"}},
	}
	// RefreshTokensTable holds the schema information for the "refresh_tokens" table.
	RefreshTokensTable = &schema.Table{
		Name:       "refresh_tokens",
		Columns:    RefreshTokensColumns,
		PrimaryKey: []*schema.Column{RefreshTokensColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "refresh_tokens_users_refresh_tokens",
				Columns:    []*schema.Column{RefreshTokensColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true, SchemaType: map[string]string{"postgres": "uuid"}},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// VideosColumns holds the columns for the "videos" table.
	VideosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "title", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "upload_date", Type: field.TypeTime},
		{Name: "uploader", Type: field.TypeString},
		{Name: "uploader_url", Type: field.TypeString},
		{Name: "thumbnail_url", Type: field.TypeString},
	}
	// VideosTable holds the schema information for the "videos" table.
	VideosTable = &schema.Table{
		Name:       "videos",
		Columns:    VideosColumns,
		PrimaryKey: []*schema.Column{VideosColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccessTokensTable,
		LikeVideosTable,
		RefreshTokensTable,
		UsersTable,
		VideosTable,
	}
)

func init() {
	AccessTokensTable.ForeignKeys[0].RefTable = RefreshTokensTable
	AccessTokensTable.ForeignKeys[1].RefTable = UsersTable
	LikeVideosTable.ForeignKeys[0].RefTable = UsersTable
	LikeVideosTable.ForeignKeys[1].RefTable = VideosTable
	RefreshTokensTable.ForeignKeys[0].RefTable = UsersTable
}
