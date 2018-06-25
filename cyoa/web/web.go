package web

import (
	"fmt"
	"gophercises/cyoa/data"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var (
	defaultHandlerTemplate *template.Template
)

type handler struct {
	story data.Story
	tmpl  *template.Template
	css   *template.CSS
}

// HandlerOption is being used for additional setting in handler
type HandlerOption func(h *handler) error

// NewHandler return the handler being used for rendering the webpage
func NewHandler(s data.Story, opts ...HandlerOption) (http.Handler, error) {
	h := handler{
		story: s,
	}

	defaultHandleTmpl, err := ioutil.ReadFile("web/web.html")
	if err != nil {
		return handler{}, err
	}
	h.tmpl = template.Must(template.New("").Parse(string(defaultHandleTmpl)))

	for _, opt := range opts {
		if err := opt(&h); err != nil {
			return handler{}, err
		}
	}

	return h, nil
}

// WithTemplateFile return the function being used for setting the web template
// in the handler
func WithTemplateFile(filename string) HandlerOption {

	handleTmpl, err := ioutil.ReadFile(filename)
	if err != nil {
		return func(h *handler) error {
			return fmt.Errorf("cannot read the file : %v", err)
		}
	}

	return func(h *handler) error {
		h.tmpl = template.Must(template.New("").Parse(string(handleTmpl)))
		return nil
	}
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	path := strings.TrimSpace(r.URL.Path)
	if path == "/story" || path == "/story/" {
		path = "/story/intro"
	}
	path = path[len("/story/"):]

	var chapter data.Chapter
	var ok bool
	if chapter, ok = h.story[path]; !ok {
		log.Printf("cannot get to the path : %v\n", path)
		http.Error(w, "Chapter not found : "+path, http.StatusNotFound)
		return
	}

	err := h.tmpl.Execute(w, chapter)
	if err != nil {
		log.Printf("cannot execute the template : %v\n", err)
		http.Error(w, "Chapter not found : "+path, http.StatusNotFound)
		return
	}
}
