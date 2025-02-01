package blogrenderer

import (
	"fmt"
	"io"
	"learn-go-with-tests/v2/17-reading_files/blogposts"
)

func Render(w io.Writer, p blogposts.Post) error {
	_, err := fmt.Fprintf(w, "<h1>%s</h1>", p.Title)
	return err
}
