package post

import (
	context "context"
)

// Controller for posts
type Controller struct {
}

// Post struct
type Post struct {
	ID int `json:"id"`
}

// Index of posts
// GET /post
func (c *Controller) Index(ctx context.Context) (posts []*Post, err error) {
	return []*Post{}, nil
}

// Show post
// GET /post/:id
func (c *Controller) Show(ctx context.Context, id int) (post *Post, err error) {
	return &Post{
		ID: id,
	}, nil
}
