package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/mfmayer/go-vue/internal/api"
	"github.com/mfmayer/go-vue/internal/log"
	"github.com/mfmayer/go-vue/internal/vfs/vfswebui"
)

//go:generate go run internal/vfs/generate_vfswebui.go

func installFileServer(router chi.Router, path string, root http.FileSystem) error {
	if strings.ContainsAny(path, "{}*") {
		return fmt.Errorf("FileServer does not permit URL parameters")
	}
	if path != "/" && path[len(path)-1] != '/' {
		router.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	fs := http.StripPrefix(path, http.FileServer(root))
	router.Route(path, func(r chi.Router) {
		//r.Use(middleware.DefaultCompress)
		r.Get("/*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fs.ServeHTTP(w, r)
		}))
	})
	return nil
}

func main() {
	router := chi.NewRouter()

	router.Route("/api/", api.InstallAPI)
	router.Get("/", http.RedirectHandler("/ui/", 301).ServeHTTP)
	if err := installFileServer(router, "/ui", vfswebui.FileSystem); err != nil {
		panic(err)
	}

	uiAddr := ":9055"
	log.Info.Printf("Web UI listening on: %v", uiAddr)
	if err := http.ListenAndServe(uiAddr, router); err != nil {
		panic(err)
	}

}
