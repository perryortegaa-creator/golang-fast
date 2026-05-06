package posts

import (
	"belajar-go/internal/model/posts"
	"context"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
)

func (s *service) CreateComment(ctx context.Context, postID, userID int64, request posts.CreateCommentRequest) error {
	now := time.Now()
	model := posts.CommentModel{
		PostID:         postID,
		UserID:         userID,
		CommentContent: request.CommentContent,
		CreatedAt:      now,
		UpdatedAt:      now,
		CreatedBy:      strconv.FormatInt(userID, 10),
		UpdatedBy:      strconv.FormatInt(userID, 10),
	}
	err := s.postsRepo.CreateComment(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("failed to create comment to repository")
		return err
	}
	return nil
}
