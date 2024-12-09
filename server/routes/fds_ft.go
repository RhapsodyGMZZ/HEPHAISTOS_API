package routes

import (
	utils "hephaistos/middleware/utils"
	"log"
	"net/http"
	"os"
	"strings"
	"unicode"
)

const _BASE_DIR = "storage/FDS_FT/"

func FDS(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/FDS" {
		w.WriteHeader(404)
		utils.RenderHtml(w, "error/404")
		return
	}

	cp := r.URL.Query().Get("code_racine")
	if cp != "" {
		file_name := get_fds(cp)
		http.ServeFile(w, r, _BASE_DIR+"FDS/"+file_name)
	} else {
		utils.RenderHtml(w, "fds")
	}
}

func get_fds(cp string) (file_name string) {
	files, err := os.ReadDir(_BASE_DIR + "FDS")
	if err != nil {
		log.Println(err)
		return
	}
	for _, file := range files {
		root_code := strings.Split(file.Name(), "_")[0]
		if is_digit(root_code) && root_code == cp {
			file_name = file.Name()
		}
	}
	return
}

func is_digit(s string) bool {
	return !strings.ContainsFunc(s, func(r rune) bool {
		return unicode.IsLetter(r)
	})
}
