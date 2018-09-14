package gosling

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"os"
	"path"
)

const defaultTemplate = `<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Redirecting to {{ .RedirectURL }}</title>
    <link rel="canonical" href="{{ .RedirectURL }}">
    <meta http-equiv="refresh" content="0; url={{ .RedirectURL }}">
    <meta name="robots" content="noindex,follow">
  </head>
  <body>
    <a href="{{ .RedirectURL }}" alt="{{ .RedirectURL }}">Click here if you are not redirected.</a>
  </body>
</html>`

type templateData struct {
	RedirectURL string
}

// Redirects holds redirects
type Redirects map[string]string

// NewRedirects parses redirects from a JSON files
func NewRedirects(inpath string) (Redirects, error) {
	jsonFile, err := os.Open(inpath)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	var redir Redirects
	if err := json.NewDecoder(jsonFile).Decode(&redir); err != nil {
		return nil, err
	}

	return redir, nil
}

// BuildRedirects builds redirects to provided folder
func BuildRedirects(redir Redirects, inpath string) error {
	for from, to := range redir {
		dstDir := path.Clean(path.Join(inpath, from))
		if err := os.MkdirAll(dstDir, 0777); err != nil {
			return fmt.Errorf("unable to create a destination folder %s: %s", dstDir, err)
		}

		dstFile := path.Join(dstDir, "index.html")
		f, err := os.Create(dstFile)
		if err != nil {
			return fmt.Errorf("unable to create a destination file %s: %s", dstFile, err)
		}

		if err := BuildRedirect(to, f); err != nil {
			return fmt.Errorf("unable to build a redirect for %s: %s", to, err)
		}
	}
	return nil
}

// BuildRedirect builds a redirect to a destination writer
func BuildRedirect(link string, wr io.Writer) error {
	data := templateData{
		RedirectURL: link,
	}

	tmpl, err := template.New("redirect").Parse(defaultTemplate)
	if err != nil {
		return err
	}

	if err := tmpl.Execute(wr, &data); err != nil {
		return err
	}

	return nil
}
