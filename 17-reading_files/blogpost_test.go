package blogpost_test

import (
	"learn-go-with-tests/v2/17-reading_files/blogposts"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: Tag1, Tag2
---
Hello
World`

		secondBody = `Title: Post 2
Description: Description 2
Tags: Tag1, Tag2
---
Hola
Mundo`
	)

	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}

	posts, err := blogposts.NewPostsFromFS(fs)
	if err != nil {
		t.Fatal(err)
	}

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}

	assertPost(t, posts[0], blogposts.Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"Tag1", "Tag2"},
		Body:        "Hello\nWorld",
	})
}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
