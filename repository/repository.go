package repository

import (
	"context"
	"go-crud-app/model"

	"go.mongodb.org/mongo-driver/mongo"
)

// Database Interface
type PostRepository interface {
	CreatePost(ctx context.Context, post *model.Post) (*mongo.InsertOneResult, error)
	GetAllPosts(ctx context.Context) ([]model.Post, error)
	GetPost(ctx context.Context, id string) (*model.Post, error)
	UpdatePost(ctx context.Context, id string, post *model.Post) (*mongo.UpdateResult, error)
	DeletePost(ctx context.Context, id string) (*mongo.DeleteResult, error)
}
