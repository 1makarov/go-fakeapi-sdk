## Installation
```sh
go get -u github.com/1makarov/go-fakeapi-sdk
```

### Initialization
```go
fake := fakeapi.New()
```

### Get Post By ID
```go
post, err := fake.GetPostByID(1)
```

### Get All Posts
```go
posts, err := fake.GetAllPosts()
```

### Create post
```go
post, err := fake.CreatePost(fakeapi.PostCreateInput{
    Title:  "title #1",
    Body:   "body #1",
    UserID: 001,
})
```

### Update post
```go
post, err := fake.UpdatePost(fakeapi.PostUpdateInput{
    Title:  "title #1",
    Body:   "body #1",
    UserID: 001,
    PostID: 23,
})
```

### Delete post
```go
err := fake.DeletePostByID(1)
```