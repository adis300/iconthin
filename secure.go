package main

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"
)

type secureHandler struct {
	code int
}

func secureRedirectHandler(code int) http.Handler {
	return &secureHandler{code}
}

func (rh *secureHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	oldpath := r.URL.Path
	if oldpath == "" { // should not happen, but avoid a crash if it does
		oldpath = "/"
	}
	// Use oldHost value to be the new url
	newURL := "https://" + strings.Replace(r.Host, PORT, PORT_SECURE, 1) + oldpath // Replace one instance
	w.Header().Set("Location", newURL)
	w.WriteHeader(rh.code)
	// RFC2616 recommends that a short note "SHOULD" be included in the
	// response because older user agents may not understand 301/307.
	// Shouldn't send the response for POST or HEAD; that leaves GET.
	if r.Method == "GET" {
		note := "<a href=\"" + template.HTMLEscapeString(newURL) + "\">" + http.StatusText(rh.code) + "</a>.\n"
		fmt.Fprintln(w, note)
	}
}
