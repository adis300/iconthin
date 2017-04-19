package main

import (
	"github.com/renstrom/shortuuid"
	"net/http"
	"strings"
	"time"
)

var sessionCounter = 0

const sessionAvailableMunutes = 20

var currentToken = UUIDGen()

func adminLoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := strings.ToLower(strings.TrimSpace(r.FormValue("uname")))
		pswd := strings.TrimSpace(r.FormValue("pswd"))
		if validateUsername(username) && validatePassword(pswd) {
			if username == ADM_UNAME && hash(pswd) == ADM_PSWD {
				sessionCounter = 0
				currentToken = UUIDGen()
				w.Write([]byte(currentToken))
				return
			}
		}
		httpError400(w, http.StatusBadRequest)
	} else {
		httpError400(w, http.StatusMethodNotAllowed)
	}
}

// UUIDGen creates a short UUID string
func UUIDGen() string {
	return shortuuid.New()
}

func startAdminSessionTokenUpdateTask() {
	go func() {
		for {
			sessionCounter++
			if sessionCounter > sessionAvailableMunutes {
				sessionCounter = 0
				currentToken = UUIDGen()
			}
			time.Sleep(time.Second * 1)
		}
	}()
}
