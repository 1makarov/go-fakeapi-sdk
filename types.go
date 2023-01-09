package fakeapi

import "fmt"

type PostResponse struct {
	UserID int    `json:"userId"`
	PostID int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func (p *PostResponse) String() string {
	return fmt.Sprintf("UserID: %d, PostID: %d, Title: %s, Body: %s", p.UserID, p.PostID, p.Title, p.Body)
}

type PostCreateRequest struct {
	UserID int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type PostUpdateRequest struct {
	UserID int    `json:"userId"`
	PostID int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
