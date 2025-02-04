package database

import (
	"AltTube-Go/ent"
	"AltTube-Go/ent/likevideo"
	"context"
)

// AddLikeVideo adds a like to a video by a user.
func AddLikeVideo(ctx context.Context, id string, videoID string) error {
	_, err := Client.LikeVideo.
		Create().
		SetUserID(id).
		SetVideoID(videoID).
		Save(ctx)
	return err
}

// ReadIsLikedVideo checks if a user has liked a video.
func ReadIsLikedVideo(ctx context.Context, id string, videoID string) (bool, error) {
	likedVideo, err := Client.LikeVideo.
		Query().
		Where(
			likevideo.UserIDEQ(id),
			likevideo.VideoIDEQ(videoID),
		).
		Only(ctx)
	return likedVideo != nil, err
}

// RemoveLikeVideo removes a like from a video by a user.
func RemoveLikeVideo(ctx context.Context, id string, videoID string) error {
	_, err := Client.LikeVideo.
		Delete().
		Where(
			likevideo.UserIDEQ(id),
			likevideo.VideoIDEQ(videoID),
		).
		Exec(ctx)
	return err
}

// RemoveAllLikesByUserID removes all likes by a user.
func RemoveAllLikesByUserID(ctx context.Context, id string) error {
	_, err := Client.LikeVideo.
		Delete().
		Where(likevideo.UserIDEQ(id)).
		Exec(ctx)
	return err
}

// GetAllLikesByUserID gets all likes by a user.
func GetAllLikesByUserID(ctx context.Context, id string) ([]*ent.LikeVideo, error) {
	likes, err := Client.LikeVideo.
		Query().
		Where(likevideo.UserIDEQ(id)).
		All(ctx)
	return likes, err
}
