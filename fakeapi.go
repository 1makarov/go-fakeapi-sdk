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

	errorLoadPage = "error, status_code: %d"
)

type FakeApi struct {
	api *api
}

func New() *FakeApi {
	t := newTransport()
	a := newAPI(t)

	return &FakeApi{api: a}
}

func (f *FakeApi) GetPostByID(ID int) (*PostOutput, error) {
	body, code, err := f.api.call(endDefault, fmt.Sprintf(endGetPostByID, ID), http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	if code != http.StatusOK {
		return nil, fmt.Errorf(errorLoadPage, code)
	}

	var post PostOutput

	if err = jsoniter.Unmarshal(body, &post); err != nil {
		return nil, err
	}

	return &post, nil
}

func (f *FakeApi) GetAllPosts() ([]PostOutput, error) {
	body, code, err := f.api.call(endDefault, endGetAllPosts, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	if code != http.StatusOK {
		return nil, fmt.Errorf(errorLoadPage, code)
	}

	var posts []PostOutput

	if err = jsoniter.Unmarshal(body, &posts); err != nil {
		return nil, err
	}

	return posts, nil
}

func (f *FakeApi) CreatePost(p PostCreateInput) (*PostOutput, error) {
	body, code, err := f.api.call(endDefault, endCreatePost, http.MethodPost, p)
	if err != nil {
		return nil, err
	}

	if code != http.StatusCreated {
		return nil, fmt.Errorf(errorLoadPage, code)
	}

	var post PostOutput

	if err = jsoniter.Unmarshal(body, &post); err != nil {
		return nil, err
	}

	return &post, nil
}

func (f *FakeApi) UpdatePost(p PostUpdateInput) (*PostOutput, error) {
	body, code, err := f.api.call(endDefault, fmt.Sprintf(endUpdatePost, p.PostID), http.MethodPut, p)
	if err != nil {
		return nil, err
	}

	if code != http.StatusOK {
		return nil, fmt.Errorf(errorLoadPage, code)
	}

	var post PostOutput

	if err = jsoniter.Unmarshal(body, &post); err != nil {
		return nil, err
	}

	return &post, nil
}

func (f *FakeApi) DeletePostByID(ID int) error {
	_, code, err := f.api.call(endDefault, fmt.Sprintf(endDeletePost, ID), http.MethodDelete, nil)
	if err != nil {
		return err
	}

	if code != http.StatusOK {
		return fmt.Errorf(errorLoadPage, code)
	}

	return nil
}
