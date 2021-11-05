package fakeapi

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"net/http"
)

const (
	endDefault = "https://jsonplaceholder.typicode.com/"

	endGetPostByID = "posts/%d"
	endGetAllPosts = "posts"
	endCreatePost  = "posts"
	endUpdatePost  = "posts/%d"
	endDeletePost  = "posts/%d"

	errLoadPage = "error, status_code: %d"
)

type Client struct {
	transport *transport
}

func New() *Client {
	return &Client{
		transport: newTransport(),
	}
}

func (f *Client) GetPostByID(ID int) (*PostOutput, error) {
	body, code, err := f.transport.call(endDefault, fmt.Sprintf(endGetPostByID, ID), http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	if code != http.StatusOK {
		return nil, fmt.Errorf(errLoadPage, code)
	}

	var post PostOutput

	if err = jsoniter.Unmarshal(body, &post); err != nil {
		return nil, err
	}

	return &post, nil
}

func (f *Client) GetAllPosts() ([]PostOutput, error) {
	body, code, err := f.transport.call(endDefault, endGetAllPosts, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	if code != http.StatusOK {
		return nil, fmt.Errorf(errLoadPage, code)
	}

	var posts []PostOutput

	if err = jsoniter.Unmarshal(body, &posts); err != nil {
		return nil, err
	}

	return posts, nil
}

func (f *Client) CreatePost(p PostCreateInput) (*PostOutput, error) {
	body, code, err := f.transport.call(endDefault, endCreatePost, http.MethodPost, p)
	if err != nil {
		return nil, err
	}
	if code != http.StatusCreated {
		return nil, fmt.Errorf(errLoadPage, code)
	}

	var post PostOutput

	if err = jsoniter.Unmarshal(body, &post); err != nil {
		return nil, err
	}

	return &post, nil
}

func (f *Client) UpdatePost(p PostUpdateInput) (*PostOutput, error) {
	body, code, err := f.transport.call(endDefault, fmt.Sprintf(endUpdatePost, p.PostID), http.MethodPut, p)
	if err != nil {
		return nil, err
	}
	if code != http.StatusOK {
		return nil, fmt.Errorf(errLoadPage, code)
	}

	var post PostOutput

	if err = jsoniter.Unmarshal(body, &post); err != nil {
		return nil, err
	}

	return &post, nil
}

func (f *Client) DeletePostByID(ID int) error {
	_, code, err := f.transport.call(endDefault, fmt.Sprintf(endDeletePost, ID), http.MethodDelete, nil)
	if err != nil {
		return err
	}
	if code != http.StatusOK {
		return fmt.Errorf(errLoadPage, code)
	}

	return nil
}
