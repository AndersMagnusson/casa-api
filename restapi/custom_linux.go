// +build linux darwin

package restapi

import (
	"net/http"
	"strings"
)

func Web(location string) func(http.Handler) http.Handler {
	fs := http.FileServer(http.Dir("home/magnusson/frontend"))
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// fmt.Print("Path: ", req.URL.Path)
			if strings.HasPrefix(req.URL.Path, "/v1") || location == "cloud" {
				next.ServeHTTP(w, req)
			} else {
				fs.ServeHTTP(w, req)
			}
			// if req.URL.Path == "/" || strings.Contains(req.URL.Path, ".js") || strings.Contains(req.URL.Path, ".css") {
			// 	fmt.Println("Serving from disk: ", req.URL.Path)
			// 	// http.StripPrefix("/", fs).ServeHTTP(w, req)
			// 	fs.ServeHTTP(w, req)
			// } else {
			// 	fmt.Println("Serving from api: ", req.URL.Path)
			// 	next.ServeHTTP(w, req)
			// }
		})
	}
}
