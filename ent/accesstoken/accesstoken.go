// Code generated by ent, DO NOT EDIT.

package accesstoken

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the accesstoken type in the database.
	Label = "access_token"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldToken holds the string denoting the token field in the database.
	FieldToken = "token"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldExpiry holds the string denoting the expiry field in the database.
	FieldExpiry = "expiry"
	// FieldRefreshTokenID holds the string denoting the refresh_token_id field in the database.
	FieldRefreshTokenID = "refresh_token_id"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeRefreshToken holds the string denoting the refresh_token edge name in mutations.
	EdgeRefreshToken = "refresh_token"
	// Table holds the table name of the accesstoken in the database.
	Table = "access_tokens"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "access_tokens"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
	// RefreshTokenTable is the table that holds the refresh_token relation/edge.
	RefreshTokenTable = "access_tokens"
	// RefreshTokenInverseTable is the table name for the RefreshToken entity.
	// It exists in this package in order to avoid circular dependency with the "refreshtoken" package.
	RefreshTokenInverseTable = "refresh_tokens"
	// RefreshTokenColumn is the table column denoting the refresh_token relation/edge.
	RefreshTokenColumn = "refresh_token_id"
)

// Columns holds all SQL columns for accesstoken fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldToken,
	FieldUserID,
	FieldExpiry,
	FieldRefreshTokenID,
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

// OrderOption defines the ordering options for the AccessToken queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByToken orders the results by the token field.
func ByToken(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldToken, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByExpiry orders the results by the expiry field.
func ByExpiry(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExpiry, opts...).ToFunc()
}

// ByRefreshTokenID orders the results by the refresh_token_id field.
func ByRefreshTokenID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRefreshTokenID, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByRefreshTokenField orders the results by refresh_token field.
func ByRefreshTokenField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRefreshTokenStep(), sql.OrderByField(field, opts...))
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
func newRefreshTokenStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RefreshTokenInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, RefreshTokenTable, RefreshTokenColumn),
	)
}
