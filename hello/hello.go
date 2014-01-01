package hello

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
)

func init() {
	m := martini.Classic()
	m.Get("/", hello)

	m.Use(render.Renderer(render.Options{
		Directory:  "templates",
		Layout:     "layout",
		Extensions: []string{".tmpl", ".html"},
		Funcs: []template.FuncMap{{
			"NowYear": NowYear,
		}},
		Delims:  render.Delims{"{[{", "}]}"},
		Charset: "UTF-8",
	}))

	http.Handle("/", m)
}

func NowYear() string {
	start := "2013"
	now := time.Now().Format("2006")
	if start == now {
		return now
	} else {
		return fmt.Sprintf("%s-%s", start, now)
	}
}

func hello(render render.Render, w http.ResponseWriter, req *http.Request) {
	render.HTML(200, "hello", struct{ Name string }{"唐古拉山"})
}
