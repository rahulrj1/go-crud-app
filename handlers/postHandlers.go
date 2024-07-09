package handlers

import (
	"context"
	"go-crud-app/model"
	"go-crud-app/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	repo repository.PostRepository
}

func NewPostHandler(repo repository.PostRepository) *PostHandler {
	return &PostHandler{repo: repo}
}

func (h *PostHandler) GetAllPosts(c *gin.Context) {
	posts, err := h.repo.GetAllPosts(context.Background())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error encountered", "error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Here is the list of all the posts", "posts": posts})
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	var newPost model.Post
	if err := c.ShouldBindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error Binding from JSON data"})
		return
	}

	result, err := h.repo.CreatePost(context.Background(), &newPost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error creating the post"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Post created successfully with title " + newPost.Title + " and body " + newPost.Body,
		"data":    result,
	})
}

func (h *PostHandler) GetPost(c *gin.Context) {
	id := c.Param("id")
	post, err := h.repo.GetPost(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Post not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully got the post", "post": post})
}

func (h *PostHandler) UpdatePost(c *gin.Context) {
	var newPost model.Post
	if err := c.ShouldBindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error Binding from JSON data"})
		return
	}

	id := c.Param("id")
	_, err := h.repo.UpdatePost(context.Background(), id, &newPost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error updating post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully updated the post", "post": newPost})
}

func (h *PostHandler) DeletePost(c *gin.Context) {
	id := c.Param("id")
	_, err := h.repo.DeletePost(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error deleting post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}
