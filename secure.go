package main

import (
	"fmt"
	"golang.org/x/crypto/scrypt"
	"log"
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

var hashSalt = []byte("iconthin!")

const (
	n      = 8192 // Default 16384
	r      = 8    // Default 8
	p      = 1    // Default 1
	keyLen = 16   // Default 32
)

func hash(text string) string {
	keyDer, err := scrypt.Key([]byte(text), hashSalt, n, r, p, keyLen)
	if err != nil {
		log.Fatal("Failed to hash text:" + text)
		log.Println(err.Error())
		return ""
	}
	return ADEncoding.EncodeToString(keyDer)
}
