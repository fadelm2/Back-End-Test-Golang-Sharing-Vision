package converter

import (
	"article/internal/entity"
	"article/internal/model"
)

func PostToResponse(post *entity.Post) *model.PostResponse {
	return &model.PostResponse{
		Title:    post.Title,
		Content:  post.Content,
		Category: post.Category,
		Status:   post.Status,
	}
}
