package utils

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func ProxyToService(targetBaseURL string, pathPrefix string) http.HandlerFunc {
	target, err := url.Parse(targetBaseURL)
	if err != nil {
		fmt.Println("Error parsing the URL")
		return nil
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	originalDirector := proxy.Director

	proxy.Director = func(r *http.Request) {
		originalDirector(r)

		r.Host = target.Host

		fmt.Println("Proxying the request to: ", r.Host+r.URL.Path)
		fmt.Println("Original request path: ", r.URL.Path)
		fmt.Println("Path prefix: ", pathPrefix)
		
		r.URL.Path = strings.TrimPrefix(r.URL.Path, pathPrefix)

		fmt.Println("Modified request path: ", r.URL.Path)
		fmt.Println("Final request url: ", r.Host+r.URL.Path)

		if userId, ok := r.Context().Value("userId").(string); ok {
			r.Header.Set("X-User-ID", userId)
		}
	}

	return proxy.ServeHTTP
}