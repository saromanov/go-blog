// Package blog defines inner representation of the blog
// models, db handling
package blog

import "time"


// Blog defines model for blog
type Blog struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Text string `json:"text"`
	CreatedAt time.Time
}