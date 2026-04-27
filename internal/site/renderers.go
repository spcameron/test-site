package site

import (
	"context"
	"io"

	"github.com/spcameron/press"
	"github.com/spcameron/test-site/templates"
)

func Renderers() press.Renderers {
	return press.Renderers{
		Home: func(w io.Writer, data press.HomePageData) error {
			return templates.Home(data).Render(context.Background(), w)
		},
		BlogIndex: func(w io.Writer, data press.BlogIndexPageData) error {
			return templates.BlogIndex(data).Render(context.Background(), w)
		},
		BlogPost: func(w io.Writer, data press.BlogPostPageData) error {
			return templates.BlogPost(data).Render(context.Background(), w)
		},
		StaticPage: func(w io.Writer, data press.StaticPageData) error {
			return templates.StaticPage(data).Render(context.Background(), w)
		},
	}
}
