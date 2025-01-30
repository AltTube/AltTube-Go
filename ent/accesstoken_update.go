// Code generated by ent, DO NOT EDIT.

package ent

import (
	"AltTube-Go/ent/accesstoken"
	"AltTube-Go/ent/predicate"
	"AltTube-Go/ent/refreshtoken"
	"AltTube-Go/ent/user"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AccessTokenUpdate is the builder for updating AccessToken entities.
type AccessTokenUpdate struct {
	config
	hooks    []Hook
	mutation *AccessTokenMutation
}

// Where appends a list predicates to the AccessTokenUpdate builder.
func (atu *AccessTokenUpdate) Where(ps ...predicate.AccessToken) *AccessTokenUpdate {
	atu.mutation.Where(ps...)
	return atu
}

// SetCreatedAt sets the "created_at" field.
func (atu *AccessTokenUpdate) SetCreatedAt(t time.Time) *AccessTokenUpdate {
	atu.mutation.SetCreatedAt(t)
	return atu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (atu *AccessTokenUpdate) SetNillableCreatedAt(t *time.Time) *AccessTokenUpdate {
	if t != nil {
		atu.SetCreatedAt(*t)
	}
	return atu
}

// ClearCreatedAt clears the value of the "created_at" field.
func (atu *AccessTokenUpdate) ClearCreatedAt() *AccessTokenUpdate {
	atu.mutation.ClearCreatedAt()
	return atu
}

// SetUpdatedAt sets the "updated_at" field.
func (atu *AccessTokenUpdate) SetUpdatedAt(t time.Time) *AccessTokenUpdate {
	atu.mutation.SetUpdatedAt(t)
	return atu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (atu *AccessTokenUpdate) SetNillableUpdatedAt(t *time.Time) *AccessTokenUpdate {
	if t != nil {
		atu.SetUpdatedAt(*t)
	}
	return atu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (atu *AccessTokenUpdate) ClearUpdatedAt() *AccessTokenUpdate {
	atu.mutation.ClearUpdatedAt()
	return atu
}

// SetDeletedAt sets the "deleted_at" field.
func (atu *AccessTokenUpdate) SetDeletedAt(t time.Time) *AccessTokenUpdate {
	atu.mutation.SetDeletedAt(t)
	return atu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (atu *AccessTokenUpdate) SetNillableDeletedAt(t *time.Time) *AccessTokenUpdate {
	if t != nil {
		atu.SetDeletedAt(*t)
	}
	return atu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (atu *AccessTokenUpdate) ClearDeletedAt() *AccessTokenUpdate {
	atu.mutation.ClearDeletedAt()
	return atu
}

// SetToken sets the "token" field.
func (atu *AccessTokenUpdate) SetToken(s string) *AccessTokenUpdate {
	atu.mutation.SetToken(s)
	return atu
}

// SetNillableToken sets the "token" field if the given value is not nil.
func (atu *AccessTokenUpdate) SetNillableToken(s *string) *AccessTokenUpdate {
	if s != nil {
		atu.SetToken(*s)
	}
	return atu
}

// ClearToken clears the value of the "token" field.
func (atu *AccessTokenUpdate) ClearToken() *AccessTokenUpdate {
	atu.mutation.ClearToken()
	return atu
}

// SetUserID sets the "user_id" field.
func (atu *AccessTokenUpdate) SetUserID(s string) *AccessTokenUpdate {
	atu.mutation.SetUserID(s)
	return atu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (atu *AccessTokenUpdate) SetNillableUserID(s *string) *AccessTokenUpdate {
	if s != nil {
		atu.SetUserID(*s)
	}
	return atu
}

// ClearUserID clears the value of the "user_id" field.
func (atu *AccessTokenUpdate) ClearUserID() *AccessTokenUpdate {
	atu.mutation.ClearUserID()
	return atu
}

// SetExpiry sets the "expiry" field.
func (atu *AccessTokenUpdate) SetExpiry(t time.Time) *AccessTokenUpdate {
	atu.mutation.SetExpiry(t)
	return atu
}

// SetNillableExpiry sets the "expiry" field if the given value is not nil.
func (atu *AccessTokenUpdate) SetNillableExpiry(t *time.Time) *AccessTokenUpdate {
	if t != nil {
		atu.SetExpiry(*t)
	}
	return atu
}

// ClearExpiry clears the value of the "expiry" field.
func (atu *AccessTokenUpdate) ClearExpiry() *AccessTokenUpdate {
	atu.mutation.ClearExpiry()
	return atu
}

// SetRefreshTokenID sets the "refresh_token_id" field.
func (atu *AccessTokenUpdate) SetRefreshTokenID(i int64) *AccessTokenUpdate {
	atu.mutation.SetRefreshTokenID(i)
	return atu
}

// SetNillableRefreshTokenID sets the "refresh_token_id" field if the given value is not nil.
func (atu *AccessTokenUpdate) SetNillableRefreshTokenID(i *int64) *AccessTokenUpdate {
	if i != nil {
		atu.SetRefreshTokenID(*i)
	}
	return atu
}

// ClearRefreshTokenID clears the value of the "refresh_token_id" field.
func (atu *AccessTokenUpdate) ClearRefreshTokenID() *AccessTokenUpdate {
	atu.mutation.ClearRefreshTokenID()
	return atu
}

// SetUser sets the "user" edge to the User entity.
func (atu *AccessTokenUpdate) SetUser(u *User) *AccessTokenUpdate {
	return atu.SetUserID(u.ID)
}

// SetRefreshToken sets the "refresh_token" edge to the RefreshToken entity.
func (atu *AccessTokenUpdate) SetRefreshToken(r *RefreshToken) *AccessTokenUpdate {
	return atu.SetRefreshTokenID(r.ID)
}

// Mutation returns the AccessTokenMutation object of the builder.
func (atu *AccessTokenUpdate) Mutation() *AccessTokenMutation {
	return atu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (atu *AccessTokenUpdate) ClearUser() *AccessTokenUpdate {
	atu.mutation.ClearUser()
	return atu
}

// ClearRefreshToken clears the "refresh_token" edge to the RefreshToken entity.
func (atu *AccessTokenUpdate) ClearRefreshToken() *AccessTokenUpdate {
	atu.mutation.ClearRefreshToken()
	return atu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (atu *AccessTokenUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, atu.sqlSave, atu.mutation, atu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (atu *AccessTokenUpdate) SaveX(ctx context.Context) int {
	affected, err := atu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (atu *AccessTokenUpdate) Exec(ctx context.Context) error {
	_, err := atu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atu *AccessTokenUpdate) ExecX(ctx context.Context) {
	if err := atu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (atu *AccessTokenUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(accesstoken.Table, accesstoken.Columns, sqlgraph.NewFieldSpec(accesstoken.FieldID, field.TypeInt64))
	if ps := atu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := atu.mutation.CreatedAt(); ok {
		_spec.SetField(accesstoken.FieldCreatedAt, field.TypeTime, value)
	}
	if atu.mutation.CreatedAtCleared() {
		_spec.ClearField(accesstoken.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := atu.mutation.UpdatedAt(); ok {
		_spec.SetField(accesstoken.FieldUpdatedAt, field.TypeTime, value)
	}
	if atu.mutation.UpdatedAtCleared() {
		_spec.ClearField(accesstoken.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := atu.mutation.DeletedAt(); ok {
		_spec.SetField(accesstoken.FieldDeletedAt, field.TypeTime, value)
	}
	if atu.mutation.DeletedAtCleared() {
		_spec.ClearField(accesstoken.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := atu.mutation.Token(); ok {
		_spec.SetField(accesstoken.FieldToken, field.TypeString, value)
	}
	if atu.mutation.TokenCleared() {
		_spec.ClearField(accesstoken.FieldToken, field.TypeString)
	}
	if value, ok := atu.mutation.Expiry(); ok {
		_spec.SetField(accesstoken.FieldExpiry, field.TypeTime, value)
	}
	if atu.mutation.ExpiryCleared() {
		_spec.ClearField(accesstoken.FieldExpiry, field.TypeTime)
	}
	if atu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   accesstoken.UserTable,
			Columns: []string{accesstoken.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   accesstoken.UserTable,
			Columns: []string{accesstoken.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if atu.mutation.RefreshTokenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   accesstoken.RefreshTokenTable,
			Columns: []string{accesstoken.RefreshTokenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(refreshtoken.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atu.mutation.RefreshTokenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   accesstoken.RefreshTokenTable,
			Columns: []string{accesstoken.RefreshTokenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(refreshtoken.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, atu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{accesstoken.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	atu.mutation.done = true
	return n, nil
}

// AccessTokenUpdateOne is the builder for updating a single AccessToken entity.
type AccessTokenUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AccessTokenMutation
}

// SetCreatedAt sets the "created_at" field.
func (atuo *AccessTokenUpdateOne) SetCreatedAt(t time.Time) *AccessTokenUpdateOne {
	atuo.mutation.SetCreatedAt(t)
	return atuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (atuo *AccessTokenUpdateOne) SetNillableCreatedAt(t *time.Time) *AccessTokenUpdateOne {
	if t != nil {
		atuo.SetCreatedAt(*t)
	}
	return atuo
}

// ClearCreatedAt clears the value of the "created_at" field.
func (atuo *AccessTokenUpdateOne) ClearCreatedAt() *AccessTokenUpdateOne {
	atuo.mutation.ClearCreatedAt()
	return atuo
}

// SetUpdatedAt sets the "updated_at" field.
func (atuo *AccessTokenUpdateOne) SetUpdatedAt(t time.Time) *AccessTokenUpdateOne {
	atuo.mutation.SetUpdatedAt(t)
	return atuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (atuo *AccessTokenUpdateOne) SetNillableUpdatedAt(t *time.Time) *AccessTokenUpdateOne {
	if t != nil {
		atuo.SetUpdatedAt(*t)
	}
	return atuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (atuo *AccessTokenUpdateOne) ClearUpdatedAt() *AccessTokenUpdateOne {
	atuo.mutation.ClearUpdatedAt()
	return atuo
}

// SetDeletedAt sets the "deleted_at" field.
func (atuo *AccessTokenUpdateOne) SetDeletedAt(t time.Time) *AccessTokenUpdateOne {
	atuo.mutation.SetDeletedAt(t)
	return atuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (atuo *AccessTokenUpdateOne) SetNillableDeletedAt(t *time.Time) *AccessTokenUpdateOne {
	if t != nil {
		atuo.SetDeletedAt(*t)
	}
	return atuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (atuo *AccessTokenUpdateOne) ClearDeletedAt() *AccessTokenUpdateOne {
	atuo.mutation.ClearDeletedAt()
	return atuo
}

// SetToken sets the "token" field.
func (atuo *AccessTokenUpdateOne) SetToken(s string) *AccessTokenUpdateOne {
	atuo.mutation.SetToken(s)
	return atuo
}

// SetNillableToken sets the "token" field if the given value is not nil.
func (atuo *AccessTokenUpdateOne) SetNillableToken(s *string) *AccessTokenUpdateOne {
	if s != nil {
		atuo.SetToken(*s)
	}
	return atuo
}

// ClearToken clears the value of the "token" field.
func (atuo *AccessTokenUpdateOne) ClearToken() *AccessTokenUpdateOne {
	atuo.mutation.ClearToken()
	return atuo
}

// SetUserID sets the "user_id" field.
func (atuo *AccessTokenUpdateOne) SetUserID(s string) *AccessTokenUpdateOne {
	atuo.mutation.SetUserID(s)
	return atuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (atuo *AccessTokenUpdateOne) SetNillableUserID(s *string) *AccessTokenUpdateOne {
	if s != nil {
		atuo.SetUserID(*s)
	}
	return atuo
}

// ClearUserID clears the value of the "user_id" field.
func (atuo *AccessTokenUpdateOne) ClearUserID() *AccessTokenUpdateOne {
	atuo.mutation.ClearUserID()
	return atuo
}

// SetExpiry sets the "expiry" field.
func (atuo *AccessTokenUpdateOne) SetExpiry(t time.Time) *AccessTokenUpdateOne {
	atuo.mutation.SetExpiry(t)
	return atuo
}

// SetNillableExpiry sets the "expiry" field if the given value is not nil.
func (atuo *AccessTokenUpdateOne) SetNillableExpiry(t *time.Time) *AccessTokenUpdateOne {
	if t != nil {
		atuo.SetExpiry(*t)
	}
	return atuo
}

// ClearExpiry clears the value of the "expiry" field.
func (atuo *AccessTokenUpdateOne) ClearExpiry() *AccessTokenUpdateOne {
	atuo.mutation.ClearExpiry()
	return atuo
}

// SetRefreshTokenID sets the "refresh_token_id" field.
func (atuo *AccessTokenUpdateOne) SetRefreshTokenID(i int64) *AccessTokenUpdateOne {
	atuo.mutation.SetRefreshTokenID(i)
	return atuo
}

// SetNillableRefreshTokenID sets the "refresh_token_id" field if the given value is not nil.
func (atuo *AccessTokenUpdateOne) SetNillableRefreshTokenID(i *int64) *AccessTokenUpdateOne {
	if i != nil {
		atuo.SetRefreshTokenID(*i)
	}
	return atuo
}

// ClearRefreshTokenID clears the value of the "refresh_token_id" field.
func (atuo *AccessTokenUpdateOne) ClearRefreshTokenID() *AccessTokenUpdateOne {
	atuo.mutation.ClearRefreshTokenID()
	return atuo
}

// SetUser sets the "user" edge to the User entity.
func (atuo *AccessTokenUpdateOne) SetUser(u *User) *AccessTokenUpdateOne {
	return atuo.SetUserID(u.ID)
}

// SetRefreshToken sets the "refresh_token" edge to the RefreshToken entity.
func (atuo *AccessTokenUpdateOne) SetRefreshToken(r *RefreshToken) *AccessTokenUpdateOne {
	return atuo.SetRefreshTokenID(r.ID)
}

// Mutation returns the AccessTokenMutation object of the builder.
func (atuo *AccessTokenUpdateOne) Mutation() *AccessTokenMutation {
	return atuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (atuo *AccessTokenUpdateOne) ClearUser() *AccessTokenUpdateOne {
	atuo.mutation.ClearUser()
	return atuo
}

// ClearRefreshToken clears the "refresh_token" edge to the RefreshToken entity.
func (atuo *AccessTokenUpdateOne) ClearRefreshToken() *AccessTokenUpdateOne {
	atuo.mutation.ClearRefreshToken()
	return atuo
}

// Where appends a list predicates to the AccessTokenUpdate builder.
func (atuo *AccessTokenUpdateOne) Where(ps ...predicate.AccessToken) *AccessTokenUpdateOne {
	atuo.mutation.Where(ps...)
	return atuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (atuo *AccessTokenUpdateOne) Select(field string, fields ...string) *AccessTokenUpdateOne {
	atuo.fields = append([]string{field}, fields...)
	return atuo
}

// Save executes the query and returns the updated AccessToken entity.
func (atuo *AccessTokenUpdateOne) Save(ctx context.Context) (*AccessToken, error) {
	return withHooks(ctx, atuo.sqlSave, atuo.mutation, atuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (atuo *AccessTokenUpdateOne) SaveX(ctx context.Context) *AccessToken {
	node, err := atuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (atuo *AccessTokenUpdateOne) Exec(ctx context.Context) error {
	_, err := atuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atuo *AccessTokenUpdateOne) ExecX(ctx context.Context) {
	if err := atuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (atuo *AccessTokenUpdateOne) sqlSave(ctx context.Context) (_node *AccessToken, err error) {
	_spec := sqlgraph.NewUpdateSpec(accesstoken.Table, accesstoken.Columns, sqlgraph.NewFieldSpec(accesstoken.FieldID, field.TypeInt64))
	id, ok := atuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AccessToken.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := atuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, accesstoken.FieldID)
		for _, f := range fields {
			if !accesstoken.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != accesstoken.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := atuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := atuo.mutation.CreatedAt(); ok {
		_spec.SetField(accesstoken.FieldCreatedAt, field.TypeTime, value)
	}
	if atuo.mutation.CreatedAtCleared() {
		_spec.ClearField(accesstoken.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := atuo.mutation.UpdatedAt(); ok {
		_spec.SetField(accesstoken.FieldUpdatedAt, field.TypeTime, value)
	}
	if atuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(accesstoken.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := atuo.mutation.DeletedAt(); ok {
		_spec.SetField(accesstoken.FieldDeletedAt, field.TypeTime, value)
	}
	if atuo.mutation.DeletedAtCleared() {
		_spec.ClearField(accesstoken.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := atuo.mutation.Token(); ok {
		_spec.SetField(accesstoken.FieldToken, field.TypeString, value)
	}
	if atuo.mutation.TokenCleared() {
		_spec.ClearField(accesstoken.FieldToken, field.TypeString)
	}
	if value, ok := atuo.mutation.Expiry(); ok {
		_spec.SetField(accesstoken.FieldExpiry, field.TypeTime, value)
	}
	if atuo.mutation.ExpiryCleared() {
		_spec.ClearField(accesstoken.FieldExpiry, field.TypeTime)
	}
	if atuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   accesstoken.UserTable,
			Columns: []string{accesstoken.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   accesstoken.UserTable,
			Columns: []string{accesstoken.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if atuo.mutation.RefreshTokenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   accesstoken.RefreshTokenTable,
			Columns: []string{accesstoken.RefreshTokenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(refreshtoken.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atuo.mutation.RefreshTokenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   accesstoken.RefreshTokenTable,
			Columns: []string{accesstoken.RefreshTokenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(refreshtoken.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &AccessToken{config: atuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, atuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{accesstoken.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	atuo.mutation.done = true
	return _node, nil
}
