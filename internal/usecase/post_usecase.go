package usecase

import (
	"article/internal/entity"
	"article/internal/model"
	"article/internal/model/converter"
	"article/internal/repository"
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PostUseCase struct {
	DB             *gorm.DB
	Log            *logrus.Logger
	Validate       *validator.Validate
	PostRepository *repository.PostsRepository
}

func NewPostUseCase(db *gorm.DB,
	logger *logrus.Logger,
	validate *validator.Validate,
	PostRepository *repository.PostsRepository) *PostUseCase {
	return &PostUseCase{
		DB:             db,
		Log:            logger,
		Validate:       validate,
		PostRepository: PostRepository,
	}
}

func (c *PostUseCase) Create(ctx context.Context,
	request *model.CreatePostRequest) (*model.PostResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, fiber.NewError(fiber.StatusBadRequest, "Input yang dimasukan ada kesalahan")
	}

	Post := &entity.Post{
		Title:    request.Title,
		Content:  request.Content, // harusnya open langsung default
		Category: request.Category,
		Status:   request.Status,
	}

	if err := c.PostRepository.Create(tx, Post); err != nil {
		c.Log.WithError(err).Error("error creating ticket Post")
		return nil, fiber.ErrInternalServerError

	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error creating ticket Post")
		return nil, fiber.ErrInternalServerError
	}

	return converter.PostToResponse(Post), nil

}
