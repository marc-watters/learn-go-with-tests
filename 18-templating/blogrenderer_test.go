package blogrenderer_test

import (
	"bytes"
	"learn-go-with-tests/v2/17-reading_files/blogposts"
	"learn-go-with-tests/v2/18-templating/blogrenderer"
	"testing"
)

func TestRender(t *testing.T) {
	aPost := blogposts.Post{
		Title:       "hello world",
		Body:        "This is a post",
		Description: "This is a description",
		Tags:        []string{"go", "tdd"},
	}

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := blogrenderer.Render(&buf, aPost)
		if err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := "<h1>hello world</h1><p>This is a description</p>Tags: <ul><li>go</li><li>tdd</li></ul>"

		if got != want {
			t.Errorf("\n\ngot:\n\n'%s'\n\nwant:\n\n'%s'", got, want)
		}
	})
}
