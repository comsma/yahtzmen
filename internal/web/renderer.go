package web

import (
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"path/filepath"

	"github.com/comsma/yahtzmen/ui"
	"github.com/labstack/echo/v5"
)

type Renderer struct {
	templates map[string]*template.Template
}

func (r *Renderer) Render(c *echo.Context, w io.Writer, name string, data any) error {
	t, ok := r.templates[name]
	if !ok {
		return fmt.Errorf("template %q not found", name)
	}
	return t.ExecuteTemplate(w, name, data)
}

func NewTemplateCache() (*Renderer, error) {

	cache := map[string]*template.Template{}

	// get a slice of pages from ./ui/html/pages
	pages, err := fs.Glob(ui.Views, "views/pages/*.gohtml")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		patterns := []string{
			"views/layouts/base.gohtml",
			"views/partials/*.gohtml",
			page,
		}

		// register template functions and parse files matching patterns
		ts, err := template.New(name).ParseFS(ui.Views, patterns...)
		if err != nil {
			return nil, err
		}

		//add to the cache map
		cache[name] = ts
	}

	return &Renderer{templates: cache}, nil
}
