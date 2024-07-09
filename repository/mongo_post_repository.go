// repository/mongo_post_repository.go
package repository

import (
	"context"
	"go-crud-app/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoPostRepository struct {
	collection *mongo.Collection
}

func NewMongoPostRepository(collection *mongo.Collection) PostRepository {
	return &MongoPostRepository{collection: collection}
}

func (r *MongoPostRepository) GetAllPosts(ctx context.Context) ([]model.Post, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var posts []model.Post
	if err := cursor.All(ctx, &posts); err != nil {
		return nil, err
	}

	return posts, nil
}

func (m *MongoPostRepository) CreatePost(ctx context.Context, post *model.Post) (*mongo.InsertOneResult, error) {
	return m.collection.InsertOne(ctx, post)
}

func (m *MongoPostRepository) GetPost(ctx context.Context, id string) (*model.Post, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var post model.Post
	err = m.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&post)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (m *MongoPostRepository) UpdatePost(ctx context.Context, id string, post *model.Post) (*mongo.UpdateResult, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objID}
	update := bson.M{
		"$set": post,
	}
	return m.collection.UpdateOne(ctx, filter, update)
}

func (m *MongoPostRepository) DeletePost(ctx context.Context, id string) (*mongo.DeleteResult, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objID}
	return m.collection.DeleteOne(ctx, filter)
}
