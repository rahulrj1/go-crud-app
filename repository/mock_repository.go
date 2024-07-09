package repository

import (
	"context"
	"go-crud-app/model"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
)

type MockPostRepository struct {
	mock.Mock
}

func (m *MockPostRepository) GetAllPosts(ctx context.Context) ([]model.Post, error) {
	args := m.Called(ctx)
	return args.Get(0).([]model.Post), args.Error(1)
}

func (m *MockPostRepository) CreatePost(ctx context.Context, post *model.Post) (*mongo.InsertOneResult, error) {
	args := m.Called(ctx, post)
	return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
}

func (m *MockPostRepository) GetPost(ctx context.Context, id string) (*model.Post, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*model.Post), args.Error(1)
}

func (m *MockPostRepository) UpdatePost(ctx context.Context, id string, post *model.Post) (*mongo.UpdateResult, error) {
	args := m.Called(ctx, id, post)
	return args.Get(0).(*mongo.UpdateResult), args.Error(1)
}

func (m *MockPostRepository) DeletePost(ctx context.Context, id string) (*mongo.DeleteResult, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*mongo.DeleteResult), args.Error(1)
}
