package posts

import (
	"belajar-go/internal/configs"
	"belajar-go/internal/model/posts"
	"context"
)

type postsRepository interface {
	CreatePost(ctx context.Context, model posts.PostModel) error
}

type service struct {
	cfg       *configs.Config
	postsRepo postsRepository
}

func NewService(cfg *configs.Config, postsRepo postsRepository) *service {
	return &service{
		cfg:       cfg,
		postsRepo: postsRepo,
	}
}
