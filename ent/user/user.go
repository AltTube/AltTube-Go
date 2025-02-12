// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// EdgeAccessTokens holds the string denoting the access_tokens edge name in mutations.
	EdgeAccessTokens = "access_tokens"
	// EdgeLikeVideos holds the string denoting the like_videos edge name in mutations.
	EdgeLikeVideos = "like_videos"
	// EdgeRefreshTokens holds the string denoting the refresh_tokens edge name in mutations.
	EdgeRefreshTokens = "refresh_tokens"
	// Table holds the table name of the user in the database.
	Table = "users"
	// AccessTokensTable is the table that holds the access_tokens relation/edge.
	AccessTokensTable = "access_tokens"
	// AccessTokensInverseTable is the table name for the AccessToken entity.
	// It exists in this package in order to avoid circular dependency with the "accesstoken" package.
	AccessTokensInverseTable = "access_tokens"
	// AccessTokensColumn is the table column denoting the access_tokens relation/edge.
	AccessTokensColumn = "user_id"
	// LikeVideosTable is the table that holds the like_videos relation/edge.
	LikeVideosTable = "like_videos"
	// LikeVideosInverseTable is the table name for the LikeVideo entity.
	// It exists in this package in order to avoid circular dependency with the "likevideo" package.
	LikeVideosInverseTable = "like_videos"
	// LikeVideosColumn is the table column denoting the like_videos relation/edge.
	LikeVideosColumn = "user_id"
	// RefreshTokensTable is the table that holds the refresh_tokens relation/edge.
	RefreshTokensTable = "refresh_tokens"
	// RefreshTokensInverseTable is the table name for the RefreshToken entity.
	// It exists in this package in order to avoid circular dependency with the "refreshtoken" package.
	RefreshTokensInverseTable = "refresh_tokens"
	// RefreshTokensColumn is the table column denoting the refresh_tokens relation/edge.
	RefreshTokensColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldEmail,
	FieldPassword,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByAccessTokensCount orders the results by access_tokens count.
func ByAccessTokensCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAccessTokensStep(), opts...)
	}
}

// ByAccessTokens orders the results by access_tokens terms.
func ByAccessTokens(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAccessTokensStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByLikeVideosCount orders the results by like_videos count.
func ByLikeVideosCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newLikeVideosStep(), opts...)
	}
}

// ByLikeVideos orders the results by like_videos terms.
func ByLikeVideos(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLikeVideosStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByRefreshTokensCount orders the results by refresh_tokens count.
func ByRefreshTokensCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRefreshTokensStep(), opts...)
	}
}

// ByRefreshTokens orders the results by refresh_tokens terms.
func ByRefreshTokens(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRefreshTokensStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newAccessTokensStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AccessTokensInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AccessTokensTable, AccessTokensColumn),
	)
}
func newLikeVideosStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LikeVideosInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, LikeVideosTable, LikeVideosColumn),
	)
}
func newRefreshTokensStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RefreshTokensInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, RefreshTokensTable, RefreshTokensColumn),
	)
}
