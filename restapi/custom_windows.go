package restapi

import (
	"fmt"
	"net/http"
	"strings"
)

func Web() func(http.Handler) http.Handler {
	fs := http.FileServer(http.Dir("frontend"))
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			fmt.Print("Path: ", req.URL.Path)
			if req.URL.Path == "/" || strings.Contains(req.URL.Path, ".js") || strings.Contains(req.URL.Path, ".css") {
				fmt.Println("Serving from disk: ", req.URL.Path)
				// http.StripPrefix("/", fs).ServeHTTP(w, req)
				fs.ServeHTTP(w, req)
			} else {
				fmt.Println("Serving from api: ", req.URL.Path)
				next.ServeHTTP(w, req)
			}
		})
	}
}
