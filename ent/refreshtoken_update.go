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
	"github.com/google/uuid"
)

// RefreshTokenUpdate is the builder for updating RefreshToken entities.
type RefreshTokenUpdate struct {
	config
	hooks    []Hook
	mutation *RefreshTokenMutation
}

// Where appends a list predicates to the RefreshTokenUpdate builder.
func (rtu *RefreshTokenUpdate) Where(ps ...predicate.RefreshToken) *RefreshTokenUpdate {
	rtu.mutation.Where(ps...)
	return rtu
}

// SetUpdateTime sets the "update_time" field.
func (rtu *RefreshTokenUpdate) SetUpdateTime(t time.Time) *RefreshTokenUpdate {
	rtu.mutation.SetUpdateTime(t)
	return rtu
}

// SetToken sets the "token" field.
func (rtu *RefreshTokenUpdate) SetToken(s string) *RefreshTokenUpdate {
	rtu.mutation.SetToken(s)
	return rtu
}

// SetNillableToken sets the "token" field if the given value is not nil.
func (rtu *RefreshTokenUpdate) SetNillableToken(s *string) *RefreshTokenUpdate {
	if s != nil {
		rtu.SetToken(*s)
	}
	return rtu
}

// ClearToken clears the value of the "token" field.
func (rtu *RefreshTokenUpdate) ClearToken() *RefreshTokenUpdate {
	rtu.mutation.ClearToken()
	return rtu
}

// SetExpiry sets the "expiry" field.
func (rtu *RefreshTokenUpdate) SetExpiry(t time.Time) *RefreshTokenUpdate {
	rtu.mutation.SetExpiry(t)
	return rtu
}

// SetNillableExpiry sets the "expiry" field if the given value is not nil.
func (rtu *RefreshTokenUpdate) SetNillableExpiry(t *time.Time) *RefreshTokenUpdate {
	if t != nil {
		rtu.SetExpiry(*t)
	}
	return rtu
}

// ClearExpiry clears the value of the "expiry" field.
func (rtu *RefreshTokenUpdate) ClearExpiry() *RefreshTokenUpdate {
	rtu.mutation.ClearExpiry()
	return rtu
}

// SetUserAgent sets the "user_agent" field.
func (rtu *RefreshTokenUpdate) SetUserAgent(s string) *RefreshTokenUpdate {
	rtu.mutation.SetUserAgent(s)
	return rtu
}

// SetNillableUserAgent sets the "user_agent" field if the given value is not nil.
func (rtu *RefreshTokenUpdate) SetNillableUserAgent(s *string) *RefreshTokenUpdate {
	if s != nil {
		rtu.SetUserAgent(*s)
	}
	return rtu
}

// ClearUserAgent clears the value of the "user_agent" field.
func (rtu *RefreshTokenUpdate) ClearUserAgent() *RefreshTokenUpdate {
	rtu.mutation.ClearUserAgent()
	return rtu
}

// SetIPAddress sets the "ip_address" field.
func (rtu *RefreshTokenUpdate) SetIPAddress(s string) *RefreshTokenUpdate {
	rtu.mutation.SetIPAddress(s)
	return rtu
}

// SetNillableIPAddress sets the "ip_address" field if the given value is not nil.
func (rtu *RefreshTokenUpdate) SetNillableIPAddress(s *string) *RefreshTokenUpdate {
	if s != nil {
		rtu.SetIPAddress(*s)
	}
	return rtu
}

// ClearIPAddress clears the value of the "ip_address" field.
func (rtu *RefreshTokenUpdate) ClearIPAddress() *RefreshTokenUpdate {
	rtu.mutation.ClearIPAddress()
	return rtu
}

// SetUserID sets the "user_id" field.
func (rtu *RefreshTokenUpdate) SetUserID(u uuid.UUID) *RefreshTokenUpdate {
	rtu.mutation.SetUserID(u)
	return rtu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (rtu *RefreshTokenUpdate) SetNillableUserID(u *uuid.UUID) *RefreshTokenUpdate {
	if u != nil {
		rtu.SetUserID(*u)
	}
	return rtu
}

// SetUser sets the "user" edge to the User entity.
func (rtu *RefreshTokenUpdate) SetUser(u *User) *RefreshTokenUpdate {
	return rtu.SetUserID(u.ID)
}

// AddAccessTokenIDs adds the "access_tokens" edge to the AccessToken entity by IDs.
func (rtu *RefreshTokenUpdate) AddAccessTokenIDs(ids ...uint) *RefreshTokenUpdate {
	rtu.mutation.AddAccessTokenIDs(ids...)
	return rtu
}

// AddAccessTokens adds the "access_tokens" edges to the AccessToken entity.
func (rtu *RefreshTokenUpdate) AddAccessTokens(a ...*AccessToken) *RefreshTokenUpdate {
	ids := make([]uint, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return rtu.AddAccessTokenIDs(ids...)
}

// Mutation returns the RefreshTokenMutation object of the builder.
func (rtu *RefreshTokenUpdate) Mutation() *RefreshTokenMutation {
	return rtu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (rtu *RefreshTokenUpdate) ClearUser() *RefreshTokenUpdate {
	rtu.mutation.ClearUser()
	return rtu
}

// ClearAccessTokens clears all "access_tokens" edges to the AccessToken entity.
func (rtu *RefreshTokenUpdate) ClearAccessTokens() *RefreshTokenUpdate {
	rtu.mutation.ClearAccessTokens()
	return rtu
}

// RemoveAccessTokenIDs removes the "access_tokens" edge to AccessToken entities by IDs.
func (rtu *RefreshTokenUpdate) RemoveAccessTokenIDs(ids ...uint) *RefreshTokenUpdate {
	rtu.mutation.RemoveAccessTokenIDs(ids...)
	return rtu
}

// RemoveAccessTokens removes "access_tokens" edges to AccessToken entities.
func (rtu *RefreshTokenUpdate) RemoveAccessTokens(a ...*AccessToken) *RefreshTokenUpdate {
	ids := make([]uint, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return rtu.RemoveAccessTokenIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rtu *RefreshTokenUpdate) Save(ctx context.Context) (int, error) {
	rtu.defaults()
	return withHooks(ctx, rtu.sqlSave, rtu.mutation, rtu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rtu *RefreshTokenUpdate) SaveX(ctx context.Context) int {
	affected, err := rtu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rtu *RefreshTokenUpdate) Exec(ctx context.Context) error {
	_, err := rtu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rtu *RefreshTokenUpdate) ExecX(ctx context.Context) {
	if err := rtu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rtu *RefreshTokenUpdate) defaults() {
	if _, ok := rtu.mutation.UpdateTime(); !ok {
		v := refreshtoken.UpdateDefaultUpdateTime()
		rtu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rtu *RefreshTokenUpdate) check() error {
	if rtu.mutation.UserCleared() && len(rtu.mutation.UserIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "RefreshToken.user"`)
	}
	return nil
}

func (rtu *RefreshTokenUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := rtu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(refreshtoken.Table, refreshtoken.Columns, sqlgraph.NewFieldSpec(refreshtoken.FieldID, field.TypeUint))
	if ps := rtu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rtu.mutation.UpdateTime(); ok {
		_spec.SetField(refreshtoken.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := rtu.mutation.Token(); ok {
		_spec.SetField(refreshtoken.FieldToken, field.TypeString, value)
	}
	if rtu.mutation.TokenCleared() {
		_spec.ClearField(refreshtoken.FieldToken, field.TypeString)
	}
	if value, ok := rtu.mutation.Expiry(); ok {
		_spec.SetField(refreshtoken.FieldExpiry, field.TypeTime, value)
	}
	if rtu.mutation.ExpiryCleared() {
		_spec.ClearField(refreshtoken.FieldExpiry, field.TypeTime)
	}
	if value, ok := rtu.mutation.UserAgent(); ok {
		_spec.SetField(refreshtoken.FieldUserAgent, field.TypeString, value)
	}
	if rtu.mutation.UserAgentCleared() {
		_spec.ClearField(refreshtoken.FieldUserAgent, field.TypeString)
	}
	if value, ok := rtu.mutation.IPAddress(); ok {
		_spec.SetField(refreshtoken.FieldIPAddress, field.TypeString, value)
	}
	if rtu.mutation.IPAddressCleared() {
		_spec.ClearField(refreshtoken.FieldIPAddress, field.TypeString)
	}
	if rtu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   refreshtoken.UserTable,
			Columns: []string{refreshtoken.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rtu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   refreshtoken.UserTable,
			Columns: []string{refreshtoken.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rtu.mutation.AccessTokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   refreshtoken.AccessTokensTable,
			Columns: []string{refreshtoken.AccessTokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(accesstoken.FieldID, field.TypeUint),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rtu.mutation.RemovedAccessTokensIDs(); len(nodes) > 0 && !rtu.mutation.AccessTokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   refreshtoken.AccessTokensTable,
			Columns: []string{refreshtoken.AccessTokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(accesstoken.FieldID, field.TypeUint),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rtu.mutation.AccessTokensIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   refreshtoken.AccessTokensTable,
			Columns: []string{refreshtoken.AccessTokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(accesstoken.FieldID, field.TypeUint),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, rtu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{refreshtoken.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	rtu.mutation.done = true
	return n, nil
}

// RefreshTokenUpdateOne is the builder for updating a single RefreshToken entity.
type RefreshTokenUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RefreshTokenMutation
}

// SetUpdateTime sets the "update_time" field.
func (rtuo *RefreshTokenUpdateOne) SetUpdateTime(t time.Time) *RefreshTokenUpdateOne {
	rtuo.mutation.SetUpdateTime(t)
	return rtuo
}

// SetToken sets the "token" field.
func (rtuo *RefreshTokenUpdateOne) SetToken(s string) *RefreshTokenUpdateOne {
	rtuo.mutation.SetToken(s)
	return rtuo
}

// SetNillableToken sets the "token" field if the given value is not nil.
func (rtuo *RefreshTokenUpdateOne) SetNillableToken(s *string) *RefreshTokenUpdateOne {
	if s != nil {
		rtuo.SetToken(*s)
	}
	return rtuo
}

// ClearToken clears the value of the "token" field.
func (rtuo *RefreshTokenUpdateOne) ClearToken() *RefreshTokenUpdateOne {
	rtuo.mutation.ClearToken()
	return rtuo
}

// SetExpiry sets the "expiry" field.
func (rtuo *RefreshTokenUpdateOne) SetExpiry(t time.Time) *RefreshTokenUpdateOne {
	rtuo.mutation.SetExpiry(t)
	return rtuo
}

// SetNillableExpiry sets the "expiry" field if the given value is not nil.
func (rtuo *RefreshTokenUpdateOne) SetNillableExpiry(t *time.Time) *RefreshTokenUpdateOne {
	if t != nil {
		rtuo.SetExpiry(*t)
	}
	return rtuo
}

// ClearExpiry clears the value of the "expiry" field.
func (rtuo *RefreshTokenUpdateOne) ClearExpiry() *RefreshTokenUpdateOne {
	rtuo.mutation.ClearExpiry()
	return rtuo
}

// SetUserAgent sets the "user_agent" field.
func (rtuo *RefreshTokenUpdateOne) SetUserAgent(s string) *RefreshTokenUpdateOne {
	rtuo.mutation.SetUserAgent(s)
	return rtuo
}

// SetNillableUserAgent sets the "user_agent" field if the given value is not nil.
func (rtuo *RefreshTokenUpdateOne) SetNillableUserAgent(s *string) *RefreshTokenUpdateOne {
	if s != nil {
		rtuo.SetUserAgent(*s)
	}
	return rtuo
}

// ClearUserAgent clears the value of the "user_agent" field.
func (rtuo *RefreshTokenUpdateOne) ClearUserAgent() *RefreshTokenUpdateOne {
	rtuo.mutation.ClearUserAgent()
	return rtuo
}

// SetIPAddress sets the "ip_address" field.
func (rtuo *RefreshTokenUpdateOne) SetIPAddress(s string) *RefreshTokenUpdateOne {
	rtuo.mutation.SetIPAddress(s)
	return rtuo
}

// SetNillableIPAddress sets the "ip_address" field if the given value is not nil.
func (rtuo *RefreshTokenUpdateOne) SetNillableIPAddress(s *string) *RefreshTokenUpdateOne {
	if s != nil {
		rtuo.SetIPAddress(*s)
	}
	return rtuo
}

// ClearIPAddress clears the value of the "ip_address" field.
func (rtuo *RefreshTokenUpdateOne) ClearIPAddress() *RefreshTokenUpdateOne {
	rtuo.mutation.ClearIPAddress()
	return rtuo
}

// SetUserID sets the "user_id" field.
func (rtuo *RefreshTokenUpdateOne) SetUserID(u uuid.UUID) *RefreshTokenUpdateOne {
	rtuo.mutation.SetUserID(u)
	return rtuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (rtuo *RefreshTokenUpdateOne) SetNillableUserID(u *uuid.UUID) *RefreshTokenUpdateOne {
	if u != nil {
		rtuo.SetUserID(*u)
	}
	return rtuo
}

// SetUser sets the "user" edge to the User entity.
func (rtuo *RefreshTokenUpdateOne) SetUser(u *User) *RefreshTokenUpdateOne {
	return rtuo.SetUserID(u.ID)
}

// AddAccessTokenIDs adds the "access_tokens" edge to the AccessToken entity by IDs.
func (rtuo *RefreshTokenUpdateOne) AddAccessTokenIDs(ids ...uint) *RefreshTokenUpdateOne {
	rtuo.mutation.AddAccessTokenIDs(ids...)
	return rtuo
}

// AddAccessTokens adds the "access_tokens" edges to the AccessToken entity.
func (rtuo *RefreshTokenUpdateOne) AddAccessTokens(a ...*AccessToken) *RefreshTokenUpdateOne {
	ids := make([]uint, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return rtuo.AddAccessTokenIDs(ids...)
}

// Mutation returns the RefreshTokenMutation object of the builder.
func (rtuo *RefreshTokenUpdateOne) Mutation() *RefreshTokenMutation {
	return rtuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (rtuo *RefreshTokenUpdateOne) ClearUser() *RefreshTokenUpdateOne {
	rtuo.mutation.ClearUser()
	return rtuo
}

// ClearAccessTokens clears all "access_tokens" edges to the AccessToken entity.
func (rtuo *RefreshTokenUpdateOne) ClearAccessTokens() *RefreshTokenUpdateOne {
	rtuo.mutation.ClearAccessTokens()
	return rtuo
}

// RemoveAccessTokenIDs removes the "access_tokens" edge to AccessToken entities by IDs.
func (rtuo *RefreshTokenUpdateOne) RemoveAccessTokenIDs(ids ...uint) *RefreshTokenUpdateOne {
	rtuo.mutation.RemoveAccessTokenIDs(ids...)
	return rtuo
}

// RemoveAccessTokens removes "access_tokens" edges to AccessToken entities.
func (rtuo *RefreshTokenUpdateOne) RemoveAccessTokens(a ...*AccessToken) *RefreshTokenUpdateOne {
	ids := make([]uint, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return rtuo.RemoveAccessTokenIDs(ids...)
}

// Where appends a list predicates to the RefreshTokenUpdate builder.
func (rtuo *RefreshTokenUpdateOne) Where(ps ...predicate.RefreshToken) *RefreshTokenUpdateOne {
	rtuo.mutation.Where(ps...)
	return rtuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rtuo *RefreshTokenUpdateOne) Select(field string, fields ...string) *RefreshTokenUpdateOne {
	rtuo.fields = append([]string{field}, fields...)
	return rtuo
}

// Save executes the query and returns the updated RefreshToken entity.
func (rtuo *RefreshTokenUpdateOne) Save(ctx context.Context) (*RefreshToken, error) {
	rtuo.defaults()
	return withHooks(ctx, rtuo.sqlSave, rtuo.mutation, rtuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rtuo *RefreshTokenUpdateOne) SaveX(ctx context.Context) *RefreshToken {
	node, err := rtuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rtuo *RefreshTokenUpdateOne) Exec(ctx context.Context) error {
	_, err := rtuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rtuo *RefreshTokenUpdateOne) ExecX(ctx context.Context) {
	if err := rtuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rtuo *RefreshTokenUpdateOne) defaults() {
	if _, ok := rtuo.mutation.UpdateTime(); !ok {
		v := refreshtoken.UpdateDefaultUpdateTime()
		rtuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rtuo *RefreshTokenUpdateOne) check() error {
	if rtuo.mutation.UserCleared() && len(rtuo.mutation.UserIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "RefreshToken.user"`)
	}
	return nil
}

func (rtuo *RefreshTokenUpdateOne) sqlSave(ctx context.Context) (_node *RefreshToken, err error) {
	if err := rtuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(refreshtoken.Table, refreshtoken.Columns, sqlgraph.NewFieldSpec(refreshtoken.FieldID, field.TypeUint))
	id, ok := rtuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "RefreshToken.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rtuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, refreshtoken.FieldID)
		for _, f := range fields {
			if !refreshtoken.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != refreshtoken.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rtuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rtuo.mutation.UpdateTime(); ok {
		_spec.SetField(refreshtoken.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := rtuo.mutation.Token(); ok {
		_spec.SetField(refreshtoken.FieldToken, field.TypeString, value)
	}
	if rtuo.mutation.TokenCleared() {
		_spec.ClearField(refreshtoken.FieldToken, field.TypeString)
	}
	if value, ok := rtuo.mutation.Expiry(); ok {
		_spec.SetField(refreshtoken.FieldExpiry, field.TypeTime, value)
	}
	if rtuo.mutation.ExpiryCleared() {
		_spec.ClearField(refreshtoken.FieldExpiry, field.TypeTime)
	}
	if value, ok := rtuo.mutation.UserAgent(); ok {
		_spec.SetField(refreshtoken.FieldUserAgent, field.TypeString, value)
	}
	if rtuo.mutation.UserAgentCleared() {
		_spec.ClearField(refreshtoken.FieldUserAgent, field.TypeString)
	}
	if value, ok := rtuo.mutation.IPAddress(); ok {
		_spec.SetField(refreshtoken.FieldIPAddress, field.TypeString, value)
	}
	if rtuo.mutation.IPAddressCleared() {
		_spec.ClearField(refreshtoken.FieldIPAddress, field.TypeString)
	}
	if rtuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   refreshtoken.UserTable,
			Columns: []string{refreshtoken.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rtuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   refreshtoken.UserTable,
			Columns: []string{refreshtoken.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rtuo.mutation.AccessTokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   refreshtoken.AccessTokensTable,
			Columns: []string{refreshtoken.AccessTokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(accesstoken.FieldID, field.TypeUint),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rtuo.mutation.RemovedAccessTokensIDs(); len(nodes) > 0 && !rtuo.mutation.AccessTokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   refreshtoken.AccessTokensTable,
			Columns: []string{refreshtoken.AccessTokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(accesstoken.FieldID, field.TypeUint),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rtuo.mutation.AccessTokensIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   refreshtoken.AccessTokensTable,
			Columns: []string{refreshtoken.AccessTokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(accesstoken.FieldID, field.TypeUint),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &RefreshToken{config: rtuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rtuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{refreshtoken.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	rtuo.mutation.done = true
	return _node, nil
}
