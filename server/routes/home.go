package routes

import (
	utils "hephaistos/middleware/utils"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(404)
		utils.RenderHtml(w, "error/404")
		return
	}
	utils.RenderHtml(w, "home")
}
