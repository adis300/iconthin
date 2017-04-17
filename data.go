package main

import (
	"net/http"
	"strings"
)

// Subscriber defines the subscriber object
type Subscriber struct {
	Email     string `json:"email"`
	Active    bool   `json:"active"`
	Timestamp int64  `json:"timestamp"`
}

// Feedback defines the feedback object
type Feedback struct {
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Subject   string `json:subject`
	Body      string `json:body`
	Active    bool   `json:"active"`
	Timestamp int64  `json:"timestamp"`
}

func subscribeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		email := strings.ToLower(strings.TrimSpace(r.FormValue("email")))
		if validateEmail(email) {
			print(email)
		} else {
			httpError400(w, http.StatusBadRequest)
		}
	} else {
		httpError400(w, http.StatusMethodNotAllowed)
	}
}

func feedbackHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		email := strings.ToLower(strings.TrimSpace(r.FormValue("email")))
		name := strings.ToLower(strings.TrimSpace(r.FormValue("name")))
		phone := strings.ToLower(strings.TrimSpace(r.FormValue("phone")))
		subject := strings.ToLower(strings.TrimSpace(r.FormValue("subject")))
		body := strings.ToLower(strings.TrimSpace(r.FormValue("body")))

		if subject != "" || body != "" {
			print(email)
			print(name)
			print(phone)
			print(subject)
			print(body)
		}
	} else {
		httpError400(w, http.StatusMethodNotAllowed)
	}
}
