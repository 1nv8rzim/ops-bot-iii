// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ritsec/ops-bot-iii/ent/predicate"
	"github.com/ritsec/ops-bot-iii/ent/shitposts"
	"github.com/ritsec/ops-bot-iii/ent/signin"
	"github.com/ritsec/ops-bot-iii/ent/user"
	"github.com/ritsec/ops-bot-iii/ent/vote"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uu *UserUpdate) SetNillableEmail(s *string) *UserUpdate {
	if s != nil {
		uu.SetEmail(*s)
	}
	return uu
}

// SetVerificationAttempts sets the "verification_attempts" field.
func (uu *UserUpdate) SetVerificationAttempts(i int8) *UserUpdate {
	uu.mutation.ResetVerificationAttempts()
	uu.mutation.SetVerificationAttempts(i)
	return uu
}

// SetNillableVerificationAttempts sets the "verification_attempts" field if the given value is not nil.
func (uu *UserUpdate) SetNillableVerificationAttempts(i *int8) *UserUpdate {
	if i != nil {
		uu.SetVerificationAttempts(*i)
	}
	return uu
}

// AddVerificationAttempts adds i to the "verification_attempts" field.
func (uu *UserUpdate) AddVerificationAttempts(i int8) *UserUpdate {
	uu.mutation.AddVerificationAttempts(i)
	return uu
}

// SetVerified sets the "verified" field.
func (uu *UserUpdate) SetVerified(b bool) *UserUpdate {
	uu.mutation.SetVerified(b)
	return uu
}

// SetNillableVerified sets the "verified" field if the given value is not nil.
func (uu *UserUpdate) SetNillableVerified(b *bool) *UserUpdate {
	if b != nil {
		uu.SetVerified(*b)
	}
	return uu
}

// AddSigninIDs adds the "signins" edge to the Signin entity by IDs.
func (uu *UserUpdate) AddSigninIDs(ids ...int) *UserUpdate {
	uu.mutation.AddSigninIDs(ids...)
	return uu
}

// AddSignins adds the "signins" edges to the Signin entity.
func (uu *UserUpdate) AddSignins(s ...*Signin) *UserUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uu.AddSigninIDs(ids...)
}

// AddVoteIDs adds the "votes" edge to the Vote entity by IDs.
func (uu *UserUpdate) AddVoteIDs(ids ...int) *UserUpdate {
	uu.mutation.AddVoteIDs(ids...)
	return uu
}

// AddVotes adds the "votes" edges to the Vote entity.
func (uu *UserUpdate) AddVotes(v ...*Vote) *UserUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return uu.AddVoteIDs(ids...)
}

// AddShitpostIDs adds the "shitposts" edge to the Shitposts entity by IDs.
func (uu *UserUpdate) AddShitpostIDs(ids ...string) *UserUpdate {
	uu.mutation.AddShitpostIDs(ids...)
	return uu
}

// AddShitposts adds the "shitposts" edges to the Shitposts entity.
func (uu *UserUpdate) AddShitposts(s ...*Shitposts) *UserUpdate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uu.AddShitpostIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearSignins clears all "signins" edges to the Signin entity.
func (uu *UserUpdate) ClearSignins() *UserUpdate {
	uu.mutation.ClearSignins()
	return uu
}

// RemoveSigninIDs removes the "signins" edge to Signin entities by IDs.
func (uu *UserUpdate) RemoveSigninIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveSigninIDs(ids...)
	return uu
}

// RemoveSignins removes "signins" edges to Signin entities.
func (uu *UserUpdate) RemoveSignins(s ...*Signin) *UserUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uu.RemoveSigninIDs(ids...)
}

// ClearVotes clears all "votes" edges to the Vote entity.
func (uu *UserUpdate) ClearVotes() *UserUpdate {
	uu.mutation.ClearVotes()
	return uu
}

// RemoveVoteIDs removes the "votes" edge to Vote entities by IDs.
func (uu *UserUpdate) RemoveVoteIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveVoteIDs(ids...)
	return uu
}

// RemoveVotes removes "votes" edges to Vote entities.
func (uu *UserUpdate) RemoveVotes(v ...*Vote) *UserUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return uu.RemoveVoteIDs(ids...)
}

// ClearShitposts clears all "shitposts" edges to the Shitposts entity.
func (uu *UserUpdate) ClearShitposts() *UserUpdate {
	uu.mutation.ClearShitposts()
	return uu
}

// RemoveShitpostIDs removes the "shitposts" edge to Shitposts entities by IDs.
func (uu *UserUpdate) RemoveShitpostIDs(ids ...string) *UserUpdate {
	uu.mutation.RemoveShitpostIDs(ids...)
	return uu
}

// RemoveShitposts removes "shitposts" edges to Shitposts entities.
func (uu *UserUpdate) RemoveShitposts(s ...*Shitposts) *UserUpdate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uu.RemoveShitpostIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, UserMutation](ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.VerificationAttempts(); ok {
		if err := user.VerificationAttemptsValidator(v); err != nil {
			return &ValidationError{Name: "verification_attempts", err: fmt.Errorf(`ent: validator failed for field "User.verification_attempts": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeString))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uu.mutation.VerificationAttempts(); ok {
		_spec.SetField(user.FieldVerificationAttempts, field.TypeInt8, value)
	}
	if value, ok := uu.mutation.AddedVerificationAttempts(); ok {
		_spec.AddField(user.FieldVerificationAttempts, field.TypeInt8, value)
	}
	if value, ok := uu.mutation.Verified(); ok {
		_spec.SetField(user.FieldVerified, field.TypeBool, value)
	}
	if uu.mutation.SigninsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SigninsTable,
			Columns: []string{user.SigninsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(signin.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedSigninsIDs(); len(nodes) > 0 && !uu.mutation.SigninsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SigninsTable,
			Columns: []string{user.SigninsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(signin.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.SigninsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SigninsTable,
			Columns: []string{user.SigninsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(signin.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.VotesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.VotesTable,
			Columns: []string{user.VotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vote.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedVotesIDs(); len(nodes) > 0 && !uu.mutation.VotesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.VotesTable,
			Columns: []string{user.VotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vote.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.VotesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.VotesTable,
			Columns: []string{user.VotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vote.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.ShitpostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ShitpostsTable,
			Columns: []string{user.ShitpostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(shitposts.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedShitpostsIDs(); len(nodes) > 0 && !uu.mutation.ShitpostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ShitpostsTable,
			Columns: []string{user.ShitpostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(shitposts.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.ShitpostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ShitpostsTable,
			Columns: []string{user.ShitpostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(shitposts.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableEmail(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetEmail(*s)
	}
	return uuo
}

// SetVerificationAttempts sets the "verification_attempts" field.
func (uuo *UserUpdateOne) SetVerificationAttempts(i int8) *UserUpdateOne {
	uuo.mutation.ResetVerificationAttempts()
	uuo.mutation.SetVerificationAttempts(i)
	return uuo
}

// SetNillableVerificationAttempts sets the "verification_attempts" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableVerificationAttempts(i *int8) *UserUpdateOne {
	if i != nil {
		uuo.SetVerificationAttempts(*i)
	}
	return uuo
}

// AddVerificationAttempts adds i to the "verification_attempts" field.
func (uuo *UserUpdateOne) AddVerificationAttempts(i int8) *UserUpdateOne {
	uuo.mutation.AddVerificationAttempts(i)
	return uuo
}

// SetVerified sets the "verified" field.
func (uuo *UserUpdateOne) SetVerified(b bool) *UserUpdateOne {
	uuo.mutation.SetVerified(b)
	return uuo
}

// SetNillableVerified sets the "verified" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableVerified(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetVerified(*b)
	}
	return uuo
}

// AddSigninIDs adds the "signins" edge to the Signin entity by IDs.
func (uuo *UserUpdateOne) AddSigninIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddSigninIDs(ids...)
	return uuo
}

// AddSignins adds the "signins" edges to the Signin entity.
func (uuo *UserUpdateOne) AddSignins(s ...*Signin) *UserUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uuo.AddSigninIDs(ids...)
}

// AddVoteIDs adds the "votes" edge to the Vote entity by IDs.
func (uuo *UserUpdateOne) AddVoteIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddVoteIDs(ids...)
	return uuo
}

// AddVotes adds the "votes" edges to the Vote entity.
func (uuo *UserUpdateOne) AddVotes(v ...*Vote) *UserUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return uuo.AddVoteIDs(ids...)
}

// AddShitpostIDs adds the "shitposts" edge to the Shitposts entity by IDs.
func (uuo *UserUpdateOne) AddShitpostIDs(ids ...string) *UserUpdateOne {
	uuo.mutation.AddShitpostIDs(ids...)
	return uuo
}

// AddShitposts adds the "shitposts" edges to the Shitposts entity.
func (uuo *UserUpdateOne) AddShitposts(s ...*Shitposts) *UserUpdateOne {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uuo.AddShitpostIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearSignins clears all "signins" edges to the Signin entity.
func (uuo *UserUpdateOne) ClearSignins() *UserUpdateOne {
	uuo.mutation.ClearSignins()
	return uuo
}

// RemoveSigninIDs removes the "signins" edge to Signin entities by IDs.
func (uuo *UserUpdateOne) RemoveSigninIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveSigninIDs(ids...)
	return uuo
}

// RemoveSignins removes "signins" edges to Signin entities.
func (uuo *UserUpdateOne) RemoveSignins(s ...*Signin) *UserUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uuo.RemoveSigninIDs(ids...)
}

// ClearVotes clears all "votes" edges to the Vote entity.
func (uuo *UserUpdateOne) ClearVotes() *UserUpdateOne {
	uuo.mutation.ClearVotes()
	return uuo
}

// RemoveVoteIDs removes the "votes" edge to Vote entities by IDs.
func (uuo *UserUpdateOne) RemoveVoteIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveVoteIDs(ids...)
	return uuo
}

// RemoveVotes removes "votes" edges to Vote entities.
func (uuo *UserUpdateOne) RemoveVotes(v ...*Vote) *UserUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return uuo.RemoveVoteIDs(ids...)
}

// ClearShitposts clears all "shitposts" edges to the Shitposts entity.
func (uuo *UserUpdateOne) ClearShitposts() *UserUpdateOne {
	uuo.mutation.ClearShitposts()
	return uuo
}

// RemoveShitpostIDs removes the "shitposts" edge to Shitposts entities by IDs.
func (uuo *UserUpdateOne) RemoveShitpostIDs(ids ...string) *UserUpdateOne {
	uuo.mutation.RemoveShitpostIDs(ids...)
	return uuo
}

// RemoveShitposts removes "shitposts" edges to Shitposts entities.
func (uuo *UserUpdateOne) RemoveShitposts(s ...*Shitposts) *UserUpdateOne {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uuo.RemoveShitpostIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks[*User, UserMutation](ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.VerificationAttempts(); ok {
		if err := user.VerificationAttemptsValidator(v); err != nil {
			return &ValidationError{Name: "verification_attempts", err: fmt.Errorf(`ent: validator failed for field "User.verification_attempts": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeString))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uuo.mutation.VerificationAttempts(); ok {
		_spec.SetField(user.FieldVerificationAttempts, field.TypeInt8, value)
	}
	if value, ok := uuo.mutation.AddedVerificationAttempts(); ok {
		_spec.AddField(user.FieldVerificationAttempts, field.TypeInt8, value)
	}
	if value, ok := uuo.mutation.Verified(); ok {
		_spec.SetField(user.FieldVerified, field.TypeBool, value)
	}
	if uuo.mutation.SigninsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SigninsTable,
			Columns: []string{user.SigninsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(signin.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedSigninsIDs(); len(nodes) > 0 && !uuo.mutation.SigninsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SigninsTable,
			Columns: []string{user.SigninsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(signin.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.SigninsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SigninsTable,
			Columns: []string{user.SigninsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(signin.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.VotesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.VotesTable,
			Columns: []string{user.VotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vote.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedVotesIDs(); len(nodes) > 0 && !uuo.mutation.VotesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.VotesTable,
			Columns: []string{user.VotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vote.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.VotesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.VotesTable,
			Columns: []string{user.VotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vote.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.ShitpostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ShitpostsTable,
			Columns: []string{user.ShitpostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(shitposts.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedShitpostsIDs(); len(nodes) > 0 && !uuo.mutation.ShitpostsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ShitpostsTable,
			Columns: []string{user.ShitpostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(shitposts.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.ShitpostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ShitpostsTable,
			Columns: []string{user.ShitpostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(shitposts.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
