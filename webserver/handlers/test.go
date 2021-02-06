package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/npillmayer/fahne/webserver/ressources"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	// fetch the url parameter `"userID"` from the request of a matching
	// routing pattern. An example routing pattern could be: /users/{userID}
	userID := chi.URLParam(r, "userID")

	// fetch `"key"` from the request context
	//ctx := r.Context()
	//key := ctx.Value("key").(string)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fullData := map[string]interface{}{
		"User": userID,
		//"Key":  key,
		//"NavigationBar": template.HTML(navigationBarHTML),
	}
	ressources.RenderTemplate(w, r, ressources.MainLayout, "main.html", fullData)
}

// respond to the client
//w.Write([]byte(fmt.Sprintf("Hi %v, %v", userID, key)))
