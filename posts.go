package fakeapi

import (
	"fmt"
	"net/http"
)

const (
	Endpoint = "https://jsonplaceholder.typicode.com/"

	_postByID   = "posts/%d"
	_allPosts   = "posts"
	_createPost = "posts"
	_updatePost = "posts/%d"
	_deletePost = "posts/%d"
)

func (c *Client) GetPostByID(ID int) (*PostResponse, error) {
	path := fmt.Sprintf(_postByID, ID)

	var post PostResponse

	return &post, c.get(c.url+path, &post, http.StatusOK)
}

func (c *Client) GetAllPosts() ([]PostResponse, error) {
	path := _allPosts

	var posts []PostResponse

	return posts, c.get(c.url+path, &posts, http.StatusOK)
}

func (c *Client) CreatePost(userID int, title, body string) (*PostResponse, error) {
	path := _createPost

	var (
		req = &PostCreateRequest{
			UserID: userID,
			Title:  title,
			Body:   body,
		}

		post PostResponse
	)

	return &post, c.post(c.url+path, &req, &post, http.StatusCreated)
}

func (c *Client) UpdatePost(userID, postID int, title, body string) (*PostResponse, error) {
	path := fmt.Sprintf(_updatePost, postID)

	var (
		req = &PostUpdateRequest{
			UserID: userID,
			PostID: postID,
			Title:  title,
			Body:   body,
		}

		post PostResponse
	)

	return &post, c.put(c.url+path, &req, &post, http.StatusOK)
}

func (c *Client) DeletePostByID(ID int) error {
	path := fmt.Sprintf(_deletePost, ID)

	return c.delete(c.url+path, nil, http.StatusOK)
}
