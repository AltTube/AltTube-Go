// Code generated by ent, DO NOT EDIT.

package ent

import (
	"AltTube-Go/ent/likevideo"
	"AltTube-Go/ent/video"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// VideoCreate is the builder for creating a Video entity.
type VideoCreate struct {
	config
	mutation *VideoMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (vc *VideoCreate) SetTitle(s string) *VideoCreate {
	vc.mutation.SetTitle(s)
	return vc
}

// SetDescription sets the "description" field.
func (vc *VideoCreate) SetDescription(s string) *VideoCreate {
	vc.mutation.SetDescription(s)
	return vc
}

// SetUploadDate sets the "upload_date" field.
func (vc *VideoCreate) SetUploadDate(t time.Time) *VideoCreate {
	vc.mutation.SetUploadDate(t)
	return vc
}

// SetUploader sets the "uploader" field.
func (vc *VideoCreate) SetUploader(s string) *VideoCreate {
	vc.mutation.SetUploader(s)
	return vc
}

// SetUploaderURL sets the "uploader_url" field.
func (vc *VideoCreate) SetUploaderURL(s string) *VideoCreate {
	vc.mutation.SetUploaderURL(s)
	return vc
}

// SetCreatedAt sets the "created_at" field.
func (vc *VideoCreate) SetCreatedAt(t time.Time) *VideoCreate {
	vc.mutation.SetCreatedAt(t)
	return vc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (vc *VideoCreate) SetNillableCreatedAt(t *time.Time) *VideoCreate {
	if t != nil {
		vc.SetCreatedAt(*t)
	}
	return vc
}

// SetUpdatedAt sets the "updated_at" field.
func (vc *VideoCreate) SetUpdatedAt(t time.Time) *VideoCreate {
	vc.mutation.SetUpdatedAt(t)
	return vc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (vc *VideoCreate) SetNillableUpdatedAt(t *time.Time) *VideoCreate {
	if t != nil {
		vc.SetUpdatedAt(*t)
	}
	return vc
}

// SetDeletedAt sets the "deleted_at" field.
func (vc *VideoCreate) SetDeletedAt(t time.Time) *VideoCreate {
	vc.mutation.SetDeletedAt(t)
	return vc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (vc *VideoCreate) SetNillableDeletedAt(t *time.Time) *VideoCreate {
	if t != nil {
		vc.SetDeletedAt(*t)
	}
	return vc
}

// SetThumbnailURL sets the "thumbnail_url" field.
func (vc *VideoCreate) SetThumbnailURL(s string) *VideoCreate {
	vc.mutation.SetThumbnailURL(s)
	return vc
}

// SetID sets the "id" field.
func (vc *VideoCreate) SetID(s string) *VideoCreate {
	vc.mutation.SetID(s)
	return vc
}

// AddLikeVideoIDs adds the "like_videos" edge to the LikeVideo entity by IDs.
func (vc *VideoCreate) AddLikeVideoIDs(ids ...int64) *VideoCreate {
	vc.mutation.AddLikeVideoIDs(ids...)
	return vc
}

// AddLikeVideos adds the "like_videos" edges to the LikeVideo entity.
func (vc *VideoCreate) AddLikeVideos(l ...*LikeVideo) *VideoCreate {
	ids := make([]int64, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return vc.AddLikeVideoIDs(ids...)
}

// Mutation returns the VideoMutation object of the builder.
func (vc *VideoCreate) Mutation() *VideoMutation {
	return vc.mutation
}

// Save creates the Video in the database.
func (vc *VideoCreate) Save(ctx context.Context) (*Video, error) {
	return withHooks(ctx, vc.sqlSave, vc.mutation, vc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (vc *VideoCreate) SaveX(ctx context.Context) *Video {
	v, err := vc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vc *VideoCreate) Exec(ctx context.Context) error {
	_, err := vc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vc *VideoCreate) ExecX(ctx context.Context) {
	if err := vc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vc *VideoCreate) check() error {
	if _, ok := vc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Video.title"`)}
	}
	if v, ok := vc.mutation.Title(); ok {
		if err := video.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Video.title": %w`, err)}
		}
	}
	if _, ok := vc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Video.description"`)}
	}
	if v, ok := vc.mutation.Description(); ok {
		if err := video.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Video.description": %w`, err)}
		}
	}
	if _, ok := vc.mutation.UploadDate(); !ok {
		return &ValidationError{Name: "upload_date", err: errors.New(`ent: missing required field "Video.upload_date"`)}
	}
	if _, ok := vc.mutation.Uploader(); !ok {
		return &ValidationError{Name: "uploader", err: errors.New(`ent: missing required field "Video.uploader"`)}
	}
	if v, ok := vc.mutation.Uploader(); ok {
		if err := video.UploaderValidator(v); err != nil {
			return &ValidationError{Name: "uploader", err: fmt.Errorf(`ent: validator failed for field "Video.uploader": %w`, err)}
		}
	}
	if _, ok := vc.mutation.UploaderURL(); !ok {
		return &ValidationError{Name: "uploader_url", err: errors.New(`ent: missing required field "Video.uploader_url"`)}
	}
	if v, ok := vc.mutation.UploaderURL(); ok {
		if err := video.UploaderURLValidator(v); err != nil {
			return &ValidationError{Name: "uploader_url", err: fmt.Errorf(`ent: validator failed for field "Video.uploader_url": %w`, err)}
		}
	}
	if _, ok := vc.mutation.ThumbnailURL(); !ok {
		return &ValidationError{Name: "thumbnail_url", err: errors.New(`ent: missing required field "Video.thumbnail_url"`)}
	}
	if v, ok := vc.mutation.ThumbnailURL(); ok {
		if err := video.ThumbnailURLValidator(v); err != nil {
			return &ValidationError{Name: "thumbnail_url", err: fmt.Errorf(`ent: validator failed for field "Video.thumbnail_url": %w`, err)}
		}
	}
	return nil
}

func (vc *VideoCreate) sqlSave(ctx context.Context) (*Video, error) {
	if err := vc.check(); err != nil {
		return nil, err
	}
	_node, _spec := vc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Video.ID type: %T", _spec.ID.Value)
		}
	}
	vc.mutation.id = &_node.ID
	vc.mutation.done = true
	return _node, nil
}

func (vc *VideoCreate) createSpec() (*Video, *sqlgraph.CreateSpec) {
	var (
		_node = &Video{config: vc.config}
		_spec = sqlgraph.NewCreateSpec(video.Table, sqlgraph.NewFieldSpec(video.FieldID, field.TypeString))
	)
	if id, ok := vc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := vc.mutation.Title(); ok {
		_spec.SetField(video.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := vc.mutation.Description(); ok {
		_spec.SetField(video.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := vc.mutation.UploadDate(); ok {
		_spec.SetField(video.FieldUploadDate, field.TypeTime, value)
		_node.UploadDate = value
	}
	if value, ok := vc.mutation.Uploader(); ok {
		_spec.SetField(video.FieldUploader, field.TypeString, value)
		_node.Uploader = value
	}
	if value, ok := vc.mutation.UploaderURL(); ok {
		_spec.SetField(video.FieldUploaderURL, field.TypeString, value)
		_node.UploaderURL = value
	}
	if value, ok := vc.mutation.CreatedAt(); ok {
		_spec.SetField(video.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := vc.mutation.UpdatedAt(); ok {
		_spec.SetField(video.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := vc.mutation.DeletedAt(); ok {
		_spec.SetField(video.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := vc.mutation.ThumbnailURL(); ok {
		_spec.SetField(video.FieldThumbnailURL, field.TypeString, value)
		_node.ThumbnailURL = value
	}
	if nodes := vc.mutation.LikeVideosIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   video.LikeVideosTable,
			Columns: []string{video.LikeVideosColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(likevideo.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// VideoCreateBulk is the builder for creating many Video entities in bulk.
type VideoCreateBulk struct {
	config
	err      error
	builders []*VideoCreate
}

// Save creates the Video entities in the database.
func (vcb *VideoCreateBulk) Save(ctx context.Context) ([]*Video, error) {
	if vcb.err != nil {
		return nil, vcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(vcb.builders))
	nodes := make([]*Video, len(vcb.builders))
	mutators := make([]Mutator, len(vcb.builders))
	for i := range vcb.builders {
		func(i int, root context.Context) {
			builder := vcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VideoMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, vcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, vcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vcb *VideoCreateBulk) SaveX(ctx context.Context) []*Video {
	v, err := vcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vcb *VideoCreateBulk) Exec(ctx context.Context) error {
	_, err := vcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vcb *VideoCreateBulk) ExecX(ctx context.Context) {
	if err := vcb.Exec(ctx); err != nil {
		panic(err)
	}
}
