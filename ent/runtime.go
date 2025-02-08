// Code generated by ent, DO NOT EDIT.

package ent

import (
	"AltTube-Go/ent/accesstoken"
	"AltTube-Go/ent/likevideo"
	"AltTube-Go/ent/refreshtoken"
	"AltTube-Go/ent/schema"
	"AltTube-Go/ent/user"
	"AltTube-Go/ent/video"
	"time"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	accesstokenMixin := schema.AccessToken{}.Mixin()
	accesstokenMixinFields0 := accesstokenMixin[0].Fields()
	_ = accesstokenMixinFields0
	accesstokenFields := schema.AccessToken{}.Fields()
	_ = accesstokenFields
	// accesstokenDescCreateTime is the schema descriptor for create_time field.
	accesstokenDescCreateTime := accesstokenMixinFields0[0].Descriptor()
	// accesstoken.DefaultCreateTime holds the default value on creation for the create_time field.
	accesstoken.DefaultCreateTime = accesstokenDescCreateTime.Default.(func() time.Time)
	// accesstokenDescUpdateTime is the schema descriptor for update_time field.
	accesstokenDescUpdateTime := accesstokenMixinFields0[1].Descriptor()
	// accesstoken.DefaultUpdateTime holds the default value on creation for the update_time field.
	accesstoken.DefaultUpdateTime = accesstokenDescUpdateTime.Default.(func() time.Time)
	// accesstoken.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	accesstoken.UpdateDefaultUpdateTime = accesstokenDescUpdateTime.UpdateDefault.(func() time.Time)
	likevideoMixin := schema.LikeVideo{}.Mixin()
	likevideoMixinFields0 := likevideoMixin[0].Fields()
	_ = likevideoMixinFields0
	likevideoFields := schema.LikeVideo{}.Fields()
	_ = likevideoFields
	// likevideoDescCreateTime is the schema descriptor for create_time field.
	likevideoDescCreateTime := likevideoMixinFields0[0].Descriptor()
	// likevideo.DefaultCreateTime holds the default value on creation for the create_time field.
	likevideo.DefaultCreateTime = likevideoDescCreateTime.Default.(func() time.Time)
	// likevideoDescUpdateTime is the schema descriptor for update_time field.
	likevideoDescUpdateTime := likevideoMixinFields0[1].Descriptor()
	// likevideo.DefaultUpdateTime holds the default value on creation for the update_time field.
	likevideo.DefaultUpdateTime = likevideoDescUpdateTime.Default.(func() time.Time)
	// likevideo.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	likevideo.UpdateDefaultUpdateTime = likevideoDescUpdateTime.UpdateDefault.(func() time.Time)
	refreshtokenMixin := schema.RefreshToken{}.Mixin()
	refreshtokenMixinFields0 := refreshtokenMixin[0].Fields()
	_ = refreshtokenMixinFields0
	refreshtokenFields := schema.RefreshToken{}.Fields()
	_ = refreshtokenFields
	// refreshtokenDescCreateTime is the schema descriptor for create_time field.
	refreshtokenDescCreateTime := refreshtokenMixinFields0[0].Descriptor()
	// refreshtoken.DefaultCreateTime holds the default value on creation for the create_time field.
	refreshtoken.DefaultCreateTime = refreshtokenDescCreateTime.Default.(func() time.Time)
	// refreshtokenDescUpdateTime is the schema descriptor for update_time field.
	refreshtokenDescUpdateTime := refreshtokenMixinFields0[1].Descriptor()
	// refreshtoken.DefaultUpdateTime holds the default value on creation for the update_time field.
	refreshtoken.DefaultUpdateTime = refreshtokenDescUpdateTime.Default.(func() time.Time)
	// refreshtoken.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	refreshtoken.UpdateDefaultUpdateTime = refreshtokenDescUpdateTime.UpdateDefault.(func() time.Time)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userMixinFields1 := userMixin[1].Fields()
	_ = userMixinFields1
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreateTime is the schema descriptor for create_time field.
	userDescCreateTime := userMixinFields0[0].Descriptor()
	// user.DefaultCreateTime holds the default value on creation for the create_time field.
	user.DefaultCreateTime = userDescCreateTime.Default.(func() time.Time)
	// userDescUpdateTime is the schema descriptor for update_time field.
	userDescUpdateTime := userMixinFields0[1].Descriptor()
	// user.DefaultUpdateTime holds the default value on creation for the update_time field.
	user.DefaultUpdateTime = userDescUpdateTime.Default.(func() time.Time)
	// user.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	user.UpdateDefaultUpdateTime = userDescUpdateTime.UpdateDefault.(func() time.Time)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[1].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescID is the schema descriptor for id field.
	userDescID := userMixinFields1[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
	videoMixin := schema.Video{}.Mixin()
	videoMixinFields0 := videoMixin[0].Fields()
	_ = videoMixinFields0
	videoFields := schema.Video{}.Fields()
	_ = videoFields
	// videoDescCreateTime is the schema descriptor for create_time field.
	videoDescCreateTime := videoMixinFields0[0].Descriptor()
	// video.DefaultCreateTime holds the default value on creation for the create_time field.
	video.DefaultCreateTime = videoDescCreateTime.Default.(func() time.Time)
	// videoDescUpdateTime is the schema descriptor for update_time field.
	videoDescUpdateTime := videoMixinFields0[1].Descriptor()
	// video.DefaultUpdateTime holds the default value on creation for the update_time field.
	video.DefaultUpdateTime = videoDescUpdateTime.Default.(func() time.Time)
	// video.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	video.UpdateDefaultUpdateTime = videoDescUpdateTime.UpdateDefault.(func() time.Time)
	// videoDescTitle is the schema descriptor for title field.
	videoDescTitle := videoFields[1].Descriptor()
	// video.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	video.TitleValidator = videoDescTitle.Validators[0].(func(string) error)
	// videoDescDescription is the schema descriptor for description field.
	videoDescDescription := videoFields[2].Descriptor()
	// video.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	video.DescriptionValidator = videoDescDescription.Validators[0].(func(string) error)
	// videoDescUploader is the schema descriptor for uploader field.
	videoDescUploader := videoFields[4].Descriptor()
	// video.UploaderValidator is a validator for the "uploader" field. It is called by the builders before save.
	video.UploaderValidator = videoDescUploader.Validators[0].(func(string) error)
	// videoDescUploaderURL is the schema descriptor for uploader_url field.
	videoDescUploaderURL := videoFields[5].Descriptor()
	// video.UploaderURLValidator is a validator for the "uploader_url" field. It is called by the builders before save.
	video.UploaderURLValidator = videoDescUploaderURL.Validators[0].(func(string) error)
	// videoDescThumbnailURL is the schema descriptor for thumbnail_url field.
	videoDescThumbnailURL := videoFields[6].Descriptor()
	// video.ThumbnailURLValidator is a validator for the "thumbnail_url" field. It is called by the builders before save.
	video.ThumbnailURLValidator = videoDescThumbnailURL.Validators[0].(func(string) error)
	// videoDescID is the schema descriptor for id field.
	videoDescID := videoFields[0].Descriptor()
	// video.IDValidator is a validator for the "id" field. It is called by the builders before save.
	video.IDValidator = videoDescID.Validators[0].(func(string) error)
}
