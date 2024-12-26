package blogrenderer_test

import (
	"blogrenderer"
	"bytes"
	"io"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

var (
	aPost = blogrenderer.Post{
		Title:       "hello world",
		Body:        "This is a post with [link](http://example.com)",
		Description: "This is a description",
		Tags:        []string{"go", "tdd"},
	}
)

func TestRender(t *testing.T) {
	postRenderer := createRenderer(t)

	t.Run("It converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := postRenderer.Render(&buf, aPost)

		if err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())

	})

	t.Run("It renders an index of posts", func(t *testing.T) {
		buf := bytes.Buffer{}
		posts := []blogrenderer.Post{{Title: "Hello World"}, {Title: "Hello World 2"}}

		if err := postRenderer.RenderIndex(&buf, posts); err != nil {
			t.Fatal(err)
		}
		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRender(b *testing.B) {
	postRenderer := createRenderer(b)

	b.ResetTimer()
	for range b.N {
		postRenderer.Render(io.Discard, aPost)
	}
}

func createRenderer(tb testing.TB) blogrenderer.PostRenderer {
	postRenderer, err := blogrenderer.NewPostRenderer()
	if err != nil {
		tb.Fatal(err)
	}
	return *postRenderer
}
