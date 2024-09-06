package template

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/google/safehtml"
)

// BasePage for structure data pass to template
type BasePage struct {
	// HTMLTitle is the value to use in the page’s <title> tag.
	HTMLTitle string

	// MetaDescription is the html used for rendering the <meta name="Description"> tag.
	MetaDescription safehtml.HTML
}

// servePage is used to execute all templates for a *Server.
func (t *Template) ServePage(w http.ResponseWriter, templateName string, page any) {
	buf, err := t.RenderPage(templateName, page)
	if err != nil {
		log.Println(err)
		http.Error(w, "server error page", http.StatusInternalServerError)
	}
	if _, err := io.Copy(w, bytes.NewReader(buf)); err != nil {
		log.Println(err)
		http.Error(w, "server error page", http.StatusInternalServerError)
	}
}

// renderPage executes the given templateName with page.
func (t *Template) RenderPage(templateName string, page any) ([]byte, error) {
	tmpl, err := t.findTemplate(templateName)
	if err != nil {
		return nil, err
	}
	return executeTemplate(templateName, tmpl, page)
}

// newBasePage returns a base page for the given request and title.
func (t *Template) NewBasePage(_ *http.Request, title string) BasePage {
	return BasePage{
		HTMLTitle: title,
	}
}

// staticPageHandler handles requests to a template that contains no dynamic
// content.
func (t *Template) StaticPageHandler(templateName, title string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t.ServePage(w, templateName, t.NewBasePage(r, title))
	}
}
