package main

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Subscriber defines the subscriber object
type Subscriber struct {
	Email     string `json:"email" gorm:"primary_key;column:email"`
	Active    bool   `json:"active" gorm:"column:active"`
	Timestamp int64  `json:"timestamp" gorm:"column:timestamp"`
}

// Feedback defines the feedback object
type Feedback struct {
	gorm.Model
	Name      string `json:"name" gorm:"column:name"`
	Phone     string `json:"phone" gorm:"column:phone"`
	Email     string `json:"email" gorm:"column:email"`
	Subject   string `json:"subject" gorm:"column:subject"`
	Body      string `json:"body" gorm:"column:body"`
	Active    bool   `json:"active" gorm:"column:active"`
	Timestamp int64  `json:"timestamp" gorm:"column:timestamp"`
}

// Response container
type Response struct {
	Data interface{} `json:"data"`
}

func createTables() {
	if !db.HasTable(&Subscriber{}) {
		log.Println("DBTableCreated:subscribers")
		db.CreateTable(&Subscriber{})
	}
	if !db.HasTable(&Feedback{}) {
		log.Println("DBTableCreated:feedbacks")
		db.CreateTable(&Feedback{})
	}
}

func subscribeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		email := strings.ToLower(strings.TrimSpace(r.FormValue("email")))
		if validateEmail(email) {
			var subscriber = &Subscriber{Email: email, Active: true, Timestamp: time.Now().Unix()}
			if err := db.Create(subscriber); err != nil {
				log.Println("WARNING:Email is already subscribed:" + email)
			}
			sendResponse(w, &Response{Data: subscriber})
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
