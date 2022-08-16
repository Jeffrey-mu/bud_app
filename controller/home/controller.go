package home

import (
	context "context"
)

// Controller for homes
type Controller struct {
}

// Home struct
type Home struct {
	ID int `json:"id"`
}

// Index of homes
// GET /home
func (c *Controller) Index(ctx context.Context) (homes []*Home, err error) {
	return []*Home{}, nil
}

// Show home
// GET /home/:id
func (c *Controller) Show(ctx context.Context, id int) (home *Home, err error) {
	return &Home{
		ID: id,
	}, nil
}
