package templates

import (
	"bytes"
	"embed"
	"log"
	"path/filepath"

	//"io"
	"text/template"

	"github.com/gouniverse/utils"
)

//go:embed *
var files embed.FS

// Store defines a session store
type Template struct {
	Path string
	// Data     map[string]string
	Data     map[string]interface{}
	Minified bool
}

// StoreOption options for the cache store
type TemplateOption func(*Template)

// WithData sets the data for the template
func WithData(data map[string]interface{}) TemplateOption {
	return func(t *Template) {
		t.Data = data
	}
}

// WithPath sets the path for the template
func WithPath(path string) TemplateOption {
	return func(t *Template) {
		t.Path = path
	}
}

// WithMinified sets the data for the template
func WithMinified(minified bool) TemplateOption {
	return func(t *Template) {
		t.Minified = minified
	}
}

// NewTemplate creates a new template
func NewTemplate(path string, opts ...TemplateOption) *Template {
	template := &Template{}

	template.Path = path

	for _, opt := range opts {
		opt(template)
	}

	if template.Path == "" {
		log.Panic("Template: path is required")
	}

	return template
}

func (t Template) Apply(data map[string]interface{}) Template {
	t.Data = data
	return t
}

func (t Template) Minify(minify bool) Template {
	t.Minified = minify
	return t
}

func (t Template) ToString() string {
	if t.Data == nil {
		str, err := files.ReadFile(t.Path)
		if err != nil {
			log.Println("Template: " + t.Path + " NOT FOUND")
			return ""
		}
		content := string(str)
		if t.Minified {
			content = minify(t.Path, content)
		}
		return content
	}

	// funcs := template.FuncMap{
	// 	"unmapJson": func(jsonStr string) interface{} {
	// 		var unknown interface{}
	// 		json.Unmarshal([]byte(jsonStr), &unknown)
	// 		return unknown
	// 	},
	// }

	parsed := template.Must(template.ParseFS(files, t.Path))
	var tpl bytes.Buffer
	if err := parsed.Execute(&tpl, t.Data); err != nil {
		log.Println(err)
		return ""
	}

	content := tpl.String()
	if t.Minified {
		content = minify(t.Path, content)
	}
	return content
}

func minify(path string, content string) string {
	extension := filepath.Ext(path)
	if extension == ".html" {
		min, err := utils.MinHTML(content)
		if err != nil {
			return content
		}
		return min
	}

	if extension == ".js" {
		min, err := utils.MinHTML(content)
		if err != nil {
			return content
		}
		return min
	}

	if extension == ".css" {
		min, err := utils.MinCSS(content)
		if err != nil {
			return content
		}
		return min
	}

	return content
}


// func base64Decode(src string) ([]byte, error) {
// 	return base64.URLEncoding.DecodeString(src)
// }
