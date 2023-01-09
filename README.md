## Installation
```sh
go get -u github.com/1makarov/go-fakeapi-sdk
```

### Initialization
```go
client = fakeapi.New(fakeapi.Endpoint, http.DefaultClient)
```

### Get Post By ID
```go
post, err := client.GetPostByID(1)
```

### Get All Posts
```go
posts, err := client.GetAllPosts()
```

### Create post
```go
post, err := client.CreatePost(100, "title #1", "body #1")
```

### Update post
```go
post, err := client.UpdatePost(100, 1, "title #1", "body #1")
```

### Delete post
```go
err := client.DeletePostByID(1)
```