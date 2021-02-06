package ressources

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

var MainLayout *template.Template

func init() {
	webRoot := os.Getenv("WEBROOT")
	//lytDir := http.Dir(filepath.Join(webRoot, "layouts"))
	mainLytPath := filepath.Join(webRoot, "layouts", "*.html")
	log.Printf("WEBROOT/layouts=%s", mainLytPath)
	MainLayout = template.Must(template.New("main").ParseGlob(mainLytPath))
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, tpl *template.Template, name string, data interface{}) {
	buf := new(bytes.Buffer)
	if err := tpl.ExecuteTemplate(buf, name, data); err != nil {
		log.Printf("Error%s", MainLayout.DefinedTemplates())
		log.Printf("\nRender Error: %v\n", err)
		return
	}
	w.Write(buf.Bytes())
}

// Push the given resource to the client (HTTP/2).
func Push(w http.ResponseWriter, resource string) {
	pusher, ok := w.(http.Pusher)
	if ok {
		if err := pusher.Push(resource, nil); err == nil {
			return
		}
	}
}
