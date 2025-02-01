package blogrenderer

import (
	"embed"
	"html/template"
	"io"
	"learn-go-with-tests/v2/17-reading_files/blogposts"
)

//go:embed "templates/*"
var postTemplates embed.FS

func Render(w io.Writer, p blogposts.Post) error {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return err
	}

	if err := templ.ExecuteTemplate(w, "blog.gohtml", p); err != nil {
		return err
	}

	return nil
}
