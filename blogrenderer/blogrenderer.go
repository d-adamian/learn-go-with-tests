package blogrenderer

import (
	"embed"
	"html/template"
	"io"
	"strings"

	"github.com/gomarkdown/markdown"
)

type Post struct {
	Title       string
	Body        string
	Description string
	Tags        []string
}

type postViewModel struct {
	Post
	HTMLBody template.HTML
}

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

type PostRenderer struct {
	templ *template.Template
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	return &PostRenderer{templ: templ}, nil
}

func (r *PostRenderer) Render(w io.Writer, p Post) error {
	viewModel := postViewModel{Post: p}
	viewModel.HTMLBody = template.HTML(markdown.ToHTML([]byte(p.Body), nil, nil))

	return r.templ.ExecuteTemplate(w, "blog.gohtml", viewModel)
}

func (p Post) SanitisedTitle() string {
	return strings.ToLower(strings.Replace(p.Title, " ", "-", -1))
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	return r.templ.ExecuteTemplate(w, "index.gohtml", posts)
}
