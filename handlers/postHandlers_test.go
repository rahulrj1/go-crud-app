package handlers

import (
	"encoding/json"
	"go-crud-app/model"
	"go-crud-app/repository"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestPostAPIs(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	mockRepo := new(repository.MockPostRepository)
	handler := NewPostHandler(mockRepo)

	r := gin.Default()
	r.GET("/posts", handler.GetAllPosts)
	r.POST("/posts", handler.CreatePost)
	r.GET("/posts/:id", handler.GetPost)
	r.PUT("posts/:id", handler.UpdatePost)
	r.DELETE("posts/:id", handler.DeletePost)

	t.Run("Get All Posts", func(t *testing.T) {
		mockPosts := []model.Post{{Title: "Test Title", Body: "Test Body"}}
		mockRepo.On("GetAllPosts", mock.Anything).Return(mockPosts, nil)

		req, _ := http.NewRequest(http.MethodGet, "/posts", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "Test Title")
		mockRepo.AssertExpectations(t)
	})

	t.Run("Create Post", func(t *testing.T) {
		mockPost := &model.Post{Title: "Test Title", Body: "Test Body"}
		mockResult := &mongo.InsertOneResult{}
		mockRepo.On("CreatePost", mock.Anything, mockPost).Return(mockResult, nil)

		// postJSON := `{"title":"Test Title","body":"Test Body"}`
		postJSON, _ := json.Marshal(mockPost)
		req, _ := http.NewRequest(http.MethodPost, "/posts", strings.NewReader(string(postJSON)))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Contains(t, w.Body.String(), "Post created successfully with title")
		mockRepo.AssertExpectations(t)
	})

	t.Run("Get Post", func(t *testing.T) {
		mockPost := &model.Post{Title: "Test Title", Body: "Test Body"}
		mockRepo.On("GetPost", mock.Anything, "1").Return(mockPost, nil)

		req, _ := http.NewRequest(http.MethodGet, "/posts/1", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "Test Title")
		mockRepo.AssertExpectations(t)
	})

	t.Run("Update Post", func(t *testing.T) {
		mockPost := &model.Post{Title: "Updated Title", Body: "Updated Body"}
		mockResult := &mongo.UpdateResult{}
		mockRepo.On("UpdatePost", mock.Anything, "1", mockPost).Return(mockResult, nil)

		postJSON := `{"title":"Updated Title","body":"Updated Body"}`
		req, _ := http.NewRequest(http.MethodPut, "/posts/1", strings.NewReader(postJSON))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "Successfully updated the post")
		mockRepo.AssertExpectations(t)
	})

	t.Run("Delete Post", func(t *testing.T) {
		mockResult := &mongo.DeleteResult{}
		mockRepo.On("DeletePost", mock.Anything, "1").Return(mockResult, nil)

		req, _ := http.NewRequest(http.MethodDelete, "/posts/1", nil)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "Post deleted successfully")
		mockRepo.AssertExpectations(t)
	})
}
