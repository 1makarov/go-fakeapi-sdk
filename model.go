package fakeapi

import "fmt"

type PostOutput struct {
	UserID int    `json:"userId"`
	PostID int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func (p *PostOutput) String() string {
	return fmt.Sprintf("UserID: %d, PostID: %d, Title: %s, Body: %s\n", p.UserID, p.PostID, p.Title, p.Body)
}

type PostCreateInput struct {
	UserID int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type PostUpdateInput struct {
	UserID int    `json:"userId"`
	PostID int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
